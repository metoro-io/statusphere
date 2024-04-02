package slack_webhook

import (
	"github.com/metoro-io/statusphere/common/api"
	"github.com/riverqueue/river"
)

type SlackWebhookArgs struct {
	WebhookUrl string       `json:"webhook_url"`
	Incident   api.Incident `json:"incident"`
}

func (SlackWebhookArgs) Kind() string {
	return "slack_webhook"
}

func (SlackWebhookArgs) InsertOpts() river.InsertOpts {
	return river.InsertOpts{
		MaxAttempts: 10,
	}
}
