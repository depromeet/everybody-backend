package noti

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/depromeet/everybody-backend/rest-api/config"
	"github.com/pkg/errors"
)

type SlackNotifierAdapter struct {
}

func NewSlackNotifierAdapter() NotifierPort {
	return &SlackNotifierAdapter{}
}

func (a *SlackNotifierAdapter) Send(sender string, body *Message) error {

	color := "#5555CC"
	payload := slack.Payload{
		Username:  config.Config.Feedback.Slack.Username,
		Channel:   config.Config.Feedback.Slack.Channel,
		IconEmoji: config.Config.Feedback.Slack.IconEmoji,
		Attachments: []slack.Attachment{slack.Attachment{
			//Fallback:     nil,
			Color:   &color,
			PreText: nil,
			//AuthorName:   nil,
			//AuthorLink:   nil,
			//AuthorIcon:   nil,
			Title: &body.Title,
			//TitleLink:    nil,
			Text: &body.Content,
			//ImageUrl:     nil,
			Fields: []*slack.Field{
				{
					Title: "유저",
					Value: sender,
					Short: false,
				},
			},
			//Footer:       nil,
			//FooterIcon:   nil,
			//Timestamp:    nil,
			//MarkdownIn:   nil,
			//Actions:      nil,
			//CallbackID:   nil,
			//ThumbnailUrl: nil,
		}},
	}
	err := slack.Send(config.Config.Feedback.Slack.Webhook, "", payload)
	if len(err) != 0 {
		return errors.WithStack(err[0])
	}

	return nil
}
