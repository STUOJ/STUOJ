//go:build ignore

package main

import (
	"STUOJ/utils/errchecker"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
			fmt.Printf("获取当前工作目录失败: %v\n", err)
			os.Exit(1)
		}

		// 构建默认路径
		*domainPath = filepath.Join(wd, "internal", "domain")
	}

	// 检查路径是否存在
	info, err := os.Stat(*domainPath)
	if err != nil {
		fmt.Printf("路径不存在: %s\n", *domainPath)
		os.Exit(1)
	}

	// 检查路径是否为目录
	if !info.IsDir() {
		fmt.Printf("%s 不是目录\n", *domainPath)
		os.Exit(1)
	}

	fmt.Printf("开始检查 %s 目录下的代码...\n", *domainPath)

	// 执行检查
	violations, err := errchecker.CheckDomainErrorTypes(*domainPath)
	if err != nil {
		fmt.Printf("检查失败: %v\n", err)
		os.Exit(1)
	}

	// 输出结果
	if len(violations) == 0 {
		fmt.Println("检查通过，未发现问题！")
		os.Exit(0)
	}

	fmt.Printf("发现 %d 个问题:\n", len(violations))
	for _, v := range violations {
		fmt.Println(v)
	}

	// 如果有违规，返回非零退出码
	os.Exit(1)
}
