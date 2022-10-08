package websocket

// MessageHandlers instance message handlers
type MessageHandlers struct {
	PersonalMessage        PersonalMessageEventHandler
	ChannelMessage         ChannelMessageEventHandler
	MessageReaction        MessageReactionEventHandler
	CardMessageButtonClick CardMessageButtonClickHandler
	CardMessageFormSubmit  CardMessageFormSubmitHandler
	CardMessageListSubmit  CardMessageListSubmitHandler
	MemberJoin             MemberJoinEventHandler
	MemberLeave            MemberLeaveEventHandler
	ChannelArticle         ChannelArticleEventHandler
	ChannelArticleComment  ChannelArticleCommentEventHandler

	PlainTextHandler PlainTextHandler
	ErrorHandler     ErrorHandler
}

// DefaultHandlers default handlers to manage all supported message
var DefaultHandlers = &MessageHandlers{
	PersonalMessage:        func(event *WSEventMessage, data *PersonalMessageEventBody) error { return nil },
	ChannelMessage:         func(event *WSEventMessage, data *ChannelMessageEventBody) error { return nil },
	MessageReaction:        func(event *WSEventMessage, data *MessageReactionEventBody) error { return nil },
	CardMessageButtonClick: func(event *WSEventMessage, data *CardMessageButtonClickEventBody) error { return nil },
	CardMessageFormSubmit:  func(event *WSEventMessage, data *CardMessageFormSubmitEventBody) error { return nil },
	CardMessageListSubmit:  func(event *WSEventMessage, data *CardMessageListSubmitEventBody) error { return nil },
	MemberJoin:             func(event *WSEventMessage, data *MemberJoinEventBody) error { return nil },
	MemberLeave:            func(event *WSEventMessage, data *MemberLeaveEventBody) error { return nil },
	ChannelArticle:         func(event *WSEventMessage, data *ChannelArticleEventBody) error { return nil },
	ChannelArticleComment:  func(event *WSEventMessage, data *ChannelArticleCommentEventBody) error { return nil },
	PlainTextHandler:       func(event *WSEventMessage, message []byte) error { return nil },
	ErrorHandler:           func(err error) {},
}

func fillHandler(handlers *MessageHandlers) *MessageHandlers {
	if handlers.PersonalMessage == nil {
		handlers.PersonalMessage = DefaultHandlers.PersonalMessage
	}
	if handlers.ChannelMessage == nil {
		handlers.ChannelMessage = DefaultHandlers.ChannelMessage
	}
	if handlers.MessageReaction == nil {
		handlers.MessageReaction = DefaultHandlers.MessageReaction
	}
	if handlers.MemberJoin == nil {
		handlers.MemberJoin = DefaultHandlers.MemberJoin
	}
	if handlers.MemberLeave == nil {
		handlers.MemberLeave = DefaultHandlers.MemberLeave
	}
	if handlers.PlainTextHandler == nil {
		handlers.PlainTextHandler = DefaultHandlers.PlainTextHandler
	}
	if handlers.ErrorHandler == nil {
		handlers.ErrorHandler = DefaultHandlers.ErrorHandler
	}
	return handlers
}

// PersonalMessageEventHandler 个人消息事件 handler
type PersonalMessageEventHandler func(event *WSEventMessage, data *PersonalMessageEventBody) error

// ChannelMessageEventHandler 频道消息事件 handler
type ChannelMessageEventHandler func(event *WSEventMessage, data *ChannelMessageEventBody) error

// MessageReactionEventHandler 消息反应事件 handler
type MessageReactionEventHandler func(event *WSEventMessage, data *MessageReactionEventBody) error

// CardMessageButtonClickHandler 卡片消息按钮事件 handler
type CardMessageButtonClickHandler func(event *WSEventMessage, data *CardMessageButtonClickEventBody) error

// CardMessageFormSubmitHandler 卡片消息表单回传事件 handler
type CardMessageFormSubmitHandler func(event *WSEventMessage, data *CardMessageFormSubmitEventBody) error

// CardMessageListSubmitHandler 卡片消息列表回传事件 handler
type CardMessageListSubmitHandler func(event *WSEventMessage, data *CardMessageListSubmitEventBody) error

// MemberJoinEventHandler 成员加入事件 handler
type MemberJoinEventHandler func(event *WSEventMessage, data *MemberJoinEventBody) error

// MemberLeaveEventHandler 成员退出事件 handler
type MemberLeaveEventHandler func(event *WSEventMessage, data *MemberLeaveEventBody) error

// ChannelArticleEventHandler 发帖事件 handler
type ChannelArticleEventHandler func(event *WSEventMessage, data *ChannelArticleEventBody) error

// ChannelArticleEventHandler 帖子评论回复事件 handler
type ChannelArticleCommentEventHandler func(event *WSEventMessage, data *ChannelArticleCommentEventBody) error

// PlainTextHandler plain text message handler
type PlainTextHandler func(event *WSEventMessage, message []byte) error

// ErrorHandler error handler
type ErrorHandler func(err error)

// RegisterHandlers Register global level event message handlers
func RegisterHandlers(handlers ...interface{}) {
	for _, h := range handlers {
		switch handle := h.(type) {
		case PersonalMessageEventHandler:
			DefaultHandlers.PersonalMessage = handle
		case ChannelMessageEventHandler:
			DefaultHandlers.ChannelMessage = handle
		case MessageReactionEventHandler:
			DefaultHandlers.MessageReaction = handle
		case MemberJoinEventHandler:
			DefaultHandlers.MemberJoin = handle
		case MemberLeaveEventHandler:
			DefaultHandlers.MemberLeave = handle
		default:
			// other handlers will be registered in the following functions
			// non-business handler will be registered here
			registerOtherHandlers(handlers)
		}
	}
}

// registerOtherHandlers Register global level non-business handlers
func registerOtherHandlers(handlers ...interface{}) {
	for _, h := range handlers {
		switch handle := h.(type) {
		case PlainTextHandler:
			DefaultHandlers.PlainTextHandler = handle
		case ErrorHandler:
			DefaultHandlers.ErrorHandler = handle
		default:
		}
	}
}
