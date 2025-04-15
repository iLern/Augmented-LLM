package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"strings"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/fatih/color"
)

// GenerateID 生成唯一ID
func GenerateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// GetTimestamp 获取当前时间戳
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// ToJSON 将对象转换为JSON字符串
func ToJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// FromJSON 从JSON字符串解析对象
func FromJSON(s string, v interface{}) error {
	return json.Unmarshal([]byte(s), v)
}

// SafeString 安全获取字符串值
func SafeString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ToJSON(v)
}

// MergeMaps 合并多个map
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// LogTitle 美化输出
func LogTitle(message string) {
	totalLength := 80
	messageLength := len(message)

	padding := max(0, totalLength-messageLength-4)
	paddedMessage := strings.Repeat("=", padding) + " " + message + " " + strings.Repeat("=", padding)

	cyanBold := color.New(color.FgCyan, color.Bold).SprintFunc()
	println(cyanBold(paddedMessage))
}

func ReportStream(sr *schema.StreamReader[*schema.Message]) {
	defer sr.Close()

	i := 0
	for {
		message, err := sr.Recv()
		if err == io.EOF { // 流式输出结束
			return
		}
		if err != nil {
			log.Fatalf("recv failed: %v", err)
		}
		log.Printf("%+v", message)
		i++
	}
}
