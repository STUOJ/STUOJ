//go:build ignore

// 为domain目录下的实体结构体生成generated_query.go文件
// 例如：go run ../../../dev/gen/query_gen.go blog

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

const queryTmpl = `// Code generated by generate_const; DO NOT EDIT.

package {{.PackageName}}

import (
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	"STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/pkg/errors"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

type _Query struct{}

var Query = new(_Query)

func (*_Query) Select(context querycontext.{{.EntityName}}QueryContext, optionFunc ...option.QueryContextOption) ([]{{.EntityName}}, []map[string]any, error) {
	for _, o := range optionFunc {
		o(&context)
	}
	map_, err := dao.{{.EntityName}}Store.Select(context.GenerateOptions())
	if err != nil {
		return nil, nil, errors.ErrInternalServer.WithMessage("查询失败")
	}
	return Dtos(map_), map_, nil
}

func (*_Query) SelectOne(context querycontext.{{.EntityName}}QueryContext, optionFunc ...option.QueryContextOption) ({{.EntityName}}, map[string]any, error) {
	for _, o := range optionFunc {
		o(&context)
	}
	context.Page = option.NewPagination(0, 1)
	map_, err := dao.{{.EntityName}}Store.SelectOne(context.GenerateOptions())
	if err != nil {
		return {{.EntityName}}{}, nil, errors.ErrNotFound.WithMessage("未查询到该{{.VarName}}")
	}
	return Dto(map_), map_, nil
}

func (*_Query) Count(context querycontext.{{.EntityName}}QueryContext, optionFunc ...option.QueryContextOption) (int64, error) {
	for _, o := range optionFunc {
		o(&context)
	}
	res, err := dao.{{.EntityName}}Store.Count(context.GenerateOptions())
	if err != nil {
		return res, errors.ErrInternalServer.WithMessage("查询失败")
	}
	return res, nil
}

func (*_Query) GroupCount(context querycontext.{{.EntityName}}QueryContext) ([]option.GroupCountResult, error) {
	options := context.GenerateGroupCountOptions()
	if ok := options.Verify();!ok {
		return nil,errors.ErrInternalServer.WithMessage("分组字段验证失败")
	}
	res, err := dao.{{.EntityName}}Store.GroupCount(options)
	if err!= nil {
		return nil, errors.ErrInternalServer.WithMessage("查询失败")
	}
	return res, nil
}

{{if .HaveId}}
func (*_Query) SelectByIds(context querycontext.{{.EntityName}}QueryContext, optionFunc ...option.QueryContextOption) (map[int64]{{.EntityName}}, map[int64]map[string]any, error) {
	for _, o := range optionFunc {
		o(&context)
	}
	context.Page = option.NewPagination(0, int64(len(context.Id.Value())))
	res, err := dao.{{.EntityName}}Store.Select(context.GenerateOptions())
	if err!= nil {
		return nil, nil, errors.ErrInternalServer.WithMessage("查询失败")
	}
	res{{.EntityName}}s := make(map[int64]{{.EntityName}}, len(res))
	resMaps := make(map[int64]map[string]any, len(res))
	for _, v := range res {
		id := int64(v["id"].(uint64))
		res{{.EntityName}}s[id] = Dto(v)
		resMaps[id] = v
	}
	return res{{.EntityName}}s, resMaps, nil
}
{{end}}
`

type TemplateData struct {
	PackageName string
	StructName  string
	EntityName  string
	VarName     string
	HaveId      bool
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
	var haveId bool
	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok || typeSpec.Type == nil {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		// 找到第一个结构体定义
		structName = typeSpec.Name.Name

		// 检查是否有Id字段
		for _, field := range structType.Fields.List {
			if len(field.Names) == 0 {
				continue
			}

			fieldName := field.Names[0].Name
			if fieldName == "Id" {
				haveId = true
				break
			}
		}

		return false
	})

	if structName == "" {
		return fmt.Errorf("在文件 %s 中未找到结构体定义", entityFile)
	}

	// 生成generated_query.go文件
	queryFile := filepath.Join("generated_query.go")

	// 准备模板数据
	varName := strings.ToLower(structName[:1]) + structName[1:]
	entityStructName := strings.Title(entityName)
	templateData := TemplateData{
		PackageName: pkgName,
		StructName:  structName,
		EntityName:  entityStructName,
		VarName:     varName,
		HaveId:      haveId,
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

	// 写入文件
	if err := os.WriteFile(queryFile, formattedCode, 0644); err != nil {
		return fmt.Errorf("写入文件 %s 失败: %v", queryFile, err)
	}

	fmt.Printf("生成成功: %s %s\n", structName, queryFile)
	return nil
}
