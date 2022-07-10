package config

type SystemConfig struct {
	SlackHost     string `env:"SLACK_HOST"`
	SlackBotToken string `env:"SLACK_BOT_TOKEN"`
	SlackAppToken string `env:"SLACK_APP_TOKEN"`

	GitHubToken string `env:"GITHUB_TOKEN"`
}

type ReacjiConfig struct {
	Settings []ReacjiSetting `toml:"settings"`
}

type ReacjiSetting struct {
	Emoji       string `toml:"emoji"`
	Description string `toml:"description"`
}
