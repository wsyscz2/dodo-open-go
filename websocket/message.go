package websocket

import (
	"github.com/dodo-open/dodo-open-go/model"
	jsoniter "github.com/json-iterator/go"
)

// TypeCode data type code
type TypeCode int

const (
	BizType       TypeCode = 0 // 业务数据
	HeartbeatType TypeCode = 1 // 心跳
)

// WSEventMessage 事件主体内容
type WSEventMessage struct {
	Type TypeCode            `json:"type"`           // 数据类型
	Data jsoniter.RawMessage `json:"data,omitempty"` // 数据内容
}

// EventData 事件业务数据
type EventData struct {
	EventId   string              `json:"eventId"`   // 事件ID
	EventType EventType           `json:"eventType"` // 事件类型
	EventBody jsoniter.RawMessage `json:"eventBody"` // 事件内容
	Timestamp uint64              `json:"timestamp"` // 发送时间戳
}

// WSBeatData 心跳类型数据
type WSBeatData struct {
	Type TypeCode `json:"type"`
}

// EventBody 事件内容
type EventBody interface {
	EventType() EventType
}

type (
	// PersonalModel 个人信息
	PersonalModel struct {
		NickName  string `json:"nickName"`  // 个人昵称
		AvatarUrl string `json:"avatarUrl"` // 个人头像
		Sex       int    `json:"sex"`       // 个人性别，-1：保密，0：女，1：男
	}

	// MemberModel 成员信息
	MemberModel struct {
		NickName string `json:"nickName"` // 在群昵称
		JoinTime string `json:"joinTime"` // 加群时间
	}

	// ReferenceModel 回复信息
	ReferenceModel struct {
		MessageId    string `json:"messageId"`    // 被回复消息ID
		DodoSourceId string `json:"dodoSourceId"` // 被回复者DoDoID
		NickName     string `json:"nickName"`     // 被回复者在群昵称
	}
)

