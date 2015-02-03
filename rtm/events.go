package rtm

import "github.com/pastjean/slackbot/api"

// Event is a Basic event you only know the type of the event which slack gives us.
// For more information on the events look at https://api.slack.com/rtm
type Event struct {
	Type string `json:"type"`
}

// HelloEvent happens when The client has succesfully connected to the server
// https://api.slack.com/events/hello
type HelloEvent struct {
	Event
}

// MessageEvent happens when A message was sent to a channel
// https://api.slack.com/events/message
// TODO: implement messages subtypes Messages subtypes are still not implemented
type MessageEvent struct {
	Event
	// Simple Message
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	TS      string `json:"ts"`

	// Edited Message is a simple message with  more attributes
	Edited *struct {
		User string `json:"user"`
		TS   string `json:"ts"`
	} `json:"edited,omitempty"`

	Subtype string `json:"string,omitempty"`
	Hidden  bool   `json:"hidden"`
}

// ChannelMarkedEvent happens when Your channel read marker was updated
// https://api.slack.com/events/channel_marked
type ChannelMarkedEvent struct {
	Event
	Channel string `json:"channel"`
	TS      string `json:"ts"`
}

// ChannelCreatedEvent happens when A new Channel was Created
// https://api.slack.com/events/channel_created
type ChannelCreatedEvent struct {
	Channel struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Created string `json:"created"`
		Creator string `json:"creator"`
	} `json:"channel"`
}

// ChannelJoinedEvent happens when You joined a channel
// https://api.slack.com/events/channel_joined
type ChannelJoinedEvent struct {
	Event
	Channel api.Channel `json:"channel"`
}

// ChannelLeftEvent happens when You left a channel
// https://api.slack.com/events/channel_left
type ChannelLeftEvent struct {
	Event
	Channel api.Channel `json:"channel"`
}

// ChannelDeletedEvent happens when A team channel was deleted
// https://api.slack.com/events/channel_deleted
type ChannelDeletedEvent struct {
	Event
	Channel string `json:"channel"`
}

// ChannelRenameEvent happens when A team channel was renamed
// https://api.slack.com/events/channel_rename
type ChannelRenameEvent struct {
	Event
	Channel struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Created int    `json:"created"`
	} `json:"channel"`
}

// ChannelArchiveEvent happens when A team channel was archived
// https://api.slack.com/events/channel_archive
type ChannelArchiveEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// ChannelUnarchiveEvent happens when A team channel was unarchived
// https://api.slack.com/events/channel_unarchive
type ChannelUnarchiveEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// ChannelHistoryChangedEvent indicates Bulk updates were made to a channel's history
// https://api.slack.com/events/channel_history_changed
type ChannelHistoryChangedEvent struct {
	Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}

// ImCreatedEvent indicates A direct message channel was created
// https://api.slack.com/events/im_created
type ImCreatedEvent struct {
	Event
	Channel api.Channel `json:"channel"`
	User    string      `json:"user"`
}

// ImOpenEvent indicates You opened a direct message channel
// https://api.slack.com/events/im_open
type ImOpenEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// ImCloseEvent indicates You closed a direct message channel
// https://api.slack.com/events/im_close
type ImCloseEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// ImMarkedEvent indicates A direct message read marker was updated
// https://api.slack.com/events/im_marked
type ImMarkedEvent struct {
	Event
	Channel string `json:"channel"`
	TS      string `json:"ts"`
}

// ImHistoryChangedEvent indicates Bulk updates were made to a DM channel's
// history
// https://api.slack.com/events/im_history_changed
type ImHistoryChangedEvent struct {
	Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}

// GroupJoinedEvent indicates You joined a private group
// https://api.slack.com/events/group_joined
type GroupJoinedEvent struct {
	Event
	Channel api.Channel `json:"channel"`
}

// GroupLeftEvent indicates You left a private group
// https://api.slack.com/events/group_left
type GroupLeftEvent struct {
	Event
	Channel api.Channel `json:"channel"`
}

// GroupOpenEvent indicates You opened a group channel
// https://api.slack.com/events/group_open
type GroupOpenEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// GroupCloseEvent indicates You closed a group channel
// https://api.slack.com/events/group_close
type GroupCloseEvent struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

// GroupArchiveEvent indicates A private group was archived
// https://api.slack.com/events/group_archive
type GroupArchiveEvent struct {
	Event
	Channel string `json:"channel"`
}

// GroupUnarchiveEvent indicates A private group was unarchived
// https://api.slack.com/events/group_unarchive
type GroupUnarchiveEvent struct {
	Event
	Channel string `json:"channel"`
}

// GroupRenameEvent indicates A private group was renamed
// https://api.slack.com/events/group_rename
type GroupRenameEvent struct {
	Event
	Channel struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Created int    `json:"created"`
	} `json:"channel"`
}

// GroupMarkedEvent indicates A private group read marker was updated
// https://api.slack.com/events/group_marked
type GroupMarkedEvent struct {
	Event
	Channel string `json:"channel"`
	TS      string `json:"ts"`
}

// GroupHistoryChangedEvent indicates Bulk updates were made to a group's history
// https://api.slack.com/events/group_history_changed
type GroupHistoryChangedEvent struct {
	Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}

// File Events Handlers

