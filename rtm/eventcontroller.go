package rtm

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type EventController struct {
	helloEventHandlers   []func(HelloEvent)
	messageEventHandlers []func(MessageEvent)

	// Channel Events
	channelMarkedEventHandlers         []func(ChannelMarkedEvent)
	channelCreatedEventHandlers        []func(ChannelCreatedEvent)
	channelJoinedEventHandlers         []func(ChannelJoinedEvent)
	channelLeftEventHandlers           []func(ChannelLeftEvent)
	channelDeletedEventHandlers        []func(ChannelDeletedEvent)
	channelRenameEventHandlers         []func(ChannelRenameEvent)
	channelArchiveEventHandlers        []func(ChannelArchiveEvent)
	channelUnarchiveEventHandlers      []func(ChannelUnarchiveEvent)
	channelHistoryChangedEventHandlers []func(ChannelHistoryChangedEvent)

	// Im Events Handlers
	imCreatedEventHandlers        []func(ImCreatedEvent)
	imOpenEventHandlers           []func(ImOpenEvent)
	imCloseEventHandlers          []func(ImCloseEvent)
	imMarkedEventHandlers         []func(ImMarkedEvent)
	imHistoryChangedEventHandlers []func(ImHistoryChangedEvent)

	// Group Events Handlers

	groupJoinedEventHandlers         []func(GroupJoinedEvent)
	groupLeftEventHandlers           []func(GroupLeftEvent)
	groupOpenEventHandlers           []func(GroupOpenEvent)
	groupCloseEventHandlers          []func(GroupCloseEvent)
	groupArchiveEventHandlers        []func(GroupArchiveEvent)
	groupUnarchiveEventHandlers      []func(GroupUnarchiveEvent)
	groupRenameEventHandlers         []func(GroupRenameEvent)
	groupMarkedEventHandlers         []func(GroupMarkedEvent)
	groupHistoryChangedEventHandlers []func(GroupHistoryChangedEvent)

	// File Events Handlers
	fileCreatedEventHandlers        []func(FileCreatedEvent)
	fileSharedEventHandlers         []func(FileSharedEvent)
	fileUnsharedEventHandlers       []func(FileUnsharedEvent)
	filePublicEventHandlers         []func(FilePublicEvent)
	filePrivateEventHandlers        []func(FilePrivateEvent)
	fileChangeEventHandlers         []func(FileChangeEvent)
	fileDeletedEventHandlers        []func(FileDeletedEvent)
	fileCommentAddedEventHandlers   []func(FileCommentAddedEvent)
	fileCommentEditedEventHandlers  []func(FileCommentEditedEvent)
	fileCommentDeletedEventHandlers []func(FileCommentDeletedEvent)

	// Presence Events Handlers
	presenceChangeEventHandlers       []func(PresenceChangeEvent)
	manualPresenceChangeEventHandlers []func(ManualPresenceChangeEvent)

	// Other Events Handlers
	prefChangeEventHandlers           []func(PrefChangeEvent)
	userChangeEventHandlers           []func(UserChangeEvent)
	teamJoinEventHandlers             []func(TeamJoinEvent)
	starAddedEventHandlers            []func(StarAddedEvent)
	starRemovedEventHandlers          []func(StarRemovedEvent)
	emojiChangedEventHandlers         []func(EmojiChangedEvent)
	commandChangedEventHandlers       []func(CommandChangedEvent)
	teamPrefChangeEventHandlers       []func(TeamPrefChangeEvent)
	teamRenameEventHandlers           []func(TeamRenameEvent)
	teamDomainChangeEventHandlers     []func(TeamDomainChangeEvent)
	emailDomainChangeEventHandlers    []func(EmailDomainChangeEvent)
	botAddedEventHandlers             []func(BotAddedEvent)
	botChangedEventHandlers           []func(BotChangedEvent)
	accountsChangedEventHandlers      []func(AccountsChangedEvent)
	teamMigrationStartedEventHandlers []func(TeamMigrationStartedEvent)

	// Meta Events
	// allEventEventHandlers    []func(Event, []byte)
	// unknownEventEventHandler []func(Event, []byte)
}

