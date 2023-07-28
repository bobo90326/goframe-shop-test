package model

type DataHeadOutput struct {
	TodayOrderCount int `json:"todayOrderCount" desc:"今日订单数"`
	DAU             int `json:"dau" desc:"日活"`
	ConversionRate  int `json:"conversionRate" desc:"转化率"`
}
