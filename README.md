# Augmented-LLM

![](https://00611-1258610398.cos.ap-nanjing.myqcloud.com/img/202504132159957.png)

## 介绍

Augmented-LLM 是在大语言模型的基础上，结合 MCP（多模态上下文处理）和 RAG（检索增强生成）等工具，使其具备信息检索、工具调用、读写记忆等功能。通过这些增强功能，Augmented-LLM 不仅能够理解和生成自然语言，还可以高效地从外部知识库中获取信息，执行复杂任务，并在交互过程中保持上下文一致性，从而提供更智能、更实用的解决方案。

> Augmented-LLM = Chat + MCP + RAG。

本项目使用的 LLM 与 Embedding 模型均为本地基于 Ollama 部署。

## 功能

- 自然语言交互
- 多模态上下文处理
- 检索增强生成
- 工具调用能力
- 上下文记忆管理

## 项目结构

```
.
├── main.go         # 程序入口
├── go.mod          # Go 模块配置
└── src/
    ├── api/        # API 接口层
    ├── config/     # 配置管理
    ├── core/       # 核心业务逻辑
    ├── models/     # 数据模型
    └── utils/      # 工具函数
```

## 环境要求

- Go 1.21 或更高版本
- 其他依赖将通过 go mod 自动管理

## 构建和运行

1. 克隆项目
```bash
git clone https://github.com/yourusername/Augmented-LLM.git
cd Augmented-LLM
```

2. 安装依赖
```bash
go mod tidy
```

3. 构建项目
```bash
go build -o augmented-llm
```

4. 运行项目
```bash
./augmented-llm
```

## 开发说明

- `src/api`: 存放 HTTP/gRPC 接口相关代码
- `src/config`: 配置文件和配置管理相关代码
- `src/core`: 核心业务逻辑，包括 LLM 调用、MCP 处理等
- `src/models`: 数据模型定义
- `src/utils`: 通用工具函数

