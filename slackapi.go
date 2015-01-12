package slack

import (
	"net/http"
	"net/url"
)

// func ChannelsJoin(channel string, token string) (*http.Response, error) {
// 	return http.PostForm("https://slack.com/api/channels.join",
// 		url.Values{"channel": {channel}, "token": {token}})
// }

func ApiTest() {}

func AuthTest(token string) (*http.Response, error) {
	return http.PostForm("https://slack.com/api/auth.test",
		url.Values{"token": {token}})
}
