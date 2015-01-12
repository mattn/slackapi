package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"net/http"
)

type RtmStartAnswer struct {
	Ok       bool
	Url      string
	Users    []User
	Channels []Channel
}

type Channel struct {
	Id          string
	Name        string
	IsChannel   bool `json:"is_channel"`
	Creator     string
	IsArchived  bool   `json:"is_archived"`
	IsGeneral   bool   `json:"is_general"`
	IsMember    bool   `json:"is_member"`
	UnreadCount int    `json:"unread_cout"`
	LastRead    string `json:"last_read"`
	// Members     []string
	Topic   ChannelTopic
	Purpose ChannelTopic
}

type ChannelTopic struct {
	Value   string
	Creator string
}

type User struct {
	Id       string
	Name     string
	Deleted  bool
	Status   string
	Color    string
	RealName string `json:"real_name"`
}

func parseWSEvent(evtReader io.Reader) error {
	return nil
}

func startWebSocketSession(token string) (RtmStartAnswer, error) {
	endpoint := "https://slack.com/api/rtm.start"

	rtmStartString := endpoint + "?token=" + token
	resp, err := http.Get(rtmStartString)

	if err != nil {
		return RtmStartAnswer{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var ans RtmStartAnswer

	json.Unmarshal(body, &ans)

	return ans, nil
}

func dialWebSocket(url string) (*websocket.Conn, error) {
	var dialer *websocket.Dialer
	var dialHeader http.Header

	conn, _, err := dialer.Dial(url, dialHeader)

	return conn, err
}

func startSlackRtm(token string) (*websocket.Conn, error) {
	ans, err := startWebSocketSession(token)
	if err != nil {
		return nil, err
	}

	conn, err := dialWebSocket(ans.Url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
