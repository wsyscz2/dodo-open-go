package model

import "errors"

// GetIntegralInfoReq 查询成员积分
// [https://open.imdodo.com/dev/api/integral.html#%E6%9F%A5%E8%AF%A2%E6%88%90%E5%91%98%E7%A7%AF%E5%88%86]
type GetIntegralInfoReq struct {
	IslandSourceId string `json:"islandSourceId"` // 群ID
	DodoSourceId   string `json:"dodoSourceId"`   // DoDoID
}

func (p *GetIntegralInfoReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter islandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter dodoSourceId (empty detected)")
	}
	return nil
}

// GetIntegralInfoRsp 查询成员积分 响应
// [https://open.imdodo.com/dev/api/integral.html#%E6%9F%A5%E8%AF%A2%E6%88%90%E5%91%98%E7%A7%AF%E5%88%86]
type GetIntegralInfoRsp struct {
	IntegralBalance int64 `json:"integralBalance"` // 积分余额
}

// SetIntegralEditReq 管理成员积分 请求
// [https://open.imdodo.com/dev/api/integral.html#%E7%AE%A1%E7%90%86%E6%88%90%E5%91%98%E7%A7%AF%E5%88%86]
type SetIntegralEditReq struct {
	IslandSourceId string `json:"islandSourceId"` // 群ID
	DodoSourceId   string `json:"dodoSourceId"`   // DoDoID
	OperateType    int    `json:"operateType"`    // 管理类型，1：发放积分，2：扣除积分
	Integral       int64  `json:"integral"`       // 积分，必须是正整数且小于1千亿
}

func (p *SetIntegralEditReq) ValidParams() error {
	if p.IslandSourceId == "" {
		return errors.New("invalid parameter islandSourceId (empty detected)")
	}
	if p.DodoSourceId == "" {
		return errors.New("invalid parameter dodoSourceId (empty detected)")
	}

	if p.Integral <= 0 || p.Integral >= 100000000000 {
		return errors.New("invalid parameter integral (must be positive and less than 100000000000)")
	}

	return nil
}

// SetIntegralEditRsp 管理成员积分 响应
// [https://open.imdodo.com/dev/api/integral.html#%E7%AE%A1%E7%90%86%E6%88%90%E5%91%98%E7%A7%AF%E5%88%86]
type SetIntegralEditRsp struct {
	IntegralBalance int64 `json:"integralBalance"` // 积分余额
}
