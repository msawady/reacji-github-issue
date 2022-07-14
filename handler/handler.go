package handler

import (
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"log"
	"reacji-github-issue/config"
	"reacji-github-issue/ext_service"
)

type CommandHandler struct {
	slackService   *ext_service.SlackService
	gitHubService  *ext_service.GitHubService
	reacjiSettings map[string]config.ReacjiSetting
}

func NewHandler(client *socketmode.Client, sc config.SystemConfig, reacjiConfig config.ReacjiConfig) *CommandHandler {

	slackService := ext_service.NewSlackService(client)

	var settings map[string]config.ReacjiSetting
	settings = make(map[string]config.ReacjiSetting)
	// create map for optimize searching.
	for _, s := range reacjiConfig.Settings {
		settings[s.Emoji] = s
	}

	log.Printf("will handling emojis: %v", settings)

	gitHubService := ext_service.NewGitHubService(sc)

	return &CommandHandler{slackService, gitHubService, settings}
}

func (handler *CommandHandler) HandleReaction(ev *slackevents.ReactionAddedEvent) {
	log.Printf("Handling reaction: %s from %s", ev.Reaction, ev.User)
	if setting, ok := handler.reacjiSettings[ev.Reaction]; ok {
		log.Printf("will process %s.", setting.Emoji)
		count, _ := handler.slackService.GetReactionCountFor(ev)
		if count == 1 {
			detail, linkErr := handler.slackService.GetMessageDetail(ev)
			if linkErr != nil {
				return
			}
			param := CreateIssueParam(&setting, detail)
			issueUrl, issueErr := handler.gitHubService.CreateIssue(param)
			if issueErr != nil {
				return
			}
			err := handler.slackService.ReplyCreatedIssueUrl(ev, *issueUrl)
			if err != nil {
				return
			}
		} else {
			log.Print("Skip processing because this is not first reaction.")
		}
	} else {
		log.Printf("skipping %s.", ev.Reaction)
	}
}
