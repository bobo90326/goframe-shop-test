package data

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-shop-test/internal/dao"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/service"
	"goframe-shop-test/utility"
)

type sData struct {
}

func init() {
	service.RegisterData(New())
}

func New() *sData {
	return &sData{}
}

func (s *sData) DataHead(ctx context.Context) (out *model.DataHeadOutput, err error) {

	return &model.DataHeadOutput{
		TodayOrderCount: s.todayOrderCount(ctx),
		DAU:             utility.RandInt(1000),
		ConversionRate:  utility.RandInt(50),
	}, nil
}

//查询今日订单数
func (s *sData) todayOrderCount(ctx context.Context) (count int) {
	count, err := dao.OrderInfo.Ctx(ctx).
		WhereBetween(dao.OrderInfo.Columns().CreatedAt, gtime.Now().StartOfDay(), gtime.Now().EndOfDay()).
		Count(dao.OrderInfo.Columns().Id)
	if err != nil {
		return -1
	}
	return count
}
