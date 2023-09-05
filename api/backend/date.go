package backend

import "github.com/gogf/gf/v2/frame/g"

type DataHeadReq struct {
	g.Meta `path:"/data/head" method:"get" tags:"数据data"  desc:"获取数据头"`
}

type DataHeadRes struct {
	TodayOrderCount int `json:"todayOrderCount" desc:"今日订单数"`
	DAU             int `json:"dau" desc:"日活"`
	ConversionRate  int `json:"conversionRate" desc:"转化率"`
}
