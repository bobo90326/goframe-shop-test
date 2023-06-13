package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" method:"post" desc:"添加角色" tags:"role"`
	Name   string `json:"name" v:"required#名称必填"`
	Desc   string `json:"desc" v:"角色描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id"`
}

type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" method:"post" desc:"修改角色" tags:"role"`
	Id     uint   `json:"id"  v:"required#id必填" desc:"id"`
	Name   string `json:"name" v:"required#名称必填" desc:"角色名称"`
	Desc   string `json:"desc" desc:"角色描述"`
}

type RoleUpdateRes struct {
	Id uint `json:"id"`
}

type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete" tags:"角色" summary:"删除角色接口"`
	Id     uint `json:"id"  v:"required#id必填" desc:"id"`
}

type RoleDeleteRes struct {
	Id uint `json:"id"`
}

type RoleGetListCommonReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"管理员" summary:"角色列表接口"`
	CommonPaginationReq
}
type RoleGetListCommonRes struct {
	//g.Meta `mime:"text/html" type:"string" example:"<html/>"`
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
