//go:build ignore

// 为domain目录下的实体结构体生成query_generated.go文件
// 例如：go run ../../../utils/gen/query_gen.go blog

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const queryTmpl = `package {{.PackageName}}

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/querymodel"
)

type _Query struct{}

var Query = new(_Query)

func (*_Query) Select(model querymodel.{{.EntityName}}QueryModel) ([]map[string]any, error) {
	res, err := dao.{{.EntityName}}Store.Select(model.GenerateOptions())
	if err != nil {
		return res, errors.ErrInternalServer.WithMessage("查询失败")
	}
	return res, &errors.NoError
}

func (*_Query) SelectOne(model querymodel.{{.EntityName}}QueryModel) (map[string]any, error) {
	res, err := dao.{{.EntityName}}Store.SelectOne(model.GenerateOptions())
	if err != nil {
		return res, errors.ErrNotFound.WithMessage("未查询到该{{.VarName}}")
	}
	return res, &errors.NoError
}

func (*_Query) Count(model querymodel.{{.EntityName}}QueryModel) (int64, error) {
	res, err := dao.{{.EntityName}}Store.Count(model.GenerateOptions())
	if err != nil {
		return res, errors.ErrInternalServer.WithMessage("查询失败")
	}
	return res, &errors.NoError
}
`

type TemplateData struct {
	PackageName string
	StructName  string
	EntityName  string
	VarName     string
}

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run query_gen.go <entity-name>")
	}
	entityName := os.Args[1]

	// 获取当前工作目录，即实体所在的目录
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("获取当前工作目录失败: %v\n", err)
	}

	// 处理单个实体
	err = processEntity(cwd, entityName)
	if err != nil {
		log.Fatalf("处理实体失败: %v\n", err)
	}
}

// 处理单个实体的query生成
func processEntity(dir string, entityName string) error {
	// 获取包名
	pkgName := filepath.Base(dir)

	// 查找实体结构体定义文件
	entityFile := filepath.Join(dir, entityName+".go")
	if _, err := os.Stat(entityFile); os.IsNotExist(err) {
		return fmt.Errorf("实体文件 %s 不存在", entityFile)
	}

	// 解析实体结构体
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, entityFile, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("解析文件 %s 失败: %v", entityFile, err)
	}

	// 查找结构体定义
	var structName string
	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok || typeSpec.Type == nil {
			return true
		}

		_, ok = typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// 找到第一个结构体定义
		structName = typeSpec.Name.Name
		return false
	})

	if structName == "" {
		return fmt.Errorf("在文件 %s 中未找到结构体定义", entityFile)
	}

	// 生成query_generated.go文件
	queryFile := filepath.Join("..", "query_generated.go")

	// 准备模板数据
	varName := strings.ToLower(structName[:1]) + structName[1:]
	entityStructName := strings.Title(entityName)
	templateData := TemplateData{
		PackageName: pkgName,
		StructName:  structName,
		EntityName:  entityStructName,
		VarName:     varName,
	}

	// 渲染模板
	tmpl, err := template.New("query").Parse(queryTmpl)
	if err != nil {
		return fmt.Errorf("解析模板失败: %v", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, templateData); err != nil {
		return fmt.Errorf("渲染模板失败: %v", err)
	}

	// 格式化代码
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("格式化代码失败: %v\n%s", err, buf.String())
	}

	// 确保query目录存在
	queryDir := filepath.Join("..", "query")
	if _, err := os.Stat(queryDir); os.IsNotExist(err) {
		if err := os.MkdirAll(queryDir, 0755); err != nil {
			return fmt.Errorf("创建目录 %s 失败: %v", queryDir, err)
		}
	}

	// 写入文件
	if err := os.WriteFile(queryFile, formattedCode, 0644); err != nil {
		return fmt.Errorf("写入文件 %s 失败: %v", queryFile, err)
	}

	fmt.Printf("成功生成 %s\n", queryFile)
	return nil
}
