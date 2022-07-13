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

type MessageDetail struct {
	Message   string
	Permalink string
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

func (ss SlackService) GetMessageDetail(event *slackevents.ReactionAddedEvent) (*MessageDetail, error) {

	message, messageErr := ss.client.GetConversationHistory(&slack.GetConversationHistoryParameters{
		// see https://api.slack.com/methods/conversations.history#retrieving-a-single-message
		ChannelID: event.Item.Channel, Oldest: event.Item.Timestamp, Inclusive: true, Limit: 1})
	if messageErr != nil {
		log.Printf("error occured on getting message details. %v", messageErr)
		return nil, messageErr
	} else if len(message.Messages) == 0 {
		log.Printf("failed to get message details.")
		return nil, nil
	}

	permalink, linkErr := ss.client.GetPermalink(&slack.PermalinkParameters{Channel: event.Item.Channel, Ts: event.Item.Timestamp})
	if linkErr != nil {
		log.Printf("error occured on getting permalink. %v", linkErr)
		return nil, messageErr
	}

	return &MessageDetail{Message: message.Messages[0].Text, Permalink: permalink}, nil
}

func (ss SlackService) ReplyCreatedIssueUrl(event *slackevents.ReactionAddedEvent, url string) error {
	_, _, err := ss.client.PostMessage(event.Item.Channel, slack.MsgOptionText(url, false), slack.MsgOptionTS(event.Item.Timestamp))
	if err != nil {
		log.Printf("falied to post reply for message. %s", err)
		return err
	}
	return nil
}
