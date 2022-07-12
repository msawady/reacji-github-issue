package config

type SystemConfig struct {
	SlackHost     string `env:"SLACK_HOST"`
	SlackBotToken string `env:"SLACK_BOT_TOKEN"`
	SlackAppToken string `env:"SLACK_APP_TOKEN"`

	GitHubToken          string `env:"GITHUB_TOKEN"`
	GitHubAppID          int64  `env:"GITHUB_APP_ID"`
	GitHubInstallationID int64  `env:"GITHUB_INSTALLATION_ID"`
	// Pem file encoded with base64
	GitHubPemBinary string `env:"GITHUB_PEM_BINARY"`
}

type ReacjiConfig struct {
	Settings []ReacjiSetting `toml:"settings"`
}

type ReacjiSetting struct {
	Emoji       string `toml:"emoji"`
	Owner       string `toml:"owner"`
	Repo        string `toml:"repo"`
	Description string `toml:"description"`
}
