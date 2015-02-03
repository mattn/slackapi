package api

import "net/url"

// emoji.list
type EmojiResponse struct {
	Response
	Emojis EmojiMap `json:"emoji"`
}

type EmojiMap map[string]string

func emojiList(token string) (*EmojiResponse, error) {
	emojilist := &EmojiResponse{}
	err := getJSON(
		"https://slack.com/api/emoji.list",
		url.Values{"token": {token}},
		emojilist)

	return emojilist, err
}
