package core

import (
	"context"
	"fmt"

	"augmented-llm/src/utils"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

type LLMService struct {
	ollamaClient *ollama.ChatModel
	tools        []tool.BaseTool
	messages     []*schema.Message
}

func NewLLMService(ctx context.Context, baseURL string, model string, systmPrompt string, tools []tool.BaseTool, context string) *LLMService {
	client, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: baseURL,
		Model:   model,
	})

	if err != nil {
		panic(err)
	}

	// Set system prompt if provided
	messages := []*schema.Message{}
	if systmPrompt != "" {
		messages = append(messages, &schema.Message{
			Role:    "system",
			Content: systmPrompt,
		})
	}
	if context != "" {
		messages = append(messages, &schema.Message{
			Role:    "user",
			Content: context,
		})
	}

	return &LLMService{
		ollamaClient: client,
		tools:        tools,
		messages:     messages,
	}
}

func (l *LLMService) Chat(prompt string) {
	utils.LogTitle("Chat")

	if prompt != "" {
		l.messages = append(l.messages, &schema.Message{
			Role:    "user",
			Content: prompt,
		})
	}

	// TODO：处理工具调用，stream不仅会返回结果，还会返回工具调用的结果
	streamResult, err := l.ollamaClient.Stream(context.Background(), l.messages)
	if err != nil {
		fmt.Errorf("流式调用失败: %w", err)
	}

	utils.ReportStream(streamResult)
}
