package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mode                         string
	WebPort                      int
	GcpProjectId                 string
	GoogleApplicationCredentials string
	MonobankToken                string
	MonobankRedirectUrl          string
	MonobankWebHookUrl           string
}

func NewConfig() *Config {
	viper.AutomaticEnv()
	return &Config{
		Mode:                         viper.GetString("SBS_MODE"),
		WebPort:                      viper.GetInt("PORT"),
		GcpProjectId:                 viper.GetString("GCP_PROJECT_ID"),
		GoogleApplicationCredentials: viper.GetString("GOOGLE_APPLICATION_CREDENTIALS"),
		MonobankToken:                viper.GetString("SBS_MONOBANK_TOKEN"),
		MonobankRedirectUrl:          viper.GetString("SBS_MONOBANK_REDIRECT_URL"),
		MonobankWebHookUrl:           viper.GetString("SBS_MONOBANK_WEBHOOK_URL"),
	}
}
