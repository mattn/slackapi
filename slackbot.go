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

// Start() starts the listening loop of the websocket, bot will parse Events
// and trigger the binded events
func (s *SlackBot) Start() error {

	rtmResponse, err := api.GetRtmStart(s.token)

	if err != nil {
		return err
	}

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

func (s SlackBot) Token() string {
	return s.token
}

// Use this before starting the bot, this is the token given to you when you add
// the integration, it can be retrieved https://stjeanme.slack.com/services
func (s *SlackBot) SetToken(tok string) {
	s.token = tok
}

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

// If you want to always trigger a function on all events happening where
// SlackBot can be found
func (s *SlackBot) OnAllEvents(handler func(rtm.Event, []byte)) {
	s.allEventsHandlers = append(s.allEventsHandlers, handler)
}

func (s SlackBot) triggerUnknownEvents(evt rtm.Event, evtstring []byte) {
	for _, handler := range s.unknownEventHandlers {
		handler(evt, evtstring)
	}
}

// Triggered on UnknownEvent
func (s *SlackBot) OnUnknownEvents(handler func(rtm.Event, []byte)) {
	s.unknownEventHandlers = append(s.unknownEventHandlers, handler)
}

func (s SlackBot) triggerHelloEvents(evt rtm.HelloEvent) {
	for _, handler := range s.helloEventHandlers {
		handler(evt)
	}
}

// Triggered on HelloEvent
func (s *SlackBot) OnHelloEvents(handler func(rtm.HelloEvent)) {
	s.helloEventHandlers = append(s.helloEventHandlers, handler)
}

func (s SlackBot) triggerPresenceChangeEvents(evt rtm.PresenceChangeEvent) {
	for _, handler := range s.presenceChangeEventHandlers {
		handler(evt)
	}
}

// Triggered on PresenceChangeEvent
func (s *SlackBot) OnPresenceChangeEvents(handler func(rtm.PresenceChangeEvent)) {
	s.presenceChangeEventHandlers = append(s.presenceChangeEventHandlers, handler)
}

func (s SlackBot) triggerMessageEvents(evt rtm.MessageEvent) {
	for _, handler := range s.messageEventHandlers {
		handler(evt)
	}
}

// Triggered on all types of MessageEvent
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
