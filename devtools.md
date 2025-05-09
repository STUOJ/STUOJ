# 开发工具使用说明

## 清理生成文件工具

用于清理项目中以 `generated_` 开头的生成文件。

### 使用方法

1. 运行工具：
   ```bash
   go run .\dev\tools\clean_generated.go
   ```
2. 如果需要跳过确认直接删除所有生成文件：
   ```bash
   go run .\dev\tools\clean_generated.go -y
   ```

## 并发执行go generate工具

用于并发执行项目中的 `go:generate` 命令。

### 使用方法

1. 运行工具：
   ```bash
   go run .\dev\tools\concurrent_generate.go
   ```
2. 可选参数：
   - `-c <并发数>`：指定并发数量，默认为CPU核心数。
   - `-v`：启用详细输出模式。
   - `<目录路径>`：指定要扫描的目录，默认为当前目录。

示例：
```bash
go run .\dev\tools\concurrent_generate.go -c 16 -v
```

## AST查看器工具

用于解析和查看Go文件的AST结构。

### 使用方法

1. 构建工具：
   ```bash
   go build -o ast_viewer ast_viewer.go
   ```
2. 运行工具：
   ```bash
   ./ast_viewer <Go文件路径>
   ```
3. 输出的AST结构会保存到 `<Go文件路径>.ast` 文件中，同时也会打印到控制台。
