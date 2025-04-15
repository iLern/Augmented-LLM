package main

import (
	"augmented-llm/src/core"
	"context"
	"fmt"
)

func main() {
	llm := core.NewLLMService(context.Background(),
		"http://localhost:11434",
		"yi:6b",
		"",
		nil,
		"",
	)

	content, toolCalls := llm.Chat("你好，你是谁？")

	fmt.Printf("%+v\n", content)
	fmt.Printf("%+v\n", toolCalls)
}
