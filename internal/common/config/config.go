package config

import (
	"github.com/spf13/viper"
)

const (
	defaultMode  = "prod"
	defaultPort  = 8080
	defaultProto = "http"
)

type Config struct {
	Mode                         string
	Port                         int
	Proto                        string
	ApiKey                       string
	GcpProjectId                 string
	GoogleApplicationCredentials string
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
	mode := viper.GetString("MODE")
	if mode == "" {
		mode = defaultMode
	}
	port := viper.GetInt("PORT")
	if port == 0 {
		port = defaultPort
	}
	proto := viper.GetString("PROTO")
	if proto == "" {
		proto = defaultProto
	}
	return &Config{
		Mode:                         mode,
		Port:                         port,
		Proto:                        proto,
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
