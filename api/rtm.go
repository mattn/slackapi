package api

import "net/url"

type RtmStartResponse struct {
	Response
	URL string `json:"url"`
	//Self     User      `json:"self"`
	Users    []User    `json:"users"`
	Channels []Channel `json:"channels"`

	// Groups
	// Ims
	// Team
	// Bots
}

// RtmStart returns all the info needed to start a realtime (websocket)
// connection to slack
// https://api.slack.com/methods/rtm.start
func RtmStart(token string) (*RtmStartResponse, error) {
	rtmStartResponse := &RtmStartResponse{}
	err := getJSON(
		"https://slack.com/api/rtm.start",
		url.Values{"token": {token}},
		rtmStartResponse)

	return rtmStartResponse, err
}
