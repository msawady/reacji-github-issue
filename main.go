package main

import (
	"context"
	"log"

	"slack-reacji-issue/bot"
	"slack-reacji-issue/config"

	"github.com/sethvargo/go-envconfig"
)

func main() {

	ctx := context.Background()
	var c config.SystemConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	bot.Run(c)
}
