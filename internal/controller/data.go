package controller

import (
	"context"
	"goframe-shop-test/api/backend"
	"goframe-shop-test/internal/service"
)

var Data = cData{}

type cData struct {
}

func (c *cData) DataHead(ctx context.Context, req backend.DataHeadReq) (res *backend.DataHeadRes, err error) {
	dataHead, err := service.Data().DataHead(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.DataHeadRes{
		TodayOrderCount: dataHead.TodayOrderCount,
		DAU:             dataHead.DAU,
		ConversionRate:  dataHead.ConversionRate,
	}, nil
}