type (
	// PersonalMessageEventBody 个人消息事件
	PersonalMessageEventBody struct {
		IslandSourceId string              `json:"islandSourceId"` // 来源群ID
		DodoSourceId   string              `json:"dodoSourceId"`   // 来源DoDoID
		Personal       *PersonalModel      `json:"personal"`       // 个人信息
		MessageId      string              `json:"messageId"`      // 消息ID
		MessageType    model.MessageType   `json:"messageType"`    // 消息类型，1：文本消息，2：图片消息，3：视频消息
		MessageBody    jsoniter.RawMessage `json:"messageBody"`    // 消息内容（model.IMessageBody）
	}

	// ChannelMessageEventBody 频道消息事件
	ChannelMessageEventBody struct {
		IslandSourceId string              `json:"islandSourceId"` // 来源群ID
		ChannelId      string              `json:"channelId"`      // 来源频道号
		DodoSourceId   string              `json:"dodoSourceId"`   // 来源DoDoID
		Personal       *PersonalModel      `json:"personal"`       // 个人信息
		Member         *MemberModel        `json:"member"`         // 成员信息
		Reference      *ReferenceModel     `json:"reference"`      // 回复信息
		MessageId      string              `json:"messageId"`      // 消息ID
		MessageType    model.MessageType   `json:"messageType"`    // 消息类型，1：文本消息，2：图片消息，3：视频消息，5：文件消息
		MessageBody    jsoniter.RawMessage `json:"messageBody"`    // 消息内容（model.IMessageBody）
	}

	// MessageReactionEventBody 消息反应事件
	MessageReactionEventBody struct {
		IslandSourceId string                `json:"islandSourceId"` // 来源群ID
		ChannelId      string                `json:"channelId"`      // 来源频道号
		DodoSourceId   string                `json:"dodoSourceId"`   // 来源DoDoID
		MessageId      string                `json:"messageId"`      // 消息ID
		Personal       *PersonalModel        `json:"personal"`       // 个人信息
		Member         *MemberModel          `json:"member"`         // 成员信息
		ReactionTarget *model.ReactionTarget `json:"reactionTarget"` // 反应对象
		ReactionEmoji  *model.ReactionEmoji  `json:"reactionEmoji"`  // 反应表情
		ReactionType   int                   `json:"reactionType"`   // 反应类型，0：删除，1：新增
	}

	// CardMessageButtonClickEventBody 卡片消息按钮事件
	CardMessageButtonClickEventBody struct {
		IslandSourceId   string         `json:"islandSourceId"`   // 来源群ID
		ChannelId        string         `json:"channelId"`        // 来源频道ID
		DodoSourceId     string         `json:"dodoSourceId"`     // 来源DoDoID
		MessageId        string         `json:"messageId"`        // 来源消息ID，频道私信触发时，返回频道私信消息ID
		Personal         *PersonalModel `json:"personal"`         // 个人信息
		Member           *MemberModel   `json:"member"`           // 成员信息
		InteractCustomId string         `json:"interactCustomId"` // 交互自定义ID
		Value            string         `json:"value"`            // Value
	}

	// CardMessageFormSubmitEventBody 卡片消息表单回传事件
	CardMessageFormSubmitEventBody struct {
		IslandSourceId   string           `json:"islandSourceId"`   // 来源群ID
		ChannelId        string           `json:"channelId"`        // 来源频道ID
		DodoSourceId     string           `json:"dodoSourceId"`     // 来源DoDoID
		MessageId        string           `json:"messageId"`        // 来源消息ID，频道私信触发时，返回频道私信消息ID
		Personal         *PersonalModel   `json:"personal"`         // 个人信息
		Member           *MemberModel     `json:"member"`           // 成员信息
		InteractCustomId string           `json:"interactCustomId"` // 交互自定义ID
		FormData         []*FormDataModel `json:"formData"`         // 表单数据
	}
	FormDataModel struct {
		Key   string `json:"key"`   // 选项自定义ID
		Value string `json:"value"` // Value
	}

	// CardMessageListSubmitEventBody 卡片消息列表回传事件
	CardMessageListSubmitEventBody struct {
		IslandSourceId   string           `json:"islandSourceId"`   // 来源群ID
		ChannelId        string           `json:"channelId"`        // 来源频道ID
		DodoSourceId     string           `json:"dodoSourceId"`     // 来源DoDoID
		MessageId        string           `json:"messageId"`        // 来源消息ID，频道私信触发时，返回频道私信消息ID
		Personal         *PersonalModel   `json:"personal"`         // 个人信息
		Member           *MemberModel     `json:"member"`           // 成员信息
		InteractCustomId string           `json:"interactCustomId"` // 交互自定义ID
		ListData         []*ListDataModel `json:"listData"`         // 列表数据
	}
	ListDataModel struct {
		Name string `json:"name"` // 选项名
	}

	// MemberJoinEventBody 成员加入事件
	MemberJoinEventBody struct {
		IslandSourceId string `json:"islandSourceId"` // 来源群ID
		DodoSourceId   string `json:"dodoSourceId"`   // 来源DoDoID
		ModifyTime     string `json:"modifyTime"`     // 变动时间
	}

	// MemberLeaveEventBody 成员退出事件
	MemberLeaveEventBody struct {
		IslandSourceId      string         `json:"islandSourceId"`      // 来源群ID
		DodoSourceId        string         `json:"dodoSourceId"`        // 来源DoDoID
		Personal            *PersonalModel `json:"personal"`            // 个人信息
		LeaveType           int            `json:"leaveType"`           // 退出类型，1：主动，2：被踢
		OperateDoDoSourceId string         `json:"operateDoDoSourceId"` // 操作者DoDoID（执行踢出操作的人）
		ModifyTime          string         `json:"modifyTime"`          // 变动时间
	}

	// PersonalMessageEventBody 个人消息事件
	ChannelArticleEventBody struct {
		IslandSourceId string         `json:"islandSourceId"` // 来源群ID
		ChannelId      string         `json:"channelId"`      // 来源帖子频道ID
		DodoSourceId   string         `json:"dodoSourceId"`   // 来源DoDoID
		Personal       *PersonalModel `json:"personal"`       // 个人信息
		Member         *MemberModel   `json:"membe"`          // 成员信息
		ArticleId      string         `json:"articleId"`      // 帖子id
		Title          string         `json:"title"`          // 标题
		ImageList      []string       `json:"imageList"`      // 图片列表
		Content        string         `json:"content"`        // 文本内容
	}

	ChannelArticleCommentEventBody struct {
		IslandSourceId string         `json:"islandSourceId"` // 来源群ID
		ChannelId      string         `json:"channelId"`      // 来源帖子频道ID
		DodoSourceId   string         `json:"dodoSourceId"`   // 来源DoDoID
		Personal       *PersonalModel `json:"personal"`       // 个人信息
		Member         *MemberModel   `json:"membe"`          // 成员信息
		ArticleId      string         `json:"articleId"`      // 帖子ID
		CommentId      string         `json:"commentId"`      // 帖子评论ID
		ReplyId        string         `json:"replyId"`        // 帖子评论回复ID，为空时：为评论事件，不为空时：为评论回复事件
		ImageList      []string       `json:"imageList"`      // 图片列表
		Content        string         `json:"content"`        // 文本内容
	}
)

func (e *PersonalMessageEventBody) EventType() EventType {
	return PersonalMessageEvent
}

func (e *ChannelMessageEventBody) EventType() EventType {
	return ChannelMessageEvent
}

func (e *MessageReactionEventBody) EventType() EventType {
	return MessageReactionEvent
}

func (e *MemberJoinEventBody) EventType() EventType {
	return MemberJoinEvent
}

func (e *MemberLeaveEventBody) EventType() EventType {
	return MemberLeaveEvent
}
