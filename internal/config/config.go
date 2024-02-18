package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Mode                string `mapstructure:"MODE"`
	PostgresDSN         string `mapstructure:"POSTGRES_DSN"`
	AccessSecret        string `mapstructure:"ACCESS_SECRET"`
	MonobankToken       string `mapstructure:"MONOBANK_TOKEN"`
	MonobankRedirectUrl string `mapstructure:"MONOBANK_REDIRECT_URL"`
	MonobankWebHookUrl  string `mapstructure:"MONOBANK_WEBHOOK_URL"`
}

func NewConfig() *Config {
	var cfg Config
	viper.AutomaticEnv()
	executablePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executableDir := filepath.Dir(executablePath)
	viper.SetConfigFile(filepath.Join(executableDir, ".env"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
