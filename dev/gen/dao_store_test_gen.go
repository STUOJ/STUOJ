//go:build ignore

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const testTmpl = `// Code generated by dao_store_test_gen; DO NOT EDIT.
package dao_test

import (
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/option"
	"testing"
)

func Test{{.StructName}}Store_Insert(t *testing.T) {
	var e entity.{{.StructName}}
	_, err := dao.{{.StructName}}Store.Insert(e)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
}

func Test{{.StructName}}Store_Select(t *testing.T) {
	opts := option.NewQueryOptions()
	_, err := dao.{{.StructName}}Store.Select(opts)
	if err != nil {
		t.Errorf("Select failed: %v", err)
	}
}

func Test{{.StructName}}Store_SelectOne(t *testing.T) {
	opts := option.NewQueryOptions()
	_, err := dao.{{.StructName}}Store.SelectOne(opts)
	if err != nil {
		t.Errorf("SelectOne failed: %v", err)
	}
}

func Test{{.StructName}}Store_Updates(t *testing.T) {
	var e entity.{{.StructName}}
	opts := option.NewQueryOptions()
	_, err := dao.{{.StructName}}Store.Updates(e, opts)
	if err != nil {
		t.Errorf("Updates failed: %v", err)
	}
}

func Test{{.StructName}}Store_Delete(t *testing.T) {
	opts := option.NewQueryOptions()
	err := dao.{{.StructName}}Store.Delete(opts)
	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
}

func Test{{.StructName}}Store_Count(t *testing.T) {
	opts := option.NewQueryOptions()
	_, err := dao.{{.StructName}}Store.Count(opts)
	if err != nil {
		t.Errorf("Count failed: %v", err)
	}
}

func Test{{.StructName}}Store_GroupCount(t *testing.T) {
	opts := &option.GroupCountOptions{}
	_, err := dao.{{.StructName}}Store.GroupCount(opts)
	if err != nil {
		t.Errorf("GroupCount failed: %v", err)
	}
}

func Test{{.StructName}}Store_Dto(t *testing.T) {
	data := map[string]any{}
	_ = dao.{{.StructName}}Store.Dto(data)
}
`

func main() {
	var typeName string
	flag.StringVar(&typeName, "struct", "", "结构体类型名称")
	flag.Parse()
	if typeName == "" {
		log.Fatal("必须指定-struct参数")
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	var found bool
	for _, pkgFiles := range pkgs {
		for _, file := range pkgFiles.Files {
			for _, decl := range file.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok || genDecl.Tok != token.TYPE {
					continue
				}
				for _, spec := range genDecl.Specs {
					typeSpec, ok := spec.(*ast.TypeSpec)
					if !ok || typeSpec.Name.Name != typeName {
						continue
					}
					if _, ok := typeSpec.Type.(*ast.StructType); ok {
						found = true
						break
					}
				}
			}
		}
	}
	if !found {
		log.Fatalf("未找到类型 %s 对应的结构体", typeName)
	}

	// 生成测试文件
	outputDir := "../../../../test/unit/dao-test"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("test").Parse(testTmpl))
	data := map[string]any{
		"StructName": typeName,
	}
	outputFile := filepath.Join(outputDir, fmt.Sprintf("generated_%s_store_test.go", toSnakeCase(typeName)))
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := t.Execute(f, data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("生成成功: %s\n", outputFile)
}

// 下划线命名
func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}
