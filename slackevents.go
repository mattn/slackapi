package main

import "encoding/json"

type Event struct {
	Type string
}

type HelloEvent struct {
	Type string
}

type PresenceChangeEvent struct {
	Event
	User     string
	Presence string
}

type MessageEventBase struct {
	Event
	Channel string
	User    string
	Text    string
	Ts      string
	Team    string
}

type MessageEvent struct {
	MessageEventBase
	Message MessageEditedEvent
	Subtype string
	Hidden  bool
}

type MessageEditedEvent struct {
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