func NewEventController() *EventController {
	return &EventController{}
}

func (e *EventController) ReceiveEvent(evtReader io.Reader) error {
	evtString, err := ioutil.ReadAll(evtReader)
	if err != nil {
		return err
	}

	var genericEvt Event
	err = json.Unmarshal(evtString, &genericEvt)

	//e.triggerAllEvents(evtString)

	switch genericEvt.Type {
	default:

	case "hello":
		e.triggerHelloEventHandlers(evtString)
	case "message":
		e.triggerMessageEventHandlers(evtString)
	case "channel_marked":
		e.triggerChannelMarkedEventHandlers(evtString)
	case "channel_created":
		e.triggerChannelCreatedEventHandlers(evtString)
	case "channel_joined":
		e.triggerChannelJoinedEventHandlers(evtString)
	case "channel_left":
		e.triggerChannelLeftEventHandlers(evtString)
	case "channel_deleted":
		e.triggerChannelDeletedEventHandlers(evtString)
	case "channel_rename":
		e.triggerChannelRenameEventHandlers(evtString)
	case "channel_archive":
		e.triggerChannelArchiveEventHandlers(evtString)
	case "channel_unarchive":
		e.triggerChannelUnarchiveEventHandlers(evtString)
	case "channel_history_changed":
		e.triggerChannelHistoryChangedEventHandlers(evtString)
	case "im_created":
		e.triggerImCreatedEventHandlers(evtString)
	case "im_open":
		e.triggerImOpenEventHandlers(evtString)
	case "im_close":
		e.triggerImCloseEventHandlers(evtString)
	case "im_marked":
		e.triggerImMarkedEventHandlers(evtString)
	case "im_history_changed":
		e.triggerImHistoryChangedEventHandlers(evtString)
	case "group_joined":
		e.triggerGroupJoinedEventHandlers(evtString)
	case "group_left":
		e.triggerGroupLeftEventHandlers(evtString)
	case "group_open":
		e.triggerGroupOpenEventHandlers(evtString)
	case "group_close":
		e.triggerGroupCloseEventHandlers(evtString)
	case "group_archive":
		e.triggerGroupArchiveEventHandlers(evtString)
	case "group_unarchive":
		e.triggerGroupUnarchiveEventHandlers(evtString)
	case "group_rename":
		e.triggerGroupRenameEventHandlers(evtString)
	case "group_marked":
		e.triggerGroupMarkedEventHandlers(evtString)
	case "group_history_changed":
		e.triggerGroupHistoryChangedEventHandlers(evtString)
	case "file_created":
		e.triggerFileCreatedEventHandlers(evtString)
	case "file_shared":
		e.triggerFileSharedEventHandlers(evtString)
	case "file_unshared":
		e.triggerFileUnsharedEventHandlers(evtString)
	case "file_public":
		e.triggerFilePublicEventHandlers(evtString)
	case "file_private":
		e.triggerFilePrivateEventHandlers(evtString)
	case "file_change":
		e.triggerFileChangeEventHandlers(evtString)
	case "file_deleted":
		e.triggerFileDeletedEventHandlers(evtString)
	case "file_comment_added":
		e.triggerFileCommentAddedEventHandlers(evtString)
	case "file_comment_edited":
		e.triggerFileCommentDeletedEventHandlers(evtString)
	case "file_comment_deleted":
		e.triggerFileCommentDeletedEventHandlers(evtString)
	case "presence_change":
		e.triggerPresenceChangeEventHandlers(evtString)
	case "manual_presence_change":
		e.triggerManualPresenceChangeEventHandlers(evtString)
	case "pref_change":
		e.triggerPrefChangeEventHandlers(evtString)
	case "user_change":
		e.triggerUserChangeEventHandlers(evtString)
	case "team_join":
		e.triggerTeamJoinEventHandlers(evtString)
	case "star_added":
		e.triggerStarAddedEventHandlers(evtString)
	case "star_removed":
		e.triggerStarRemovedEventHandlers(evtString)
	case "emoji_changed":
		e.triggerEmojiChangedEventHandlers(evtString)
	case "commands_changed":
		e.triggerCommandChangedEventHandlers(evtString)
	case "team_pref_change":
		e.triggerTeamPrefChangeEventHandlers(evtString)
	case "team_rename":
		e.triggerTeamRenameEventHandlers(evtString)
	case "team_domain_change":
		e.triggerTeamDomainChangeEventHandlers(evtString)
	case "bot_added":
		e.triggerBotAddedEventHandlers(evtString)
	case "bot_changed":
		e.triggerBotChangedEventHandlers(evtString)
	case "accounts_changed":
		e.triggerAccountsChangedEventHandlers(evtString)
	case "team_migration_started":
		e.triggerTeamMigrationStartedEventHandlers(evtString)
	}
	return nil
}