// FileCreatedEvent indicates A file was created
// https://api.slack.com/events/file_created
type FileCreatedEvent struct {
	Event
	File api.File `json:"file"`
}

// FileSharedEvent indicates A file was shared
// https://api.slack.com/events/file_shared
type FileSharedEvent struct {
	Event
	File api.File `json:"file"`
}

// FileUnsharedEvent indicates A file was unshared
// https://api.slack.com/events/file_unshared
type FileUnsharedEvent struct {
	Event
	File api.File `json:"file"`
}

// FilePublicEvent indicates A file was made public
// https://api.slack.com/events/file_public
type FilePublicEvent struct {
	Event
	File api.File `json:"file"`
}

// FilePrivateEvent indicates A file was made private
// https://api.slack.com/events/file_private
type FilePrivateEvent struct {
	Event
	File string `json:"file"`
}

// FileChangeEvent indicates A file was changed
// https://api.slack.com/events/file_change
type FileChangeEvent struct {
	Event
	File api.File `json:"file"`
}

// FileDeletedEvent indicates A file was deleted
// https://api.slack.com/events/file_deleted
type FileDeletedEvent struct {
	Event
	FileID  string `json:"file_id"`
	EventTS string `json:"event_ts"`
}

// FileCommentAddedEvent indicates A file comment was added
// https://api.slack.com/events/file_comment_added
type FileCommentAddedEvent struct {
	Event
	File    api.File    `json:"file"`
	Comment interface{} `json:"comment"` // TODO: Comment
}

// FileCommentEditedEvent indicates A file comment was edited
// https://api.slack.com/events/file_comment_edited
type FileCommentEditedEvent struct {
	Event
	File api.File `json:"file"`
	// TODO: Comment
	Comment interface{} `json:"comment"` // TODO: Comment
}

// FileCommentDeletedEvent indicates A file comment was deleted
// https://api.slack.com/events/file_comment_deleted
type FileCommentDeletedEvent struct {
	Event
	File    api.File `json:"file"`
	Comment string   `json:"comment"`
}

// PresenceChangeEvent A team member's presence changed
// https://api.slack.com/events/presence_change
type PresenceChangeEvent struct {
	Event
	User     string `json:"user"`
	Presence string `json:"presence"`
}

// ManualPresenceChangeEvent You manually updated your presence
// https://api.slack.com/events/manual_presence_change
type ManualPresenceChangeEvent struct {
	Event
	Presence string `json:"presence"`
}

// PrefChangeEvent indicates You have updated your preferences
// https://api.slack.com/events/pref_change
type PrefChangeEvent struct {
	Event
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// UserChangeEvent indicates A team member's data has changed
// https://api.slack.com/events/user_change
type UserChangeEvent struct {
	Event
	User api.User `json:"user"`
}

// TeamJoinEvent indicates a new team member has joined
// https://api.slack.com/events/team_join
type TeamJoinEvent struct {
	Event
	User api.User `json:"user"`
}

// StarAddedEvent indicates A team member has starred an item
// https://api.slack.com/events/star_added
// TODO:
type StarAddedEvent struct{ Event }

// StarRemovedEvent indicates A team member removed a star
// https://api.slack.com/events/star_removed
// TODO:
type StarRemovedEvent struct{ Event }

// EmojiChangedEvent indicates the lists of emoji has changed
// https://api.slack.com/events/emoji_changed
type EmojiChangedEvent struct {
	Event
	EventTS string `json:"event_ts"`
}

// CommandChangedEvent indicates a team slash command has been added or changed
// https://api.slack.com/events/commands_changed
type CommandChangedEvent struct {
	Event
	EventTS string `json:"event_ts"`
}

// TeamPrefChangeEvent indicates a team preference has been updated
// https://api.slack.com/events/team_pref_change
type TeamPrefChangeEvent struct {
	Event
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// TeamRenameEvent happens when an admin changes the team name
// https://api.slack.com/events/team_rename
type TeamRenameEvent struct {
	Event
	Name string `json:"name"`
}

// TeamDomainChangeEvent happens when the team domain has changed
// https://api.slack.com/events/team_domain_change
type TeamDomainChangeEvent struct {
	Event
	URL    string `json:"url"`
	Domain string `json:"domain"`
}

// EmailDomainChangeEvent happens when the team email domain has changed
// https://api.slack.com/events/email_domain_changed
type EmailDomainChangeEvent struct {
	Event
	EmailDomain string `json:"email_domain"`
	EventTS     string `json:"event_ts"`
}

// BotAddedEvent happens when an integration bot was added
// https://api.slack.com/events/bot_added
type BotAddedEvent struct {
	Event
	Bot struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Icons struct {
			Image48 string `json:"image_48"`
		} `json:"icons"`
	} `json:"bot"`
}

// BotChangedEvent happens when an integration bot was changed
// https://api.slack.com/events/bot_changed
type BotChangedEvent struct {
	Event
	Bot struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Icons struct {
			Image48 string `json:"image_48"`
		} `json:"icons"`
	} `json:"bot"`
}

// AccountsChangedEvent is used by slack web client to maintain a list of
// logged-in accounts. Other clients should ignore this event.
type AccountsChangedEvent struct{ Event }

// TeamMigrationStartedEvent happens when the team is being migrated between
// servers
// https://api.slack.com/events/team_migration_started
type TeamMigrationStartedEvent struct{ Event }
