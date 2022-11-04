package model

import "errors"

// GetBotInfoRsp 获取机器人信息 response
type GetBotInfoRsp struct {
	ClientId     string `json:"clientId"`     // 机器人唯一标识
	DodoSourceId string `json:"dodoSourceId"` // 机器人DoDoID
	NickName     string `json:"nickName"`     // 机器人昵称
	AvatarUrl    string `json:"avatarUrl"`    // 机器人图标
}

// SetBotLeaveIslandReq 机器人退群 request
type SetBotLeaveIslandReq struct {
	IslandSourceId string `json:"islandSourceId" binding:"required"` // 群ID
}

func (p *SetBotLeaveIslandReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter IslandId (empty detected)")
	}
	return nil
}

// GetBotInviteListReq
type GetBotInviteListReq struct {
	PageSize int   `json:"pageSize"` // 页大小，最大100
	MaxId    int64 `json:"maxId"`    // 上一页最大ID值，为提升分页查询性能，需要传入上一页查询记录中的最大ID值，首页请传0
}

// GetBotInviteListRsp
type GetBotInviteListRsp struct {
	MaxId int64            `json:"maxId"` // 最大ID值
	List  []*BotInviteItem `json:"list"`  // 数据列表
}
type BotInviteItem struct {
	DodoSourceId string `json:"dodoSourceId"` // DoDoID
	NickName     string `json:"nickName"`     // DoDo昵称
	AvatarUrl    string `json:"avatarUrl"`    // 头像
}
