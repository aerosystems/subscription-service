package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Mode                      string `mapstructure:"MODE" required:"true"`
	PostgresDSN               string `mapstructure:"POSTGRES_DSN" required:"true"`
	AccessSecret              string `mapstructure:"ACCESS_SECRET" required:"true"`
	MonobankToken             string `mapstructure:"MONOBANK_TOKEN" required:"true"`
	MonobankRedirectUrl       string `mapstructure:"MONOBANK_REDIRECT_URL" required:"true"`
	MonobankWebHookUrl        string `mapstructure:"MONOBANK_WEBHOOK_URL" required:"true"`
	GcpProjectId              string `mapstructure:"GCP_PROJECT_ID" required:"true"`
	GcpServiceAccountFilePath string `mapstructure:"GCP_SERVICE_ACCOUNT_FILE_PATH" required:"true"`
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
