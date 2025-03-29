package errchecker

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// CheckDomainErrorTypes 检查domain领域下的实体导出函数返回的error是否为自定义错误类型
func CheckDomainErrorTypes(domainPath string) ([]string, error) {
	var violations []string

	// 遍历domain目录
	err := filepath.Walk(domainPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录
		if info.IsDir() {
			return nil
		}

		// 只检查Go文件
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// 跳过值对象目录
		if strings.Contains(path, "valueobject") {
			return nil
		}

		// 跳过防腐层
		if strings.Contains(path, "judge0") || strings.Contains(path, "yuki") {
			return nil
		}

		// 解析Go文件
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		// 检查文件中的函数
		fileViolations := checkFile(node, fset, path)
		violations = append(violations, fileViolations...)

		return nil
	})

	return violations, err
}

// checkFile 检查单个文件中的函数
func checkFile(node *ast.File, fset *token.FileSet, filePath string) []string {
	var violations []string

	// 遍历所有声明
	ast.Inspect(node, func(n ast.Node) bool {
		// 检查函数声明
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		// 只检查导出的函数（首字母大写）
		if !funcDecl.Name.IsExported() {
			return true
		}

		// 检查函数是否有返回值
		if funcDecl.Type.Results == nil {
			return true
		}

		// 检查返回值中是否包含error类型，并记录error类型在返回值中的位置
		hasErrorReturn := false
		errorPositions := []int{} // 记录error类型在返回值中的位置
		for i, field := range funcDecl.Type.Results.List {
			// 处理多个相同类型的返回值，如 (a, b int, err error)
			if len(field.Names) > 0 {
				// 如果字段有多个名称，则计算实际位置
				basePos := 0
				for j := 0; j < i; j++ {
					basePos += len(funcDecl.Type.Results.List[j].Names)
				}

				for idx := range field.Names {
					expr, ok := field.Type.(*ast.Ident)
					if ok && expr.Name == "error" {
						hasErrorReturn = true
						errorPositions = append(errorPositions, basePos+idx)
					}
				}
			} else {
				// 计算当前位置
				pos := 0
				for j := 0; j < i; j++ {
					if len(funcDecl.Type.Results.List[j].Names) > 0 {
						pos += len(funcDecl.Type.Results.List[j].Names)
					} else {
						pos++
					}
				}

				expr, ok := field.Type.(*ast.Ident)
				if ok && expr.Name == "error" {
					hasErrorReturn = true
					errorPositions = append(errorPositions, pos)
				}
			}
		}

		// 如果没有error返回值，则跳过
		if !hasErrorReturn {
			return true
		}

		// 检查函数体中的return语句
		ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
			retStmt, ok := n.(*ast.ReturnStmt)
			if !ok {
				return true
			}

			// 只检查error位置的返回值
			for _, pos := range errorPositions {
				// 确保返回值数量足够
				if pos >= len(retStmt.Results) {
					continue
				}

				expr := retStmt.Results[pos]

				// 检查是否返回nil
				if isNil(expr) {
					continue
				}

				// 检查是否返回自定义错误类型
				if !isCustomError(expr) {
					pos := fset.Position(expr.Pos())
					violation := fmt.Sprintf("%s:%d:%d: 函数 %s 返回的error不是自定义错误类型",
						filePath, pos.Line, pos.Column, funcDecl.Name.Name)
					violations = append(violations, violation)
				}
			}
			return true
		})

		return true
	})

	return violations
}

// isNil 检查表达式是否为nil
func isNil(expr ast.Expr) bool {
	ident, ok := expr.(*ast.Ident)
	return ok && ident.Name == "nil"
}

// isCustomError 检查表达式是否为自定义错误类型
func isCustomError(expr ast.Expr) bool {
	// 检查是否为 nil
	if isNil(expr) {
		return true
	}

	// 检查是否为 &errors.XXX 形式
	if unary, ok := expr.(*ast.UnaryExpr); ok && unary.Op == token.AND {
		if selector, ok := unary.X.(*ast.SelectorExpr); ok {
			if x, ok := selector.X.(*ast.Ident); ok && x.Name == "errors" {
				return true
			}
		}
	}

	// 检查是否为 errors.XXX 形式（直接返回错误变量）
	if selector, ok := expr.(*ast.SelectorExpr); ok {
		if x, ok := selector.X.(*ast.Ident); ok && x.Name == "errors" {
			return true
		}
	}

	// 检查是否为方法调用形式
	if call, ok := expr.(*ast.CallExpr); ok {
		// 检查是否为 errors.XXX.WithXXX() 形式
		if selector, ok := call.Fun.(*ast.SelectorExpr); ok {
			// 检查是否为 errors.XXX.WithXXX()
			if inner, ok := selector.X.(*ast.SelectorExpr); ok {
				if x, ok := inner.X.(*ast.Ident); ok && x.Name == "errors" {
					return true
				}
			}

			// 检查是否为 errors.NewError() 形式
			if x, ok := selector.X.(*ast.Ident); ok && x.Name == "errors" {
				return true
			}
		}
	}

	return false
}
