package config

import "github.com/kelseyhightower/envconfig"

// Config aaa
type Config struct {
	TelegramBotToken string `envconfig:"TELEGRAM_API_TOKEN" required:"true"`
	WebhookURL       string `envconfig:"WEBHOOK_SERVICE_URL" required:"true"`
	Port             string `envconfig:"PORT" required:"true"`
	Db               string `envconfig:"MONGO_DB_URI" required:"true"`
	UseWebhook       bool   `envconfig:"USE_WEBHOOK" required:"true"`
	Debug            bool   `envconfig:"DEBUG"`
}

// Setup aaa
func Setup() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return &c, err
	}
	return &c, nil
}
