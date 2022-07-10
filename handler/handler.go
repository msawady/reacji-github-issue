package handler

import (
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"log"
	"slack-reacji-issue/config"
	"slack-reacji-issue/ext_service"
)

type CommandHandler struct {
	slackService   *ext_service.SlackService
	rc             config.ReacjiConfig
	reacjiSettings map[string]config.ReacjiSetting
}

func NewHandler(client *socketmode.Client, reacjiConfig config.ReacjiConfig) *CommandHandler {

	slackService := ext_service.NewSlackService(client)

	var settings map[string]config.ReacjiSetting
	settings = make(map[string]config.ReacjiSetting)
	for _, s := range reacjiConfig.Settings {
		settings[s.Emoji] = s
	}

	log.Printf("will handling emojis: %v", settings)

	return &CommandHandler{slackService, reacjiConfig, settings}
}

func (handler *CommandHandler) HandleReaction(ev *slackevents.ReactionAddedEvent) {
	log.Printf("Handling reaction: %s from %s", ev.Reaction, ev.User)
	if v, ok := handler.reacjiSettings[ev.Reaction]; ok {
		log.Printf("will process %s.", v.Emoji)
		count, _ := handler.slackService.GetReactionCountFor(ev)
		log.Printf("reaction count %d.", count)
	} else {
		log.Printf("skipping %s.", ev.Reaction)
	}
}
