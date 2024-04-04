package twitter_post

import (
	"github.com/metoro-io/statusphere/common/api"
	"github.com/riverqueue/river"
)

type TwitterPostArgs struct {
	Incident   api.Incident `json:"incident"`
	WebhookUrl string       `json:"webhook_url"`
}

func (TwitterPostArgs) Kind() string {
	return "twitter_post"
}

func (TwitterPostArgs) InsertOpts() river.InsertOpts {
	return river.InsertOpts{
		MaxAttempts: 10,
	}
}
