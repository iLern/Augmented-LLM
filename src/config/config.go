package config

type Config struct {
	Server   ServerConfig   `json:"server"`
	LLM      LLMConfig      `json:"llm"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

type LLMConfig struct {
	Model     string `json:"model"`
	APIKey    string `json:"api_key"`
	APIBase   string `json:"api_base"`
	MaxTokens int    `json:"max_tokens"`
}

type DatabaseConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

var GlobalConfig Config

func Init() error {
	// TODO: 从配置文件或环境变量加载配置
	GlobalConfig = Config{
		Server: ServerConfig{
			Port: 8080,
			Host: "localhost",
		},
		LLM: LLMConfig{
			Model:     "gpt-4",
			MaxTokens: 2000,
		},
	}
	return nil
}
