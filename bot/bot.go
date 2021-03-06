package bot

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"log"
	"os"
	"reacji-github-issue/config"
	"reacji-github-issue/handler"
)

func Run(sc config.SystemConfig, rc config.ReacjiConfig) {

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
		log.Fatalf("SLACK_BOT_TOKEN is invalid: %v\n", authTestErr)
	}
	//selfUserId := authTest.UserID

	go func() {
		handleEvents(socket, sc, rc)
	}()
	socket.Run()
}

func handleEvents(client *socketmode.Client, sc config.SystemConfig, rc config.ReacjiConfig) {

	var h *handler.CommandHandler
	for envelope := range client.Events {
		switch envelope.Type {
		case socketmode.EventTypeConnected:
			log.Println("Connection Established.")
			h = handler.NewHandler(client, sc, rc)
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
