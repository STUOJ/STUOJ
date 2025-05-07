//go:build ignore

// AST查看器工具，用于解析和查看Go文件的AST结构
// 使用方法: go build -o ast_viewer ast_viewer.go
// 然后运行: ./ast_viewer <Go文件路径>
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

// parseFile 解析指定Go文件并返回AST
func parseFile(filename string) (*ast.File, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("解析文件失败: %v", err)
	}
	return file, nil
}

// printAST 打印AST结构到控制台
func printAST(node ast.Node, indent string) {
	fmt.Printf("%s%T\n", indent, node)
	if node == nil {
		return
	}

	switch n := node.(type) {
	case *ast.File:
		fmt.Printf("%s  Name: %s\n", indent, n.Name)
		for _, decl := range n.Decls {
			printAST(decl, indent+"  ")
		}
	case *ast.GenDecl:
		fmt.Printf("%s  Tok: %s\n", indent, n.Tok)
		for _, spec := range n.Specs {
			printAST(spec, indent+"  ")
		}
	case *ast.FuncDecl:
		fmt.Printf("%s  Name: %s\n", indent, n.Name)
		printAST(n.Type, indent+"  ")
	case *ast.FuncType:
		if n.Params != nil {
			printAST(n.Params, indent+"  ")
		}
		if n.Results != nil {
			printAST(n.Results, indent+"  ")
		}
	case *ast.FieldList:
		for _, field := range n.List {
			printAST(field, indent+"  ")
		}
	case *ast.Field:
		if len(n.Names) > 0 {
			fmt.Printf("%s  Names: %v\n", indent, n.Names)
		}
		printAST(n.Type, indent+"  ")
	case *ast.Ident:
		fmt.Printf("%s  Name: %s\n", indent, n.Name)
	case *ast.StructType:
		printAST(n.Fields, indent+"  ")
	case *ast.InterfaceType:
		printAST(n.Methods, indent+"  ")
	default:
		fmt.Printf("%s  %#v\n", indent, n)
	}
}

// saveASTToFile 将AST结构保存到文件
func saveASTToFile(filename string, node ast.Node) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 重定向标准输出到文件
	oldStdout := os.Stdout
	os.Stdout = file
	defer func() { os.Stdout = oldStdout }()

	printAST(node, "")
	return nil
}

func main() {
	// 解析命令行参数
	flag.Parse()

	filename := flag.Arg(0)
	file, err := parseFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// 处理输出
	outputFile := filepath.Base(filename) + ".ast"
	err = saveASTToFile(outputFile, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AST已保存到: %s\n", outputFile)
	printAST(file, "")
}
