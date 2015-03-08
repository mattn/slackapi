package rtm

import (
	"github.com/gorilla/websocket"
	"github.com/pastjean/slackapi/api"
)

// SlackRTM needs to be documented
type SlackRTM struct {
	EventController
	conn             *websocket.Conn
	RtmStartResponse *api.RtmStartResponse

	connectEventHandlers []func()

	// The ID we use to send messages (must change everytime we send a msg)
	messageID int
}

func NewSlackRTM(resp *api.RtmStartResponse) *SlackRTM {
	return &SlackRTM{messageID: 1, RtmStartResponse: resp}
}

func (rtm *SlackRTM) Start() error {
	conn, err := dialWebSocket(rtm.RtmStartResponse.URL)
	if err != nil {
		return err
	}

	rtm.conn = conn
	err = rtm.runLoop()
	return err
}

func (rtm *SlackRTM) runLoop() error {
	c := rtm.conn
	for {
		_, r, err := c.NextReader()

		if err != nil {
			return err
		}
		err = rtm.ReceiveEvent(r)
		if err != nil {
			return err
		}
	}
}

func (rtm *SlackRTM) SendMessage(channel string, message string) error {
	msg := struct {
		ID      int    `json:"id"`
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Text    string `json:"message"`
	}{ID: rtm.messageID, Type: "message", Channel: channel, Text: message}

	err := rtm.conn.WriteJSON(msg)

	if err != nil {
		return err
	}
	rtm.messageID++

	return nil
}

func (rtm *SlackRTM) OnConnectEvent(handler func()) {
	if rtm.connectEventHandlers == nil {
		rtm.connectEventHandlers = make([]func(), 0)
	}
	rtm.connectEventHandlers = append(rtm.connectEventHandlers, handler)
}
func (rtm *SlackRTM) triggerConnectEventHandlers() {
	for _, handler := range rtm.connectEventHandlers {
		handler()
	}
}
