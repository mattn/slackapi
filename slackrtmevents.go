package slackbot

import "encoding/json"

// Basic event you only know the type of the event which slack gives us.
// For more information on the events look at https://api.slack.com/rtm
type Event struct {
	Type string
}

// Hello Event doesnt have more information than the basic event, this one only
// tells us we can make call & that the server responds (simple ping)
// https://api.slack.com/events/hello
type HelloEvent struct {
	Event
}

// PresenceChange Event when a user gets
// https://api.slack.com/events/presence_change
type PresenceChangeEvent struct {
	Event
	User     string
	Presence string
}

// Both MessageEvent and MessageEditedEvent Share this structure
type MessageEventBase struct {
	Event
	Channel string
	User    string
	Text    string
	Ts      string
	Team    string
}

// A message was sent to a channel
// https://api.slack.com/events/message
type MessageEvent struct {
	MessageEventBase
	Message MessageChangeEvent
	Subtype string
	Hidden  bool
}

// This is when a MessageEvent was edited
// The basic MessageEvent will have a Subtype="message_changed"
// https://api.slack.com/events/message/message_changed
type MessageChangeEvent struct {
	MessageEventBase
	Edited EditedEvent
}

type EditedEvent struct {
	User string
	Ts   string
}

func parseEvent(evtstring []byte) Event {
	var evt Event
	json.Unmarshal(evtstring, &evt)
	return evt
}

func parseHelloEvent(evtstring []byte) HelloEvent {
	var evt HelloEvent
	json.Unmarshal(evtstring, &evt)
	return evt
}

func parsePresenceChangeEvent(evtstring []byte) PresenceChangeEvent {
	var evt PresenceChangeEvent
	json.Unmarshal(evtstring, &evt)
	return evt
}

func parseMessageEvent(evtstring []byte) MessageEvent {
	var evt MessageEvent
	json.Unmarshal(evtstring, &evt)
	return evt
}
