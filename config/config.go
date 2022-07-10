package config

type SystemConfig struct {
	SlackHost  string `env:"SLACK_HOST"`
	SlackToken string `env:"SLACK_TOKEN"`

	GitHubToken string `env:"GITHUB_TOKEN"`
}
