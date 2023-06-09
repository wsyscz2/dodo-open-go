package client

import (
	"context"

	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

// GetIntegralInfo 查询成员积分
func (c *client) GetIntegralInfo(ctx context.Context, req *model.GetIntegralInfoReq) (*model.GetIntegralInfoRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(getIntegralInfoUri))
	if err != nil {
		return nil, err
	}

	result := &model.GetIntegralInfoRsp{}
	if err := tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) SetIntegralEdit(ctx context.Context, req *model.SetIntegralEditReq) (*model.SetIntegralEditRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}
	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setIntegralEditUri))
	if err != nil {
		return nil, err
	}

	result := &model.SetIntegralEditRsp{}
	if err := tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
