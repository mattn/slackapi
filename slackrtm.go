package main

import (
	"io"
	"io/ioutil"
)

type SlackRtm struct {
	unknownEventHandlers        []func(Event, []byte)
	helloEventHandlers          []func(HelloEvent)
	presenceChangeEventHandlers []func(PresenceChangeEvent)
	messageEventHandlers        []func(MessageEvent)
}

func NewSlackRtm() SlackRtm {
	s := SlackRtm{}

	s.unknownEventHandlers = make([]func(Event, []byte), 0)
	s.helloEventHandlers = make([]func(HelloEvent), 0)
	s.presenceChangeEventHandlers = make([]func(PresenceChangeEvent), 0)
	s.messageEventHandlers = make([]func(MessageEvent), 0)

	return s
}

func (s SlackRtm) triggerUnknownEvents(evt Event, evtstring []byte) {
	for _, handler := range s.unknownEventHandlers {
		handler(evt, evtstring)
	}
}
func (s SlackRtm) OnUnknownEvent(handler func(Event, []byte)) {
	s.unknownEventHandlers = append(s.unknownEventHandlers, handler)
}

func (s SlackRtm) triggerHelloEvents(evt HelloEvent) {
	for _, handler := range s.helloEventHandlers {
		handler(evt)
	}
}
func (s SlackRtm) OnHelloEvent(handler func(HelloEvent)) {
	s.helloEventHandlers = append(s.helloEventHandlers, handler)
}

func (s SlackRtm) triggerPresenceChangeEvents(evt PresenceChangeEvent) {
	for _, handler := range s.presenceChangeEventHandlers {
		handler(evt)
	}
}
func (s SlackRtm) OnPresenceChangeEvent(handler func(PresenceChangeEvent)) {
	s.presenceChangeEventHandlers = append(s.presenceChangeEventHandlers, handler)
}

func (s SlackRtm) triggerMessageEvents(evt MessageEvent) {
	for _, handler := range s.messageEventHandlers {
		handler(evt)
	}
}

func (s SlackRtm) OnMessageEvent(handler func(MessageEvent)) {
	s.messageEventHandlers = append(s.messageEventHandlers, handler)
}

func (s SlackRtm) parseEvent(evtReader io.Reader) error {
	evtstring, _ := ioutil.ReadAll(evtReader)

	genericEvt := parseEvent(evtstring)

	switch genericEvt.Type {
	default:
		s.triggerUnknownEvents(genericEvt, evtstring)
	case "hello":
		evt := parseHelloEvent(evtstring)
		s.triggerHelloEvents(evt)
	case "presence_change":
		evt := parsePresenceChangeEvent(evtstring)
		s.triggerPresenceChangeEvents(evt)
	case "message":
		evt := parseMessageEvent(evtstring)
		s.triggerMessageEvents(evt)
	}

	return nil
}
