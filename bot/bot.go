package bot

import (
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"log"
	"os"
	"slack-reacji-issue/config"
	handler2 "slack-reacji-issue/handler"
)

func Run(sc config.SystemConfig, rc config.ReacjiConfig) {

	log.Printf("bot_token: %s", sc.SlackBotToken)
	log.Printf("app_token: %s", sc.SlackAppToken)
	api := slack.New(
		sc.SlackBotToken,
		slack.OptionAppLevelToken(sc.SlackAppToken),
		slack.OptionDebug(false),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	socket := socketmode.New(
		api,
		socketmode.OptionDebug(true),
		socketmode.OptionLog(log.New(os.Stdout, "sm: ", log.Lshortfile|log.LstdFlags)),
	)
	_, authTestErr := api.AuthTest()
	if authTestErr != nil {
		fmt.Fprintf(os.Stderr, "SLACK_BOT_TOKEN is invalid: %v\n", authTestErr)
		os.Exit(1)
	}
	//selfUserId := authTest.UserID

	go func() {
		handleCommand(socket, rc)
	}()
	socket.Run()

}

func handleCommand(client *socketmode.Client, rc config.ReacjiConfig) {

	var h *handler2.CommandHandler
	for envelope := range client.Events {
		switch envelope.Type {
		case socketmode.EventTypeConnected:
			log.Println("Connection Established.")
			h = handler2.NewHandler(client, rc)
		case socketmode.EventTypeEventsAPI:
			eventsAPIEvent, _ := envelope.Data.(slackevents.EventsAPIEvent)
			client.Ack(*envelope.Request)
			switch eventsAPIEvent.Type {
			case slackevents.CallbackEvent:
				innerEvent := eventsAPIEvent.InnerEvent
				switch ev := innerEvent.Data.(type) {
				case *slackevents.ReactionAddedEvent:
					h.HandleReaction(ev)
				}
			}
		default:
		}
	}
}
