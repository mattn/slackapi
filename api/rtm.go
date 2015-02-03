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

// rtm.start

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
