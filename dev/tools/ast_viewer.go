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

// saveASTToFile 将AST结构保存到文件
func saveASTToFile(filename string, fset *token.FileSet, node ast.Node) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 重定向标准输出到文件
	oldStdout := os.Stdout
	os.Stdout = file
	defer func() { os.Stdout = oldStdout }()

	ast.Print(fset, node)
	return nil
}

func main() {
	// 解析命令行参数
	flag.Parse()

	filename := flag.Arg(0)
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// 处理输出
	outputFile := filepath.Base(filename) + ".ast"
	err = saveASTToFile(outputFile, fset, file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AST已保存到: %s\n", outputFile)
	ast.Print(fset, file)
}
