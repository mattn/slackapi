// Package slackbot serves the purpose to create bots for slack
//
// Here is an example integration:
//
//   package main
//
//   import (
//   	"fmt"
//   	"github.com/pastjean/slackbot"
//   	"log"
//   	"os"
//   )
//
//   func main() {
//     s := slackbot.NewSlackBot()
//
//     s.OnMessageEvents(func(evt slackbot.MessageEvent) {
//       fmt.Printf("%v\n", evt)
//     })
//
//     token := os.Getenv("SLACK_TOKEN")
//     if token == "" {
// 	   	 log.Fatal("SLACK_TOKEN environment variable should be set")
// 	   }
//
//     s.SetToken(token)
//     log.Fatal(s.Start())
//   }
package slackbot

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/gorilla/websocket"
	"github.com/pastjean/slackbot/api"
	"github.com/pastjean/slackbot/rtm"
)

// SlackBot is a statefull instance, you can have multiples SlackBot in the same process
// independant of eachother
type SlackBot struct {
	allEventsHandlers           []func(rtm.Event, []byte)
	unknownEventHandlers        []func(rtm.Event, []byte)
	helloEventHandlers          []func(rtm.HelloEvent)
	presenceChangeEventHandlers []func(rtm.PresenceChangeEvent)
	messageEventHandlers        []func(rtm.MessageEvent)
	onConnectEventHandlers      []func()

	token string
	conn  *websocket.Conn

	started bool

	RtmStartResponse *api.RtmStartResponse

	messageID int
}

// NewSlackBot Instanciates
func NewSlackBot() *SlackBot {
	s := SlackBot{}
	s.started = false
	s.unknownEventHandlers = make([]func(rtm.Event, []byte), 0)
	s.helloEventHandlers = make([]func(rtm.HelloEvent), 0)
	s.presenceChangeEventHandlers = make([]func(rtm.PresenceChangeEvent), 0)
	s.messageEventHandlers = make([]func(rtm.MessageEvent), 0)
	s.messageID = 1

	return &s
}

// Start starts the listening loop of the websocket, bot will parse Events
// and trigger the binded events
func (s *SlackBot) Start() error {

	rtmResponse, err := api.GetRtmStart(s.token)
	s.RtmStartResponse = rtmResponse

	if err != nil {
		return err
	}

	fmt.Printf("URL:%v\n", rtmResponse.Url)
	conn, err := rtm.Dial(rtmResponse.Url)

	if err != nil {
		return err
	}
	s.conn = conn

	s.started = true

	s.triggerOnConnectEvent()
	err = s.runLoop()
	s.started = false
	return err

}

func (s SlackBot) sendTextMessage(msg []byte) error {
	return s.conn.WriteMessage(websocket.TextMessage, msg)
}

func (s SlackBot) runLoop() error {
	c := s.conn
	for {
		_, r, err := c.NextReader()

		if err != nil {
			return err
		}

		err = s.parseEvent(r)
		if err != nil {
			return err
		}

	}
}

// Token returns the api token
func (s SlackBot) Token() string {
	return s.token
}

// SetToken is used to set the api token, use this before starting the bot
// https://{domain}.slack.com/services
func (s *SlackBot) SetToken(tok string) {
	s.token = tok
}

// OnConnectEvent sets up a callback that will happen when the bot successfully
// connects to the api
func (s *SlackBot) OnConnectEvent(handler func()) {
	s.onConnectEventHandlers = append(s.onConnectEventHandlers, handler)
}

func (s SlackBot) triggerOnConnectEvent() {
	for _, handler := range s.onConnectEventHandlers {
		handler()
	}

}

func (s SlackBot) triggerAllEvents(evt rtm.Event, evtstring []byte) {
	for _, handler := range s.allEventsHandlers {
		handler(evt, evtstring)
	}
}

// OnAllEvents sets up a callback func that will happen on all events the user
// receives
func (s *SlackBot) OnAllEvents(handler func(rtm.Event, []byte)) {
	s.allEventsHandlers = append(s.allEventsHandlers, handler)
}

func (s SlackBot) triggerUnknownEvents(evt rtm.Event, evtstring []byte) {
	for _, handler := range s.unknownEventHandlers {
		handler(evt, evtstring)
	}
}

// OnUnknownEvents sets up a callback func that will be called when this api
// receives an event unknown to its api
func (s *SlackBot) OnUnknownEvents(handler func(rtm.Event, []byte)) {
	s.unknownEventHandlers = append(s.unknownEventHandlers, handler)
}

func (s SlackBot) triggerHelloEvents(evt rtm.HelloEvent) {
	for _, handler := range s.helloEventHandlers {
		handler(evt)
	}
}

// OnHelloEvents sets up a callback that happens on the hello event
func (s *SlackBot) OnHelloEvents(handler func(rtm.HelloEvent)) {
	s.helloEventHandlers = append(s.helloEventHandlers, handler)
}

func (s SlackBot) triggerPresenceChangeEvents(evt rtm.PresenceChangeEvent) {
	for _, handler := range s.presenceChangeEventHandlers {
		handler(evt)
	}
}

// OnPresenceChangeEvents sets up an event that happens on PresenceChangeEvent
func (s *SlackBot) OnPresenceChangeEvents(handler func(rtm.PresenceChangeEvent)) {
	s.presenceChangeEventHandlers = append(s.presenceChangeEventHandlers, handler)
}

func (s SlackBot) triggerMessageEvents(evt rtm.MessageEvent) {
	for _, handler := range s.messageEventHandlers {
		handler(evt)
	}
}

// OnMessageEvents sets up a callback that happens on all MessageEvent type events
func (s *SlackBot) OnMessageEvents(handler func(rtm.MessageEvent)) {
	s.messageEventHandlers = append(s.messageEventHandlers, handler)
}

func (s SlackBot) parseEvent(evtReader io.Reader) error {

	evtString, err := ioutil.ReadAll(evtReader)
	if err != nil {
		return err
	}
	genericEvt, err := rtm.ParseEvent(evtString)
	s.triggerAllEvents(genericEvt, evtString)

	switch genericEvt.Type {
	default:
		s.triggerUnknownEvents(genericEvt, evtString)
	case "hello":
		evt, err := rtm.ParseHelloEvent(evtString)
		if err != nil {
			return err
		}
		s.triggerHelloEvents(evt)
	case "presence_change":
		evt, err := rtm.ParsePresenceChangeEvent(evtString)
		if err != nil {
			return err
		}
		s.triggerPresenceChangeEvents(evt)
	case "message":
		evt, err := rtm.ParseMessageEvent(evtString)
		if err != nil {
			return err
		}
		s.triggerMessageEvents(evt)
	}

	return err
}

// SendMessage Uses the rtm api to Send a messages channel must be the channel
// ID and not the channel name
func (s *SlackBot) SendMessage(channel string, message string) error {
	evt := rtm.MessageSendEvent{Id: s.messageID, Channel: channel, Text: message}
	evt.Type = "message"

	err := s.conn.WriteJSON(evt)

	if err != nil {
		return err
	}
	s.messageID = s.messageID + 1

	return nil
}