func (e *EventController) OnHelloEvents(handler func(HelloEvent)) {
	if e.helloEventHandlers == nil {
		e.helloEventHandlers = make([]func(HelloEvent), 0)
	}
	e.helloEventHandlers = append(e.helloEventHandlers, handler)
}
func (e *EventController) triggerHelloEventHandlers(evtString []byte) {
	var evt HelloEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.helloEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnMessageEvents(handler func(MessageEvent)) {
	if e.messageEventHandlers == nil {
		e.messageEventHandlers = make([]func(MessageEvent), 0)
	}
	e.messageEventHandlers = append(e.messageEventHandlers, handler)
}
func (e *EventController) triggerMessageEventHandlers(evtString []byte) {
	var evt MessageEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.messageEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelMarkedEvents(handler func(ChannelMarkedEvent)) {
	if e.channelMarkedEventHandlers == nil {
		e.channelMarkedEventHandlers = make([]func(ChannelMarkedEvent), 0)
	}
	e.channelMarkedEventHandlers = append(e.channelMarkedEventHandlers, handler)
}
func (e *EventController) triggerChannelMarkedEventHandlers(evtString []byte) {
	var evt ChannelMarkedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelMarkedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelCreatedEvents(handler func(ChannelCreatedEvent)) {
	if e.channelCreatedEventHandlers == nil {
		e.channelCreatedEventHandlers = make([]func(ChannelCreatedEvent), 0)
	}
	e.channelCreatedEventHandlers = append(e.channelCreatedEventHandlers, handler)
}
func (e *EventController) triggerChannelCreatedEventHandlers(evtString []byte) {
	var evt ChannelCreatedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelCreatedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelJoinedEvents(handler func(ChannelJoinedEvent)) {
	if e.channelJoinedEventHandlers == nil {
		e.channelJoinedEventHandlers = make([]func(ChannelJoinedEvent), 0)
	}
	e.channelJoinedEventHandlers = append(e.channelJoinedEventHandlers, handler)
}
func (e *EventController) triggerChannelJoinedEventHandlers(evtString []byte) {
	var evt ChannelJoinedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelJoinedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelLeftEvents(handler func(ChannelLeftEvent)) {
	if e.channelLeftEventHandlers == nil {
		e.channelLeftEventHandlers = make([]func(ChannelLeftEvent), 0)
	}
	e.channelLeftEventHandlers = append(e.channelLeftEventHandlers, handler)
}
func (e *EventController) triggerChannelLeftEventHandlers(evtString []byte) {
	var evt ChannelLeftEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelLeftEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelDeletedEvents(handler func(ChannelDeletedEvent)) {
	if e.channelDeletedEventHandlers == nil {
		e.channelDeletedEventHandlers = make([]func(ChannelDeletedEvent), 0)
	}
	e.channelDeletedEventHandlers = append(e.channelDeletedEventHandlers, handler)
}
func (e *EventController) triggerChannelDeletedEventHandlers(evtString []byte) {
	var evt ChannelDeletedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelDeletedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelRenameEvents(handler func(ChannelRenameEvent)) {
	if e.channelRenameEventHandlers == nil {
		e.channelRenameEventHandlers = make([]func(ChannelRenameEvent), 0)
	}
	e.channelRenameEventHandlers = append(e.channelRenameEventHandlers, handler)
}
func (e *EventController) triggerChannelRenameEventHandlers(evtString []byte) {
	var evt ChannelRenameEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelRenameEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelArchiveEvents(handler func(ChannelArchiveEvent)) {
	if e.channelArchiveEventHandlers == nil {
		e.channelArchiveEventHandlers = make([]func(ChannelArchiveEvent), 0)
	}
	e.channelArchiveEventHandlers = append(e.channelArchiveEventHandlers, handler)
}
func (e *EventController) triggerChannelArchiveEventHandlers(evtString []byte) {
	var evt ChannelArchiveEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelArchiveEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelUnarchiveEvents(handler func(ChannelUnarchiveEvent)) {
	if e.channelUnarchiveEventHandlers == nil {
		e.channelUnarchiveEventHandlers = make([]func(ChannelUnarchiveEvent), 0)
	}
	e.channelUnarchiveEventHandlers = append(e.channelUnarchiveEventHandlers, handler)
}
func (e *EventController) triggerChannelUnarchiveEventHandlers(evtString []byte) {
	var evt ChannelUnarchiveEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelUnarchiveEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnChannelHistoryChangedEvents(handler func(ChannelHistoryChangedEvent)) {
	if e.channelHistoryChangedEventHandlers == nil {
		e.channelHistoryChangedEventHandlers = make([]func(ChannelHistoryChangedEvent), 0)
	}
	e.channelHistoryChangedEventHandlers = append(e.channelHistoryChangedEventHandlers, handler)
}
func (e *EventController) triggerChannelHistoryChangedEventHandlers(evtString []byte) {
	var evt ChannelHistoryChangedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.channelHistoryChangedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnImCreatedEvents(handler func(ImCreatedEvent)) {
	if e.imCreatedEventHandlers == nil {
		e.imCreatedEventHandlers = make([]func(ImCreatedEvent), 0)
	}
	e.imCreatedEventHandlers = append(e.imCreatedEventHandlers, handler)
}
func (e *EventController) triggerImCreatedEventHandlers(evtString []byte) {
	var evt ImCreatedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.imCreatedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnImOpenEvents(handler func(ImOpenEvent)) {
	if e.imOpenEventHandlers == nil {
		e.imOpenEventHandlers = make([]func(ImOpenEvent), 0)
	}
	e.imOpenEventHandlers = append(e.imOpenEventHandlers, handler)
}
func (e *EventController) triggerImOpenEventHandlers(evtString []byte) {
	var evt ImOpenEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.imOpenEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnImCloseEvents(handler func(ImCloseEvent)) {
	if e.imCloseEventHandlers == nil {
		e.imCloseEventHandlers = make([]func(ImCloseEvent), 0)
	}
	e.imCloseEventHandlers = append(e.imCloseEventHandlers, handler)
}
func (e *EventController) triggerImCloseEventHandlers(evtString []byte) {
	var evt ImCloseEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.imCloseEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnImMarkedEvents(handler func(ImMarkedEvent)) {
	if e.imMarkedEventHandlers == nil {
		e.imMarkedEventHandlers = make([]func(ImMarkedEvent), 0)
	}
	e.imMarkedEventHandlers = append(e.imMarkedEventHandlers, handler)
}
func (e *EventController) triggerImMarkedEventHandlers(evtString []byte) {
	var evt ImMarkedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.imMarkedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnImHistoryChangedEvents(handler func(ImHistoryChangedEvent)) {
	if e.imHistoryChangedEventHandlers == nil {
		e.imHistoryChangedEventHandlers = make([]func(ImHistoryChangedEvent), 0)
	}
	e.imHistoryChangedEventHandlers = append(e.imHistoryChangedEventHandlers, handler)
}
func (e *EventController) triggerImHistoryChangedEventHandlers(evtString []byte) {
	var evt ImHistoryChangedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.imHistoryChangedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupJoinedEvents(handler func(GroupJoinedEvent)) {
	if e.groupJoinedEventHandlers == nil {
		e.groupJoinedEventHandlers = make([]func(GroupJoinedEvent), 0)
	}
	e.groupJoinedEventHandlers = append(e.groupJoinedEventHandlers, handler)
}
func (e *EventController) triggerGroupJoinedEventHandlers(evtString []byte) {
	var evt GroupJoinedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupJoinedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupLeftEvents(handler func(GroupLeftEvent)) {
	if e.groupLeftEventHandlers == nil {
		e.groupLeftEventHandlers = make([]func(GroupLeftEvent), 0)
	}
	e.groupLeftEventHandlers = append(e.groupLeftEventHandlers, handler)
}
func (e *EventController) triggerGroupLeftEventHandlers(evtString []byte) {
	var evt GroupLeftEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupLeftEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupOpenEvents(handler func(GroupOpenEvent)) {
	if e.groupOpenEventHandlers == nil {
		e.groupOpenEventHandlers = make([]func(GroupOpenEvent), 0)
	}
	e.groupOpenEventHandlers = append(e.groupOpenEventHandlers, handler)
}
func (e *EventController) triggerGroupOpenEventHandlers(evtString []byte) {
	var evt GroupOpenEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupOpenEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupCloseEvents(handler func(GroupCloseEvent)) {
	if e.groupCloseEventHandlers == nil {
		e.groupCloseEventHandlers = make([]func(GroupCloseEvent), 0)
	}
	e.groupCloseEventHandlers = append(e.groupCloseEventHandlers, handler)
}
func (e *EventController) triggerGroupCloseEventHandlers(evtString []byte) {
	var evt GroupCloseEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupCloseEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupArchiveEvents(handler func(GroupArchiveEvent)) {
	if e.groupArchiveEventHandlers == nil {
		e.groupArchiveEventHandlers = make([]func(GroupArchiveEvent), 0)
	}
	e.groupArchiveEventHandlers = append(e.groupArchiveEventHandlers, handler)
}
func (e *EventController) triggerGroupArchiveEventHandlers(evtString []byte) {
	var evt GroupArchiveEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupArchiveEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupUnarchiveEvents(handler func(GroupUnarchiveEvent)) {
	if e.groupUnarchiveEventHandlers == nil {
		e.groupUnarchiveEventHandlers = make([]func(GroupUnarchiveEvent), 0)
	}
	e.groupUnarchiveEventHandlers = append(e.groupUnarchiveEventHandlers, handler)
}
func (e *EventController) triggerGroupUnarchiveEventHandlers(evtString []byte) {
	var evt GroupUnarchiveEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupUnarchiveEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupRenameEvents(handler func(GroupRenameEvent)) {
	if e.groupRenameEventHandlers == nil {
		e.groupRenameEventHandlers = make([]func(GroupRenameEvent), 0)
	}
	e.groupRenameEventHandlers = append(e.groupRenameEventHandlers, handler)
}
func (e *EventController) triggerGroupRenameEventHandlers(evtString []byte) {
	var evt GroupRenameEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupRenameEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupMarkedEvents(handler func(GroupMarkedEvent)) {
	if e.groupMarkedEventHandlers == nil {
		e.groupMarkedEventHandlers = make([]func(GroupMarkedEvent), 0)
	}
	e.groupMarkedEventHandlers = append(e.groupMarkedEventHandlers, handler)
}
func (e *EventController) triggerGroupMarkedEventHandlers(evtString []byte) {
	var evt GroupMarkedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupMarkedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnGroupHistoryChangedEvents(handler func(GroupHistoryChangedEvent)) {
	if e.groupHistoryChangedEventHandlers == nil {
		e.groupHistoryChangedEventHandlers = make([]func(GroupHistoryChangedEvent), 0)
	}
	e.groupHistoryChangedEventHandlers = append(e.groupHistoryChangedEventHandlers, handler)
}
func (e *EventController) triggerGroupHistoryChangedEventHandlers(evtString []byte) {
	var evt GroupHistoryChangedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.groupHistoryChangedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFileCreatedEvents(handler func(FileCreatedEvent)) {
	if e.fileCreatedEventHandlers == nil {
		e.fileCreatedEventHandlers = make([]func(FileCreatedEvent), 0)
	}
	e.fileCreatedEventHandlers = append(e.fileCreatedEventHandlers, handler)
}
func (e *EventController) triggerFileCreatedEventHandlers(evtString []byte) {
	var evt FileCreatedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.fileCreatedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFileSharedEvents(handler func(FileSharedEvent)) {
	if e.fileSharedEventHandlers == nil {
		e.fileSharedEventHandlers = make([]func(FileSharedEvent), 0)
	}
	e.fileSharedEventHandlers = append(e.fileSharedEventHandlers, handler)
}
func (e *EventController) triggerFileSharedEventHandlers(evtString []byte) {
	var evt FileSharedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.fileSharedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFileUnsharedEvents(handler func(FileUnsharedEvent)) {
	if e.fileUnsharedEventHandlers == nil {
		e.fileUnsharedEventHandlers = make([]func(FileUnsharedEvent), 0)
	}
	e.fileUnsharedEventHandlers = append(e.fileUnsharedEventHandlers, handler)
}
func (e *EventController) triggerFileUnsharedEventHandlers(evtString []byte) {
	var evt FileUnsharedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.fileUnsharedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFilePublicEvents(handler func(FilePublicEvent)) {
	if e.filePublicEventHandlers == nil {
		e.filePublicEventHandlers = make([]func(FilePublicEvent), 0)
	}
	e.filePublicEventHandlers = append(e.filePublicEventHandlers, handler)
}
func (e *EventController) triggerFilePublicEventHandlers(evtString []byte) {
	var evt FilePublicEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.filePublicEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFilePrivateEvents(handler func(FilePrivateEvent)) {
	if e.filePrivateEventHandlers == nil {
		e.filePrivateEventHandlers = make([]func(FilePrivateEvent), 0)
	}
	e.filePrivateEventHandlers = append(e.filePrivateEventHandlers, handler)
}
func (e *EventController) triggerFilePrivateEventHandlers(evtString []byte) {
	var evt FilePrivateEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.filePrivateEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFileChangeEvents(handler func(FileChangeEvent)) {
	if e.fileChangeEventHandlers == nil {
		e.fileChangeEventHandlers = make([]func(FileChangeEvent), 0)
	}
	e.fileChangeEventHandlers = append(e.fileChangeEventHandlers, handler)
}
func (e *EventController) triggerFileChangeEventHandlers(evtString []byte) {
	var evt FileChangeEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.fileChangeEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFileDeletedEvents(handler func(FileDeletedEvent)) {
	if e.fileDeletedEventHandlers == nil {
		e.fileDeletedEventHandlers = make([]func(FileDeletedEvent), 0)
	}
	e.fileDeletedEventHandlers = append(e.fileDeletedEventHandlers, handler)
}
func (e *EventController) triggerFileDeletedEventHandlers(evtString []byte) {
	var evt FileDeletedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.fileDeletedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFileCommentAddedEvents(handler func(FileCommentAddedEvent)) {
	if e.fileCommentAddedEventHandlers == nil {
		e.fileCommentAddedEventHandlers = make([]func(FileCommentAddedEvent), 0)
	}
	e.fileCommentAddedEventHandlers = append(e.fileCommentAddedEventHandlers, handler)
}
func (e *EventController) triggerFileCommentAddedEventHandlers(evtString []byte) {
	var evt FileCommentAddedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.fileCommentAddedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFileCommentEditedEvents(handler func(FileCommentEditedEvent)) {
	if e.fileCommentEditedEventHandlers == nil {
		e.fileCommentEditedEventHandlers = make([]func(FileCommentEditedEvent), 0)
	}
	e.fileCommentEditedEventHandlers = append(e.fileCommentEditedEventHandlers, handler)
}
func (e *EventController) triggerFileCommentEditedEventHandlers(evtString []byte) {
	var evt FileCommentEditedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.fileCommentEditedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnFileCommentDeletedEvents(handler func(FileCommentDeletedEvent)) {
	if e.fileCommentDeletedEventHandlers == nil {
		e.fileCommentDeletedEventHandlers = make([]func(FileCommentDeletedEvent), 0)
	}
	e.fileCommentDeletedEventHandlers = append(e.fileCommentDeletedEventHandlers, handler)
}
func (e *EventController) triggerFileCommentDeletedEventHandlers(evtString []byte) {
	var evt FileCommentDeletedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.fileCommentDeletedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnPresenceChangeEvents(handler func(PresenceChangeEvent)) {
	if e.presenceChangeEventHandlers == nil {
		e.presenceChangeEventHandlers = make([]func(PresenceChangeEvent), 0)
	}
	e.presenceChangeEventHandlers = append(e.presenceChangeEventHandlers, handler)
}
func (e *EventController) triggerPresenceChangeEventHandlers(evtString []byte) {
	var evt PresenceChangeEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.presenceChangeEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnManualPresenceChangeEvents(handler func(ManualPresenceChangeEvent)) {
	if e.manualPresenceChangeEventHandlers == nil {
		e.manualPresenceChangeEventHandlers = make([]func(ManualPresenceChangeEvent), 0)
	}
	e.manualPresenceChangeEventHandlers = append(e.manualPresenceChangeEventHandlers, handler)
}
func (e *EventController) triggerManualPresenceChangeEventHandlers(evtString []byte) {
	var evt ManualPresenceChangeEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.manualPresenceChangeEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnPrefChangeEvents(handler func(PrefChangeEvent)) {
	if e.prefChangeEventHandlers == nil {
		e.prefChangeEventHandlers = make([]func(PrefChangeEvent), 0)
	}
	e.prefChangeEventHandlers = append(e.prefChangeEventHandlers, handler)
}
func (e *EventController) triggerPrefChangeEventHandlers(evtString []byte) {
	var evt PrefChangeEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.prefChangeEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnUserChangeEvents(handler func(UserChangeEvent)) {
	if e.userChangeEventHandlers == nil {
		e.userChangeEventHandlers = make([]func(UserChangeEvent), 0)
	}
	e.userChangeEventHandlers = append(e.userChangeEventHandlers, handler)
}
func (e *EventController) triggerUserChangeEventHandlers(evtString []byte) {
	var evt UserChangeEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.userChangeEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnTeamJoinEvents(handler func(TeamJoinEvent)) {
	if e.teamJoinEventHandlers == nil {
		e.teamJoinEventHandlers = make([]func(TeamJoinEvent), 0)
	}
	e.teamJoinEventHandlers = append(e.teamJoinEventHandlers, handler)
}
func (e *EventController) triggerTeamJoinEventHandlers(evtString []byte) {
	var evt TeamJoinEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.teamJoinEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnStarAddedEvents(handler func(StarAddedEvent)) {
	if e.starAddedEventHandlers == nil {
		e.starAddedEventHandlers = make([]func(StarAddedEvent), 0)
	}
	e.starAddedEventHandlers = append(e.starAddedEventHandlers, handler)
}
func (e *EventController) triggerStarAddedEventHandlers(evtString []byte) {
	var evt StarAddedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.starAddedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnStarRemovedEventHandlers(handler func(StarRemovedEvent)) {
	if e.starRemovedEventHandlers == nil {
		e.starRemovedEventHandlers = make([]func(StarRemovedEvent), 0)
	}
	e.starRemovedEventHandlers = append(e.starRemovedEventHandlers, handler)
}
func (e *EventController) triggerStarRemovedEventHandlers(evtString []byte) {
	var evt StarRemovedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.starRemovedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnEmojiChangedEvents(handler func(EmojiChangedEvent)) {
	if e.emojiChangedEventHandlers == nil {
		e.emojiChangedEventHandlers = make([]func(EmojiChangedEvent), 0)
	}
	e.emojiChangedEventHandlers = append(e.emojiChangedEventHandlers, handler)
}
func (e *EventController) triggerEmojiChangedEventHandlers(evtString []byte) {
	var evt EmojiChangedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.emojiChangedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnCommandChangedEvents(handler func(CommandChangedEvent)) {
	if e.commandChangedEventHandlers == nil {
		e.commandChangedEventHandlers = make([]func(CommandChangedEvent), 0)
	}
	e.commandChangedEventHandlers = append(e.commandChangedEventHandlers, handler)
}
func (e *EventController) triggerCommandChangedEventHandlers(evtString []byte) {
	var evt CommandChangedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.commandChangedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnTeamPrefChangeEvents(handler func(TeamPrefChangeEvent)) {
	if e.teamPrefChangeEventHandlers == nil {
		e.teamPrefChangeEventHandlers = make([]func(TeamPrefChangeEvent), 0)
	}
	e.teamPrefChangeEventHandlers = append(e.teamPrefChangeEventHandlers, handler)
}
func (e *EventController) triggerTeamPrefChangeEventHandlers(evtString []byte) {
	var evt TeamPrefChangeEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.teamPrefChangeEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnTeamRenameEvents(handler func(TeamRenameEvent)) {
	if e.teamRenameEventHandlers == nil {
		e.teamRenameEventHandlers = make([]func(TeamRenameEvent), 0)
	}
	e.teamRenameEventHandlers = append(e.teamRenameEventHandlers, handler)
}
func (e *EventController) triggerTeamRenameEventHandlers(evtString []byte) {
	var evt TeamRenameEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.teamRenameEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnTeamDomainChangeEvents(handler func(TeamDomainChangeEvent)) {
	if e.teamDomainChangeEventHandlers == nil {
		e.teamDomainChangeEventHandlers = make([]func(TeamDomainChangeEvent), 0)
	}
	e.teamDomainChangeEventHandlers = append(e.teamDomainChangeEventHandlers, handler)
}
func (e *EventController) triggerTeamDomainChangeEventHandlers(evtString []byte) {
	var evt TeamDomainChangeEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.teamDomainChangeEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnEmailDomainChangeEvents(handler func(EmailDomainChangeEvent)) {
	if e.emailDomainChangeEventHandlers == nil {
		e.emailDomainChangeEventHandlers = make([]func(EmailDomainChangeEvent), 0)
	}
	e.emailDomainChangeEventHandlers = append(e.emailDomainChangeEventHandlers, handler)
}
func (e *EventController) triggerEmailDomainChangeEventHandlers(evtString []byte) {
	var evt EmailDomainChangeEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.emailDomainChangeEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnBotAddedEvents(handler func(BotAddedEvent)) {
	if e.botAddedEventHandlers == nil {
		e.botAddedEventHandlers = make([]func(BotAddedEvent), 0)
	}
	e.botAddedEventHandlers = append(e.botAddedEventHandlers, handler)
}
func (e *EventController) triggerBotAddedEventHandlers(evtString []byte) {
	var evt BotAddedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.botAddedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnBotChangedEvents(handler func(BotChangedEvent)) {
	if e.botChangedEventHandlers == nil {
		e.botChangedEventHandlers = make([]func(BotChangedEvent), 0)
	}
	e.botChangedEventHandlers = append(e.botChangedEventHandlers, handler)
}
func (e *EventController) triggerBotChangedEventHandlers(evtString []byte) {
	var evt BotChangedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.botChangedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnAccountsChangedEvents(handler func(AccountsChangedEvent)) {
	if e.accountsChangedEventHandlers == nil {
		e.accountsChangedEventHandlers = make([]func(AccountsChangedEvent), 0)
	}
	e.accountsChangedEventHandlers = append(e.accountsChangedEventHandlers, handler)
}
func (e *EventController) triggerAccountsChangedEventHandlers(evtString []byte) {
	var evt AccountsChangedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.accountsChangedEventHandlers {
		handler(evt)
	}
}
func (e *EventController) OnTeamMigrationStartedEvents(handler func(TeamMigrationStartedEvent)) {
	if e.teamMigrationStartedEventHandlers == nil {
		e.teamMigrationStartedEventHandlers = make([]func(TeamMigrationStartedEvent), 0)
	}
	e.teamMigrationStartedEventHandlers = append(e.teamMigrationStartedEventHandlers, handler)
}
func (e *EventController) triggerTeamMigrationStartedEventHandlers(evtString []byte) {
	var evt TeamMigrationStartedEvent
	err := json.Unmarshal(evtString, &evt)
	if err != nil {
		return
	}
	for _, handler := range e.teamMigrationStartedEventHandlers {
		handler(evt)
	}
}
