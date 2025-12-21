package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config 应用配置结构
type Config struct {
	WoodDetection WoodDetectionConfig `yaml:"wood_detection"`
	Server        ServerConfig        `yaml:"server"`
	Log           LogConfig           `yaml:"log"`
}

// WoodDetectionConfig 木材检测 API 配置
type WoodDetectionConfig struct {
	BaseURL       string `yaml:"base_url"`
	APIKey        string `yaml:"api_key"`
	Timeout       int    `yaml:"timeout"`        // 秒
	MaxRetries    int    `yaml:"max_retries"`    // 最大重试次数
	RetryInterval int    `yaml:"retry_interval"` // 重试间隔（秒）
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	ReadTimeout  int    `yaml:"read_timeout"`  // 秒
	WriteTimeout int    `yaml:"write_timeout"` // 秒
}

// LogConfig 日志配置
type LogConfig struct {
	Level  string `yaml:"level"`  // debug, info, warn, error
	Format string `yaml:"format"` // json, text
}

// LoadConfig 从文件加载配置
func LoadConfig(configPath string) (*Config, error) {
	// 如果配置文件不存在，使用默认配置
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return DefaultConfig(), nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 从环境变量覆盖配置（优先级更高）
	config.LoadFromEnv()

	return &config, nil
}

// LoadFromEnv 从环境变量加载配置（覆盖文件配置）
func (c *Config) LoadFromEnv() {
	if baseURL := os.Getenv("WOOD_DETECTION_BASE_URL"); baseURL != "" {
		c.WoodDetection.BaseURL = baseURL
	}
	if apiKey := os.Getenv("WOOD_DETECTION_API_KEY"); apiKey != "" {
		c.WoodDetection.APIKey = apiKey
	}
	if timeout := os.Getenv("WOOD_DETECTION_TIMEOUT"); timeout != "" {
		var t int
		if _, err := fmt.Sscanf(timeout, "%d", &t); err == nil {
			c.WoodDetection.Timeout = t
		}
	}
	if port := os.Getenv("SERVER_PORT"); port != "" {
		var p int
		if _, err := fmt.Sscanf(port, "%d", &p); err == nil {
			c.Server.Port = p
		}
	}
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		WoodDetection: WoodDetectionConfig{
			BaseURL:       "http://localhost:8000",
			APIKey:        "test-key-123",
			Timeout:       30,
			MaxRetries:    3,
			RetryInterval: 1,
		},
		Server: ServerConfig{
			Host:         "0.0.0.0",
			Port:         8080,
			ReadTimeout:  30,
			WriteTimeout: 30,
		},
		Log: LogConfig{
			Level:  "info",
			Format: "json",
		},
	}
}

// GetTimeout 获取超时时间
func (c *WoodDetectionConfig) GetTimeout() time.Duration {
	return time.Duration(c.Timeout) * time.Second
}

// GetRetryInterval 获取重试间隔
func (c *WoodDetectionConfig) GetRetryInterval() time.Duration {
	return time.Duration(c.RetryInterval) * time.Second
}

// GetAddress 获取服务器监听地址
func (s *ServerConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
