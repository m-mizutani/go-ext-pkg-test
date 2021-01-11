package iface_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/m-mizutani/go-ext-pkg-test/iface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dummyHTTPClient struct {
	requests []*http.Request
}

func (x *dummyHTTPClient) Do(req *http.Request) (*http.Response, error) {
	x.requests = append(x.requests, req)

	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(strings.NewReader(sampleResp)),
	}, nil
}

const sampleResp = `{
    "ok": true,
    "channel": "C1H9RESGL",
    "ts": "1503435956.000247",
    "message": {
        "text": "Here's a message for you",
        "username": "ecto1",
        "bot_id": "B19LU7CSY",
        "attachments": [
            {
                "text": "This is an attachment",
                "id": 1,
                "fallback": "This is an attachment's fallback"
            }
        ],
        "type": "message",
        "subtype": "bot_message",
        "ts": "1503435956.000247"
    }
}`

func TestSlack(t *testing.T) {
	var dummy dummyHTTPClient
	err := iface.SendMessage("MY_TOKEN", "test-message", &dummy)
	require.NoError(t, err)
	require.Equal(t, 1, len(dummy.requests))
	req := dummy.requests[0]
	assert.Equal(t, "POST", req.Method)
	assert.Equal(t, "slack.com", req.Host)

	raw, err := ioutil.ReadAll(req.Body)
	require.NoError(t, err)
	body := string(raw)
	assert.Contains(t, body, `channel=%23test`)   // Channel
	assert.Contains(t, body, `text=test-message`) // Text
	assert.Contains(t, body, `token=MY_TOKEN`)    // Token
}
