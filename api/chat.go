package api

import (
	"net/http"
	"net/url"
)

// chat.delete
// chat.postMessage

// ChatPostMessage posts a message to a slack channel
func ChatPostMessage(channel string, text string, token string) (*http.Response, error) {
	return http.PostForm(
		"chat.postMessage",
		url.Values{
			"token":   {token},
			"channel": {channel},
			"text":    {text},
		})
}

// chat.update
