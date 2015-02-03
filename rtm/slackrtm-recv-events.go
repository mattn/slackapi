package rtm

import "encoding/json"

// Basic event you only know the type of the event which slack gives us.
// For more information on the events look at https://api.slack.com/rtm
type Event struct {
	Id   string `json:"_id,omitempty"`
	Type string `json:"type"`
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
	User     string `json:"user"`
	Presence string `json:"presence"`
}

// Both MessageEvent and MessageEditedEvent Share this structure
type MessageEventBase struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user,omitempty"`
	Text    string `json:"text,omitempty"`
	Ts      string `json:"ts"`
	Team    string `json:"team,omitempty"`
}

// A message was sent to a channel
// https://api.slack.com/events/message
type MessageEvent struct {
	MessageEventBase
	Message *MessageChangeEvent `json:"message,omitempty"`
	Subtype string              `json:"subtype,omitempty"`
	Hidden  bool                `json:"hidden,omitempty"`
}

// This is when a MessageEvent was edited
// The basic MessageEvent will have a Subtype="message_changed"
// https://api.slack.com/events/message/message_changed
type MessageChangeEvent struct {
	MessageEventBase
	Edited *EditedEvent `json:"edited"`
}

type EditedEvent struct {
	User string `json:"user"`
	Ts   string `json:"ts"`
}

// A new Channel was Created
// https://api.slack.com/events/channel_created
type ChannelCreatedEvent struct {
	Event
	Channel struct {
		ID      string `json:id`
		Name    string `json:name`
		Created string `json:created`
		Creator string `json:creator`
	} `json:"channel"`
}

//-----------------------------------------
// TODO: ChannelRenamed
// TODO: EmojiChanged
// TODO: Users
// TODO: Files
//-----------------------------------------

func ParseEvent(evtString []byte) (Event, error) {
	var evt Event
	err := json.Unmarshal(evtString, &evt)
	return evt, err
}

func ParseHelloEvent(evtString []byte) (HelloEvent, error) {
	var evt HelloEvent
	err := json.Unmarshal(evtString, &evt)
	return evt, err
}

func ParsePresenceChangeEvent(evtString []byte) (PresenceChangeEvent, error) {
	var evt PresenceChangeEvent
	err := json.Unmarshal(evtString, &evt)
	return evt, err
}

func ParseMessageEvent(evtString []byte) (MessageEvent, error) {
	var evt MessageEvent
	err := json.Unmarshal(evtString, &evt)
	return evt, err
}

func ParseChannelCreatedEvent(evtString []byte) (ChannelCreatedEvent, error) {
	var evt ChannelCreatedEvent
	err := json.Unmarshal(evtString, &evt)
	return evt, err
}
