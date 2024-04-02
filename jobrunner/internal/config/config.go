package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	SlackWebhookUrl string `envconfig:"SLACK_WEBHOOK_URL"`
}

func GetConfigFromEnvironment() (Config, error) {
	var config Config
	err := envconfig.Process("STATUSPHERE", &config)
	return config, err
}
