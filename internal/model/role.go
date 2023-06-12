package model

type RoleCreateUpdateBase struct {
	Name string
	Desc string
}

// RoleCreateInput 创建内容
type RoleCreateInput struct {
	RoleCreateUpdateBase
}

// RoleCreateOutput 创建内容返回结果
type RoleCreateOutput struct {
	RoleId int `json:"role_id"`
}
