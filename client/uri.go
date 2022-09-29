package client

import "fmt"

type uri string

const (
	// 机器人 API
	getBotInfoUri        uri = "/api/v1/bot/info"         // 获取机器人信息
	setBotIslandLeaveUri uri = "/api/v1/bot/island/leave" // 机器人退群

	// 群 API
	getIslandListUri          uri = "/api/v1/island/list"            // 获取群列表
	getIslandInfoUri          uri = "/api/v1/island/info"            // 获取群信息
	getIslandLevelRankListUri uri = "/api/v1/island/level/rank/list" // 获取群等级排行榜
	getIslandMuteListUri      uri = "/api/v1/island/mute/list"       // 获取群禁言名单
	getIslandBanListUri       uri = "/api/v1/island/ban/list"        // 获取群封禁名单

	// 频道 API
	getChannelListUri uri = "/api/v1/channel/list"   // 获取频道列表
	getChannelInfoUri uri = "/api/v1/channel/info"   // 获取频道信息
	createChannelUri  uri = "/api/v1/channel/add"    // 创建频道
	editChannelUri    uri = "/api/v1/channel/edit"   // 编辑频道
	removeChannelUri  uri = "/api/v1/channel/remove" // 删除频道

	// 文字频道 API
	sendChannelMessageUri     uri = "/api/v1/channel/message/send"            // 发送消息
	editChannelMessageUri     uri = "/api/v1/channel/message/edit"            // 编辑消息
	withdrawChannelMessageUri uri = "/api/v1/channel/message/withdraw"        // 撤回消息
	addChannelMessageReaction uri = "/api/v1/channel/message/reaction/add"    // 添加表情反应
	remChannelMessageReaction uri = "/api/v1/channel/message/reaction/remove" // 取消表情反应

	// 帖子频道 API
	setChannelArticleAdd uri = "/api/v1/channel/article/add" // 发布帖子

	// 身份组 API
	getRoleListUri      uri = "/api/v1/role/list"          // 获取身份组列表
	createRoleUri       uri = "/api/v1/role/add"           // 创建身份组
	editRoleUri         uri = "/api/v1/role/edit"          // 编辑身份组
	removeRoleUri       uri = "/api/v1/role/remove"        // 删除身份组
	addRoleMemberUri    uri = "/api/v1/role/member/add"    // 赋予成员身份组
	removeRoleMemberUri uri = "/api/v1/role/member/remove" // 取消成员身份组

	// 成员 API
	getMemberListUri     uri = "/api/v1/member/list"            // 获取成员列表
	getMemberInfoUri     uri = "/api/v1/member/info"            // 获取成员信息
	getMemberRoleListUri uri = "/api/v1/member/role/list"       // 获取成员身份组列表
	getMemberInviteInfo  uri = "/api/v1/member/invitation/info" // 获取成员邀请信息
	setMemberNickUri     uri = "/api/v1/member/nickname/edit"   // 编辑成员群昵称
	muteMemberUri        uri = "/api/v1/member/mute/add"        // 禁言成员
	unmuteMemberUri      uri = "/api/v1/member/mute/remove"     // 取消成员禁言
	banMemberUri         uri = "/api/v1/member/ban/set"         // 永久封禁成员
	unbanMemberUri       uri = "/api/v1/member/ban/remove"      // 取消成员永久封禁

	// 数字藏品 API
	getMemberNFTStatusUri    uri = "/api/v1/member/nft/status"       // 获取成员数字藏品判断
	getMemberUPowerchainInfo uri = "/api/v1/member/upowerchain/info" // 取成员高能链数字藏品信息

	// 私信 API
	sendDirectMessageUri uri = "/api/v1/personal/message/send" // 发送私信

	// 资源 API
	uploadImageUri uri = "/api/v1/resource/picture/upload" // 上传图片资源

	// 事件 API
	getWebsocketConnectionUri uri = "/api/v1/websocket/connection" // 获取 Websocket 连接
)

// getApi build the full api url
func (c *client) getApi(u uri) string {
	return fmt.Sprintf("%s%s", c.conf.BaseApi, u)
}
