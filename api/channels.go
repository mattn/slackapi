package api

import "net/url"

// channels.archive
// channels.create
// channels.history
// channels.info
// channels.invite
// channels.join
// channels.kick
// channels.leave
// channels.list

// GetChannelList returns the slack domain channel list
func GetChannelList(token string) (*ChannelListResponse, error) {
	channelList := &ChannelListResponse{}
	err := getJSON(
		"https://slack.com/api/channels.list",
		url.Values{"token": {token}},
		channelList)

	return channelList, err
}

// channels.mark
// channels.rename
// channels.setPurpose
// channels.setTopic
// channels.unarchive
