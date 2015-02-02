package api

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func get(url string, params url.Values) (*http.Response, error) {
	return http.Get(url + "?" + params.Encode())
}
func getJSON(url string, params url.Values, v interface{}) error {
	resp, err := get(url, params)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// AuthTest tests if our auth token is good
func AuthTest(token string) (*http.Response, error) {
	return http.PostForm("https://slack.com/api/auth.test",
		url.Values{"token": {token}})
}

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

// GetRtmStart returns all the info needed to start a realtime (websocket)
// connection to slack
func GetRtmStart(token string) (*RtmStartResponse, error) {
	rtmStartResponse := &RtmStartResponse{}
	err := getJSON(
		"https://slack.com/api/rtm.start",
		url.Values{"token": {token}},
		rtmStartResponse)

	return rtmStartResponse, err
}

// GetChannelList returns the slack domain channel list
func GetChannelList(token string) (*ChannelListResponse, error) {
	channelList := &ChannelListResponse{}
	err := getJSON(
		"https://slack.com/api/channels.list",
		url.Values{"token": {token}},
		channelList)

	return channelList, err
}

// GetUserList returns the slack domain user list
func GetUserList(token string) (*UserListResponse, error) {
	userList := &UserListResponse{}
	err := getJSON(
		"https://slack.com/api/users.list",
		url.Values{"token": {token}},
		userList)

	return userList, err
}
