//go:build ignore

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// 解析命令行参数
	domainPath := flag.String("path", "", "domain目录路径，默认为internal/domain")
	flag.Parse()

	// 如果未指定路径，使用默认路径
	if *domainPath == "" {
		// 获取当前工作目录
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "获取当前工作目录失败: %v\n", err)
			os.Exit(1)
		}

		// 构建默认路径
		*domainPath = filepath.Join(wd, "internal", "domain")
	}

	// 检查路径是否存在
	info, err := os.Stat(*domainPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "路径不存在: %s\n", *domainPath)
		os.Exit(1)
	}

	// 检查路径是否为目录
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "%s 不是目录\n", *domainPath)
		os.Exit(1)
	}

	fmt.Printf("开始检查 %s 目录下的代码...\n", *domainPath)

	// 执行检查
	violations, err := CheckDomainErrorTypes(*domainPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "检查失败: %v\n", err)
		os.Exit(1)
	}

	// 输出结果
	if len(violations) == 0 {
		fmt.Println("检查通过，未发现问题！")
		os.Exit(0)
	}

	fmt.Fprintf(os.Stderr, "发现 %d 个问题:\n", len(violations))
	for _, v := range violations {
		fmt.Fprintln(os.Stderr, v)
	}

	// 如果有违规，返回非零退出码
	os.Exit(1)
}

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

		// 打印正在检查的文件路径
		fmt.Printf("正在检查文件: %s\n", path)

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

		// 打印正在检查的函数名称和返回类型
		returnTypes := "无返回值"
		if funcDecl.Type.Results != nil {
			returnTypes = formatReturnTypes(funcDecl.Type.Results)
		}
		fmt.Printf("  正在检查函数: %s，返回类型: %s\n", funcDecl.Name.Name, returnTypes)

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
					pos := fset.Position(expr.Pos())
					violation := fmt.Sprintf("%s:%d:%d: 函数 %s 返回nil而不是自定义错误类型",
						filePath, pos.Line, pos.Column, funcDecl.Name.Name)
					violations = append(violations, violation)
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

// formatReturnTypes 格式化函数的返回类型为字符串
func formatReturnTypes(results *ast.FieldList) string {
	var types []string

	for _, field := range results.List {
		// 获取类型字符串
		typeStr := ""
		switch t := field.Type.(type) {
		case *ast.Ident:
			// 简单类型，如 int, string, error
			typeStr = t.Name
		case *ast.SelectorExpr:
			// 带包名的类型，如 time.Time
			if x, ok := t.X.(*ast.Ident); ok {
				typeStr = x.Name + "." + t.Sel.Name
			}
		case *ast.StarExpr:
			// 指针类型，如 *User
			if ident, ok := t.X.(*ast.Ident); ok {
				typeStr = "*" + ident.Name
			} else if sel, ok := t.X.(*ast.SelectorExpr); ok {
				// 带包名的指针类型，如 *errors.Error
				if x, ok := sel.X.(*ast.Ident); ok {
					typeStr = "*" + x.Name + "." + sel.Sel.Name
				}
			}
		case *ast.ArrayType:
			// 数组类型，如 []string
			typeStr = "[]"
			if ident, ok := t.Elt.(*ast.Ident); ok {
				typeStr += ident.Name
			}
		case *ast.MapType:
			// 映射类型，如 map[string]int
			typeStr = "map"
		case *ast.InterfaceType:
			// 接口类型
			typeStr = "interface{}"
		default:
			typeStr = "未知类型"
		}

		// 如果字段有多个名称，则为每个名称添加类型
		if len(field.Names) > 0 {
			for range field.Names {
				types = append(types, typeStr)
			}
		} else {
			// 匿名返回值
			types = append(types, typeStr)
		}
	}

	return strings.Join(types, ", ")
}

// isCustomError 检查表达式是否为自定义错误类型
func isCustomError(expr ast.Expr) bool {
	// 检查是否为 nil
	if isNil(expr) {
		return true
	}

	// 检查是否为标准库errors包的错误（通常是errors.New()形式）
	if call, ok := expr.(*ast.CallExpr); ok {
		if selector, ok := call.Fun.(*ast.SelectorExpr); ok {
			if x, ok := selector.X.(*ast.Ident); ok {
				// 如果是标准库errors包的New方法，则不是自定义错误
				if x.Name == "errors" && selector.Sel.Name == "New" {
					return false
				}
			}
		}
	}

	// 检查是否为 &errors.XXX 形式（自定义错误包的指针形式）
	if unary, ok := expr.(*ast.UnaryExpr); ok && unary.Op == token.AND {
		if selector, ok := unary.X.(*ast.SelectorExpr); ok {
			if x, ok := selector.X.(*ast.Ident); ok && x.Name == "errors" {
				return true
			}
		}
	}

	// 检查是否为 errors.XXX 形式（直接返回自定义错误变量）
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

	// 如果不是上述任何一种形式，则不是自定义错误
	return false
}
