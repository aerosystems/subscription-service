package main

import (
	"github.com/spf13/viper"
)

const (
	defaultMode  = "prod"
	defaultPort  = "8080"
	defaultProto = "http"
)

type Config struct {
	Mode                         string
	Host                         string
	Port                         string
	Proto                        string
	GcpProjectId                 string
	GoogleApplicationCredentials string
	ProjectServiceGRPCAddr       string
	ApiKey                       string
	MonobankToken                string
	MonobankRedirectUrl          string
	MonobankWebHookUrl           string
	ProjectTopicId               string
	ProjectSubName               string
	ProjectCreateEndpoint        string
	ProjectServiceApiKey         string
}

func NewConfig() *Config {
	viper.AutomaticEnv()
	viper.SetDefault("MODE", defaultMode)
	viper.SetDefault("PORT", defaultPort)
	viper.SetDefault("PROTO", defaultProto)
	return &Config{
		Mode:                         viper.GetString("MODE"),
		Host:                         viper.GetString("HOST"),
		Port:                         viper.GetString("PORT"),
		Proto:                        viper.GetString("PROTO"),
		GcpProjectId:                 viper.GetString("GCP_PROJECT_ID"),
		GoogleApplicationCredentials: viper.GetString("GOOGLE_APPLICATION_CREDENTIALS"),
		ApiKey:                       viper.GetString("SBS_API_KEY"),
		MonobankToken:                viper.GetString("SBS_MONOBANK_TOKEN"),
		MonobankRedirectUrl:          viper.GetString("SBS_MONOBANK_REDIRECT_URL"),
		MonobankWebHookUrl:           viper.GetString("SBS_MONOBANK_WEBHOOK_URL"),
		ProjectTopicId:               viper.GetString("SBS_PROJECT_TOPIC_ID"),
		ProjectSubName:               viper.GetString("SBS_PROJECT_SUB_NAME"),
		ProjectCreateEndpoint:        viper.GetString("SBS_PROJECT_CREATE_ENDPOINT"),
		ProjectServiceApiKey:         viper.GetString("PRJCT_API_KEY"),
	}
}
