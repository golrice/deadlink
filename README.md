# DeadLink

DeadLink 是一个用于检测网页中死链（无效链接）的工具。它可以抓取指定网站的所有链接，验证其有效性，并生成详细的检测报告。

---

## **功能特性**

- **网页爬取**：递归抓取目标网站的所有页面链接。
- **链接检查**：验证每个链接的有效性，包括 HTTP 状态码和网络错误。
- **报告生成**：支持多种格式的检测报告（如控制台输出、CSV、JSON）。
- **并发处理**：通过 Goroutine 实现高效的并发爬取和检查。
- **可扩展性**：模块化设计，便于添加新功能或自定义行为。

---

## **目录结构**

```plaintext
deadlink/
├── cmd/                  # 程序入口点
│   └── client/           # 客户端
│       ├── cmd/          # cli接口
│       └── main.go       # 主程序入口
├── internal/             # 内部实现模块
│   ├── config/           # 配置管理模块
│   ├── crawler/          # 爬虫模块
│   ├── checker/          # 链接检查模块
│   ├── reporter/         # 报告生成模块
│   └── utils/            # 工具函数模块
├── pkg/                  # 公共模块
│   └── models/           # 数据模型定义
├── scripts/              # 辅助脚本
├── tests/                # 单元测试和集成测试
├── go.mod                # Go 模块依赖文件
├── go.sum                # Go 模块依赖校验文件
└── README.md             # 项目说明文档
```

---

## **安装与运行**

### **1. 安装依赖**
确保您的系统已安装 Go（建议版本 >= 1.20）。然后克隆项目并安装依赖：

```bash
git clone https://github.com/your-repo/deadlink.git
cd deadlink
make deps
```

### **2. 构建项目**
使用以下命令构建项目：

```bash
make build
```

构建完成后，会在项目根目录生成名为 `deadlink` 的可执行文件。

### **3. 运行项目**
直接运行项目，指定起始 URL：

```bash
./deadlink https://example.com
```

默认情况下，检测结果会保存为 `output.csv` 文件。

---

## **使用方法**

### **命令行参数**
```bash
Usage:
  deadlink [URL]

Flags:
  -h, --help   显示帮助信息
```

### **示例**
```bash
# 检测 example.com 的所有链接
./deadlink https://example.com

# 查看检测结果
cat output.csv
```

---

## **报告格式**

DeadLink 支持以下报告格式：

1. **控制台输出**：直接在终端显示检测结果。
2. **CSV 文件**：以表格形式保存检测结果。
3. **JSON 文件**：以结构化数据格式保存检测结果。

您可以修改 `cmd/client/cmd/root.go` 中的代码，更改输出路径或格式。

---

## **开发与测试**

### **1. 运行测试**
使用以下命令运行单元测试和集成测试：

```bash
make test
```

### **2. 代码质量检查**
运行代码静态分析和格式化工具：

```bash
make lint
```

### **3. 清理构建文件**
清理生成的二进制文件和其他临时文件：

```bash
make clean
```

---

## **模块说明**

### **(1) 爬虫模块**
- 负责抓取网页内容并提取所有链接。
- 支持递归爬取子页面，并限制最大深度。
- 使用 `goquery` 解析 HTML 页面。

### **(2) 检查器模块**
- 验证每个链接的有效性，返回状态码或错误信息。
- 使用并发处理提高性能。
- 支持忽略特定错误（如超时）。

### **(3) 报告生成器模块**
- 将检测结果格式化为人类可读的形式。
- 支持导出为 CSV、JSON 或 HTML 文件。

### **(4) 工具函数模块**
- 提供日志记录、字符串处理、文件操作等通用工具。

---

## **贡献指南**

欢迎为该项目贡献代码！请遵循以下步骤：

1. Fork 项目仓库。
2. 创建一个新的分支 (`git checkout -b feature/YourFeatureName`)。
3. 提交更改 (`git commit -m "Add some feature"`)。
4. 推送分支 (`git push origin feature/YourFeatureName`)。
5. 提交 Pull Request。

---

## **许可证**

本项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。
