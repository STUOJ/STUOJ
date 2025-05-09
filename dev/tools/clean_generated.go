//go:build ignore

// 清理生成文件的工具
// 使用方法: go run .\dev\tools\clean_generated.go [-y]
// -y 表示跳过确认，直接删除所有生成文件
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CollectGeneratedFiles 收集所有生成文件
func CollectGeneratedFiles(rootPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("访问路径失败: %w", err)
		}

		// 跳过.git目录
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		// 收集生成文件
		if !info.IsDir() && strings.HasPrefix(info.Name(), "generated_") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// CleanGeneratedFiles 执行批量删除操作
func CleanGeneratedFiles(files []string, skipConfirm bool) error {
	if len(files) == 0 {
		fmt.Println("未找到生成文件")
		return nil
	}

	if !skipConfirm {
		fmt.Println("发现生成文件:")
		for _, f := range files {
			fmt.Println("  ", f)
		}

		fmt.Printf("\n确认删除全部 %d 个文件? (y/n) ", len(files))
		var confirm string
		fmt.Scanln(&confirm)
		if strings.ToLower(confirm) != "y" {
			fmt.Println("取消删除操作")
			return nil
		}
	}

	for _, f := range files {
		if err := os.Remove(f); err != nil {
			return fmt.Errorf("删除文件失败: %w", err)
		}
		fmt.Printf("成功删除: %s\n", f)
	}
	return nil
}

func main() {
	// 默认项目根目录为当前目录的父目录
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取工作目录失败: %v\n", err)
		os.Exit(1)
	}
	rootPath := filepath.Join(wd)
	skipConfirm := len(os.Args) > 1 && os.Args[1] == "-y"

	fmt.Printf("扫描根目录: %s\n", rootPath)
	fmt.Println("扫描生成文件...")
	files, err := CollectGeneratedFiles(rootPath)
	if err != nil {
		fmt.Printf("扫描失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("开始清理生成文件...")
	if err := CleanGeneratedFiles(files, skipConfirm); err != nil {
		fmt.Printf("清理失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("清理完成")
}
