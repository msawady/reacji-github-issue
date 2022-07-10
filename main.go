package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type SystemConfig struct {
	SlackHost  string `env:"SLACK_HOST,required"`
	SlackToken string `env:"SLACK_TOKEN,required"`

	GitHubToken string `env:"GITHUB_TOKEN,required"`
}

func main() {

	ctx := context.Background()
	fmt.Println("hogehoge")
	var c SystemConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}
