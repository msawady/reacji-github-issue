package ext_service

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"log"
)

type SlackService struct {
	client *socketmode.Client
}

func NewSlackService(client *socketmode.Client) *SlackService {
	return &SlackService{client}
}

func (ss SlackService) GetReactionCountFor(event *slackevents.ReactionAddedEvent) (int, error) {

	reactions, err := ss.client.GetReactions(slack.NewRefToMessage(event.Item.Channel, event.Item.Timestamp), slack.NewGetReactionsParameters())
	if err != nil {
		log.Printf("failed to get reactions. %v", err)
		return 0, err
	}

	for _, r := range reactions {
		if r.Name != event.Reaction {
			continue
		}
		return r.Count, nil
	}

	log.Printf("failed to get reaction count for %s", event.Reaction)
	return 0, nil
}

func (ss SlackService) GetPermalink(event *slackevents.ReactionAddedEvent) (string, error) {

	permalink, err := ss.client.GetPermalink(&slack.PermalinkParameters{Channel: event.Item.Channel, Ts: event.Item.Timestamp})
	if err != nil {
		log.Printf("failed to get permalink. %v", err)
		return "", err
	}

	return permalink, nil
}

func (ss SlackService) ReplyCreatedIssueUrl(event *slackevents.ReactionAddedEvent, url string) error {
	_, _, err := ss.client.PostMessage(event.Item.Channel, slack.MsgOptionText(url, false), slack.MsgOptionTS(event.Item.Timestamp))
	if err != nil {
		log.Printf("falied to post reply for message. %s", err)
		return err
	}
	return nil
}
