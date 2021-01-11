package iface

import (
	"net/http"

	"github.com/slack-go/slack"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func SendMessage(token string, text string, client httpClient) error {
	slackClient := slack.New(token, slack.OptionHTTPClient(client))

	if _, _, err := slackClient.PostMessage("#test", slack.MsgOptionText(text, false)); err != nil {
		return err
	}

	return nil
}
