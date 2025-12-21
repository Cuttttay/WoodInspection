package app

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      Appconfig      `yaml:"app" mapstructure:"app"`
	Database DatabaseConfig `yaml:"database" mapstructure:"database"`
	JWT      JWTConfig      `yaml:"jwt" mapstructure:"jwt"`
}

type Appconfig struct {
	Port int `yaml:"port" mapstructure:"port"`
}

type JwcConfig struct {
	LoginURL string `yaml:"login_url" mapstructure:"login_url"`
}

type JWTConfig struct {
	Secret string `yaml:"secret" mapstructure:"secret"`
	Issuer string `yaml:"issuer" mapstructure:"issuer"`
}

type DatabaseConfig struct {
	Host string `yaml:"source" mapstructure:"source"`
	Port int    `yaml:"port" mapstructure:"port"`
	User string `yaml:"user" mapstructure:"user"`
	Pass string `yaml:"pass" mapstructure:"pass"`
	Name string `yaml:"name" mapstructure:"name"`
}

var Conf *Config

// LoadConfigFromPath 从指定路径加载配置
func LoadConfigFromPath(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Load config failed: %s", err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("Unmarshal config failed: %s", err)
	}

	return config, nil
}
