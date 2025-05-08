//go:build ignore

// 并发执行go generate命令的工具
// 使用方法: go run .\dev\tools\concurrent_generate.go [-c <并发数>] [-v] [<目录路径>]
// -c 指定并发数量，默认为CPU核心数
// -v 启用详细输出模式
// <目录路径> 指定要扫描的目录，默认为当前目录
package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

// GenerateTask 表示一个go:generate任务
type GenerateTask struct {
	FilePath string // 文件路径
	Command  string // 生成命令
	Dir      string // 命令执行目录
}

// GenerateResult 表示一个生成任务的结果
type GenerateResult struct {
	Task      GenerateTask
	Success   bool
	Output    string
	Error     error
	TimeTaken time.Duration
}

// FindGenerateTasks 扫描目录查找所有包含go:generate注释的文件
func FindGenerateTasks(rootPath string) ([]GenerateTask, error) {
	var tasks []GenerateTask
	cmdSet := make(map[string]struct{}) // 用于去重的命令集合

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("访问路径失败: %w", err)
		}

		// 跳过.git目录
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		// 只处理Go文件
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			// 解析文件查找go:generate注释
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				return nil // 忽略解析错误，继续处理其他文件
			}

			// 检查文件中的注释
			for _, cg := range f.Comments {
				for _, c := range cg.List {
					if strings.HasPrefix(c.Text, "//go:generate") {
						cmd := strings.TrimPrefix(c.Text, "//go:generate")
						cmd = strings.TrimSpace(cmd)
						if cmd != "" {
							// 检查命令是否已存在
							if _, exists := cmdSet[cmd]; !exists {
								cmdSet[cmd] = struct{}{}
								// 获取文件所在目录作为命令执行目录
								dir := filepath.Dir(path)
								tasks = append(tasks, GenerateTask{
									FilePath: path,
									Command:  cmd,
									Dir:      dir,
								})
							}
						}
					}
				}
			}
		}
		return nil
	})

	return tasks, err
}

// ExecuteGenerateTasks 并发执行生成任务
func ExecuteGenerateTasks(tasks []GenerateTask, concurrency int, verbose bool) []GenerateResult {
	if concurrency <= 0 {
		concurrency = runtime.NumCPU()
	}

	taskCh := make(chan GenerateTask, len(tasks))
	resultCh := make(chan GenerateResult, len(tasks))

	// 启动工作协程
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskCh {
				result := executeTask(task, verbose)
				resultCh <- result
			}
		}()
	}

	// 发送任务
	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)

	// 等待所有工作协程完成
	wg.Wait()
	close(resultCh)

	// 收集结果
	var results []GenerateResult
	for result := range resultCh {
		results = append(results, result)
	}

	return results
}

// executeTask 执行单个生成任务
func executeTask(task GenerateTask, verbose bool) GenerateResult {
	result := GenerateResult{
		Task: task,
	}

	// 执行命令
	startTime := time.Now()

	if verbose {
		fmt.Printf("执行任务: %s\n", task.Command)
		fmt.Printf("文件: %s\n", task.FilePath)
		fmt.Printf("目录: %s\n\n", task.Dir)
	}

	// 创建命令对象
	cmd := exec.Command("cmd.exe", "/c", task.Command)
	cmd.Dir = task.Dir // 设置工作目录

	// 设置环境变量
	cmd.Env = os.Environ()
	// 提取文件名并设置GOFILE环境变量
	fileName := filepath.Base(task.FilePath)
	cmd.Env = append(cmd.Env, fmt.Sprintf("GOFILE=%s", fileName))

	// 获取命令输出
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		result.Success = false
		result.Error = fmt.Errorf("获取标准输出失败: %w", err)
		return result
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		result.Success = false
		result.Error = fmt.Errorf("获取标准错误输出失败: %w", err)
		return result
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		result.Success = false
		result.Error = fmt.Errorf("启动命令失败: %w", err)
		return result
	}

	// 读取命令输出
	var output strings.Builder

	// 处理标准输出
	stdoutScanner := bufio.NewScanner(stdoutPipe)
	go func() {
		for stdoutScanner.Scan() {
			line := stdoutScanner.Text()
			output.WriteString(line)
			output.WriteString("\n")
			if verbose {
				fmt.Println(line)
			}
		}
	}()

	// 处理标准错误输出
	stderrScanner := bufio.NewScanner(stderrPipe)
	go func() {
		for stderrScanner.Scan() {
			line := stderrScanner.Text()
			output.WriteString(line)
			output.WriteString("\n")
			if verbose {
				fmt.Println(line)
			}
		}
	}()

	// 等待命令执行完成
	err = cmd.Wait()
	result.TimeTaken = time.Since(startTime)
	result.Output = output.String()

	if err != nil {
		result.Success = false
		result.Error = fmt.Errorf("命令执行失败: %w", err)
	} else {
		result.Success = true
	}

	return result
}

// PrintSummary 打印执行结果摘要
func PrintSummary(results []GenerateResult) {
	successCount := 0
	failCount := 0
	totalTime := time.Duration(0)

	for _, result := range results {
		if result.Success {
			successCount++
		} else {
			failCount++
		}
		totalTime += result.TimeTaken
	}

	fmt.Println("\n执行结果摘要:")
	fmt.Printf("总任务数: %d\n", len(results))
	fmt.Printf("成功: %d\n", successCount)
	fmt.Printf("失败: %d\n", failCount)
	fmt.Printf("总耗时: %s\n", totalTime)

	// 如果有失败的任务，打印详细信息
	if failCount > 0 {
		fmt.Println("\n失败任务详情:")
		for _, result := range results {
			if !result.Success {
				fmt.Printf("文件: %s\n", result.Task.FilePath)
				fmt.Printf("命令: %s\n", result.Task.Command)
				fmt.Printf("错误: %v\n\n", result.Error)
			}
		}
	}
}

func main() {
	// 解析命令行参数
	concurrency := flag.Int("c", runtime.NumCPU(), "并发执行的任务数量")
	verbose := flag.Bool("v", false, "启用详细输出模式")
	flag.Parse()

	// 获取目录路径
	rootPath := "."
	if flag.NArg() > 0 {
		rootPath = flag.Arg(0)
	}

	// 获取绝对路径
	absRootPath, err := filepath.Abs(rootPath)
	if err != nil {
		log.Fatalf("获取绝对路径失败: %v\n", err)
	}

	fmt.Printf("扫描目录: %s\n", absRootPath)
	fmt.Printf("并发数量: %d\n\n", *concurrency)

	// 查找所有生成任务
	tasks, err := FindGenerateTasks(absRootPath)
	if err != nil {
		log.Fatalf("查找生成任务失败: %v\n", err)
	}

	if len(tasks) == 0 {
		fmt.Println("未找到任何go:generate任务")
		return
	}

	fmt.Printf("找到 %d 个go:generate任务\n\n", len(tasks))

	// 执行任务
	startTime := time.Now()
	results := ExecuteGenerateTasks(tasks, *concurrency, *verbose)
	totalTime := time.Since(startTime)

	// 打印摘要
	PrintSummary(results)
	fmt.Printf("\n总执行时间: %s\n", totalTime)
}
