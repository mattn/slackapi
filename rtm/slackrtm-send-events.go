package rtm

type MessageSendEvent struct {
	Event
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}
