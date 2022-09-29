package client

import (
	"context"

	"github.com/dodo-open/dodo-open-go/model"
	"github.com/dodo-open/dodo-open-go/tools"
)

func (c *client) SetChannelArticleAdd(ctx context.Context, req *model.SetChannelArticleAddReq) (*model.SetChannelArticleAddRsp, error) {
	if err := req.ValidParams(); err != nil {
		return nil, err
	}

	resp, err := c.request(ctx).SetBody(req).Post(c.getApi(setChannelArticleAdd))
	if err != nil {
		return nil, err
	}

	result := &model.SetChannelArticleAddRsp{}
	if err = tools.JSON.Unmarshal(c.unmarshalResult(resp).Data, &result); err != nil {
		return nil, err
	}
	return result, nil
}
