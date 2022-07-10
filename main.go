package main

import (
	"context"
	"github.com/BurntSushi/toml"
	"log"

	"slack-reacji-issue/bot"
	"slack-reacji-issue/config"

	"github.com/sethvargo/go-envconfig"
)

func main() {

	ctx := context.Background()
	var sc config.SystemConfig
	if scErr := envconfig.Process(ctx, &sc); scErr != nil {
		log.Fatal(scErr)
	}

	var rc config.ReacjiConfig
	if _, rcErr := toml.DecodeFile("reacji.toml", &rc); rcErr != nil {
		log.Fatal(rcErr)
	}

	bot.Run(sc, rc)
}
