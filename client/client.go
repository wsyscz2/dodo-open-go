package client

import (
	"context"
	"errors"
	"time"

	"github.com/dodo-open/dodo-open-go/model"
	"github.com/go-resty/resty/v2"
)

type (
	// Client DoDoBot API Interface
	Client interface {
		Base
		BotAPI
		IslandAPI
		ChannelAPI
		ChannelVoiceAPI
		ChannelArticleAPI
		MessageAPI
		RoleAPI
		MemberAPI
		DigitalAssetAPI
		DirectMessageAPI
		ResourceUploadAPI
		WebsocketAPI
		GiftAPI
		IntegralAPI

		GetMemberDodoIdMapList(ctx context.Context, req *model.GetMemberDodoIdMapListReq) ([]*model.MemberDodoIdMap, error)
	}

	// Base client basic interface
	Base interface {
		GetConfig() *Config
	}

	// BotAPI bot API interface
	BotAPI interface {
		GetBotInfo(ctx context.Context) (*model.GetBotInfoRsp, error)                                             // GetBotInfo 获取机器人信息
		SetBotIslandLeave(ctx context.Context, req *model.SetBotLeaveIslandReq) (bool, error)                     // SetBotIslandLeave 机器人退群
		GetBotInviteList(ctx context.Context, req *model.GetBotInviteListReq) (*model.GetBotInviteListRsp, error) // GetBotInviteList 获取机器人邀请列表
		SetBotInviteAdd(ctx context.Context, req *model.SetBotInviteAddReq) (bool, error)                         // SetBotInviteAdd 添加成员到机器人邀请列表
		SetBotInviteRemove(ctx context.Context, req *model.SetBotInviteRemoveReq) (bool, error)                   // SetBotInviteRemove 移除成员出机器人邀请列表
	}

	// IslandAPI island API interface
	IslandAPI interface {
		GetIslandList(ctx context.Context) ([]*model.IslandElement, error)                                                            // GetIslandList 获取群列表
		GetIslandInfo(ctx context.Context, req *model.GetIslandInfoReq) (*model.GetIslandInfoRsp, error)                              // GetIslandInfo 获取群信息
		GetIslandLevelRankList(ctx context.Context, req *model.GetIslandLevelRankListReq) ([]*model.GetIslandLevelRankElement, error) // GetIslandLevelRankList 获取群等级排行榜
		GetIslandMuteList(ctx context.Context, req *model.GetIslandMuteListReq) (*model.GetIslandMuteListRsp, error)                  // GetIslandMuteList 获取群禁言名单
		GetIslandBanList(ctx context.Context, req *model.GetIslandBanListReq) (*model.GetIslandBanListRsp, error)                     // GetIslandBanList 获取群封禁名单
	}

	// ChannelAPI channel basic API interface
	ChannelAPI interface {
		GetChannelList(ctx context.Context, req *model.GetChannelListReq) ([]*model.ChannelElement, error)  // GetChannelList 获取频道列表
		GetChannelInfo(ctx context.Context, req *model.GetChannelInfoReq) (*model.GetChannelInfoRsp, error) // GetChannelInfo 获取频道信息
		CreateChannel(ctx context.Context, req *model.CreateChannelReq) (*model.CreateChannelRsp, error)    // CreateChannel 创建频道
		EditChannel(ctx context.Context, req *model.EditChannelReq) (bool, error)                           // EditChannel 编辑频道
		RemoveChannel(ctx context.Context, req *model.RemoveChannelReq) (bool, error)                       // RemoveChannel 编辑频道
	}

	// ChannelVoiceAPI channel voice API interface
	ChannelVoiceAPI interface {
		GetChannelVoiceMemberStatus(ctx context.Context, req *model.GetChannelVoiceMemberStatusReq) (*model.GetChannelVoiceMemberStatusRsp, error) // GetChannelVoiceMemberStatus 获取成员语音频道状态
		SetChannelVoiceMemberMove(ctx context.Context, req *model.SetChannelVoiceMemberMoveReq) (bool, error)                                      // SetChannelVoiceMemberMove 移动语音频道成员
		SetChannelVoiceMemberEdit(ctx context.Context, req *model.SetChannelVoiceMemberEditReq) (bool, error)                                      // SetChannelVoiceMemberEdit 管理语音中的成员
	}
	// ChannelArticleAPI  channel article API interface
	ChannelArticleAPI interface {
		SetChannelArticleAdd(ctx context.Context, req *model.SetChannelArticleAddReq) (*model.SetChannelArticleAddRsp, error) // SetChannelArticleAdd 发布帖子
		SetChannelArticleRemove(ctx context.Context, req *model.SetChannelArticleRemoveReq) (bool, error)                     // SetChannelArticleRemove 删除帖子评论回复
	}

	// MessageAPI message API interface
	MessageAPI interface {
		SendChannelMessage(ctx context.Context, req *model.SendChannelMessageReq) (*model.SendChannelMessageRsp, error)                                                    // SendChannelMessage 发送消息
		EditChannelMessage(ctx context.Context, req *model.EditChannelMessageReq) (*model.EditChannelMessageRsp, error)                                                    // EditChannelMessage 编辑消息
		WithdrawChannelMessage(ctx context.Context, req *model.WithdrawChannelMessageReq) (bool, error)                                                                    // WithdrawChannelMessage 撤回消息
		AddChannelMessageReaction(ctx context.Context, req *model.AddChannelMessageReactionReq) (bool, error)                                                              // AddChannelMessageReaction 添加表情反应
		RemChannelMessageReaction(ctx context.Context, req *model.RemChannelMessageReactionReq) (bool, error)                                                              // RemChannelMessageReaction 取消表情反应
		SetChannelMessageTop(ctx context.Context, req *model.SetChannelMessageTopReq) (bool, error)                                                                        // SetChannelMessageTop 置顶消息
		GetChannelMessageReactionList(ctx context.Context, req *model.GetChannelMessageReactionListReq) ([]*model.GetChannelMessageReactionListRsp, error)                 // GetChannelMessageReactionList 获取消息反应列表
		GetChannelMessageReactionMemberList(ctx context.Context, req *model.GetChannelMessageReactionMemberListReq) (*model.GetChannelMessageReactionMemberListRsp, error) // GetChannelMessageReactionMemberList 获取消息反应内成员列表
	}

	// RoleAPI role API interface
	RoleAPI interface {
		GetRoleList(ctx context.Context, req *model.GetRoleListReq) ([]*model.RoleElement, error)                    // GetRoleList 获取身份组列表
		CreateRole(ctx context.Context, req *model.CreateRoleReq) (*model.CreateRoleRsp, error)                      // CreateRole 创建身份组
		EditRole(ctx context.Context, req *model.EditRoleReq) (bool, error)                                          // EditRole 编辑身份组
		RemoveRole(ctx context.Context, req *model.RemoveRoleReq) (bool, error)                                      // RemoveRole 删除身份组
		AddRoleMember(ctx context.Context, req *model.AddRoleMemberReq) (bool, error)                                // AddRoleMember 赋予成员身份组
		RemoveRoleMember(ctx context.Context, req *model.RemoveRoleMemberReq) (bool, error)                          // RemoveRoleMember 取消成员身份组
		GetRoleMemberList(ctx context.Context, req *model.GetRoleMemberListReq) (*model.GetRoleMemberListRsp, error) // GetRoleMemberList 获取身份组成员列表
	}

	// MemberAPI member API interface
	MemberAPI interface {
		GetMemberList(ctx context.Context, req *model.GetMemberListReq) (*model.GetMemberListRsp, error)                   // GetMemberList 获取成员列表
		GetMemberInfo(ctx context.Context, req *model.GetMemberInfoReq) (*model.GetMemberInfoRsp, error)                   // GetMemberInfo 获取成员信息
		GetMemberRoleList(ctx context.Context, req *model.GetMemberRoleListReq) ([]*model.RoleElement, error)              // GetMemberRoleList 获取成员身份组列表
		GetMemberInviteInfo(ctx context.Context, req *model.GetMemberInviteInfoReq) (*model.GetMemberInviteInfoRsp, error) // GetMemberInviteInfo 获取成员邀请信息
		SetMemberNick(ctx context.Context, req *model.SetMemberNickReq) (bool, error)                                      // SetMemberNick 编辑成员群昵称
		MuteMember(ctx context.Context, req *model.MuteMemberReq) (bool, error)                                            // MuteMember 禁言成员
		UnmuteMember(ctx context.Context, req *model.UnmuteMemberReq) (bool, error)                                        // UnmuteMember 取消成员禁言
		BanMember(ctx context.Context, req *model.BanMemberReq) (bool, error)                                              // BanMember 永久封禁成员
		UnbanMember(ctx context.Context, req *model.UnbanMemberReq) (bool, error)                                          // UnbanMember 取消成员永久封禁
	}

	// GiftAPI gift API interface
	GiftAPI interface {
		GetGiftAccount(ctx context.Context, req *model.GetGiftAccountReq) (*model.GetGiftAccountRsp, error)                      // GetGiftAccount 获取群收入
		GetGiftShareRatioInfo(ctx context.Context, req *model.GetGiftShareRatioInfoReq) (*model.GetGiftShareRatioInfoRsp, error) // GetGiftShareRatioInfo 获取成员分成管理
		GetGiftList(ctx context.Context, req *model.GetGiftListReq) ([]*model.GetGiftListRsp, error)                             // GetGiftList 获取内容礼物列表
		GetGiftMemberList(ctx context.Context, req *model.GetGiftMemberListReq) (*model.GetGiftMemberListRsp, error)             // GetGiftMemberList 获取内容礼物内成员列表
		GetGiftGrossValueList(ctx context.Context, req *model.GetGiftGrossValueListReq) (*model.GetGiftGrossValueListRsp, error) // GetGiftGrossValueList 获取内容礼物总值列表
	}

	IntegralAPI interface {
		GetIntegralInfo(ctx context.Context, req *model.GetIntegralInfoReq) (*model.GetIntegralInfoRsp, error) // GetIntegralInfo 查询成员积分
		SetIntegralEdit(ctx context.Context, req *model.SetIntegralEditReq) (*model.SetIntegralEditRsp, error) // SetIntegralEdit 管理成员积分
	}

	// DigitalAssetAPI digital asset API interface
	DigitalAssetAPI interface {
		GetMemberNFTStatus(ctx context.Context, req *model.GetMemberNFTStatusReq) (*model.GetMemberNFTStatusRsp, error)                   // GetMemberNFTStatus 获取成员数字藏品判断
		GetMemberUPowerchainInfo(ctx context.Context, req *model.GetMemberUPowerchainInfoReq) (*model.GetMemberUPowerchainInfoRsp, error) // GetMemberUPowerchainInfo 取成员高能链数字藏品信息
	}

	// DirectMessageAPI direct message (a.k.a. DM) API interface
	DirectMessageAPI interface {
		SendDirectMessage(ctx context.Context, req *model.SendDirectMessageReq) (*model.SendDirectMessageRsp, error) // SendDirectMessage 发送私信
	}

	// ResourceUploadAPI resource upload API interface
	ResourceUploadAPI interface {
		UploadImageByBytes(ctx context.Context, req *model.UploadImageByBytesReq) (*model.UploadImageRsp, error)
		UploadImageByPath(ctx context.Context, req *model.UploadImageByPathReq) (*model.UploadImageRsp, error)
	}

	// WebsocketAPI websocket API interface
	WebsocketAPI interface {
		GetWebsocketConnection(ctx context.Context) (*model.GetWebsocketConnectionRsp, error)
	}
)

type (
	// client DoDoBot Instance
	client struct {
		conf *Config
		r    *resty.Client // resty client
	}

	// Config DoDoBot client configuration
	Config struct {
		BaseApi  string        // DoDo OpenAPI Server Domain
		ClientId string        // DoDoBot ClientID
		Token    string        // DoDoBot Bot token
		IsDebug  bool          // debug mode
		Timeout  time.Duration // resty client request timeout
	}
)

// New create a new DoDoBot instance
func New(clientId, token string, options ...OptionHandler) (Client, error) {
	config := getDefaultConfig()
	config.ClientId = clientId
	config.Token = token

	// handle custom options
	for _, optionHandler := range options {
		if optionHandler == nil {
			return nil, errors.New("invalid OptionHandler (nil detected)")
		}
		if err := optionHandler(config); err != nil {
			return nil, err
		}
	}

	instance := &client{conf: config}
	instance.setupResty()

	return instance, nil
}

// getDefaultConfig Get the default configuration
func getDefaultConfig() *Config {
	return &Config{
		BaseApi: "https://botopen.imdodo.com",
		IsDebug: false,
		Timeout: time.Second * 5,
	}
}

// GetConfig get instance configuration
func (c *client) GetConfig() *Config {
	return c.conf
}
