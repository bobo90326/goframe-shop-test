package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleReq struct {
	g.Meta `path:"/backend/role/add" method:"post"  desc:"添加角色" tags:"role"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc" v:"required#角色描述不能为空" dc:"角色描述"`
}

type RoleRes struct {
	RoleId int `json:"role_id"`
}

//修改角色
type RoleUpdateReq struct {
	g.Meta `path:"/backend/role/update" method:"post"  desc:"修改角色" tags:"role"`
	Id     uint   `json:"id" v:"required#角色id不能为空" dc:"角色id"`
	Name   string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Desc   string `json:"desc"  dc:"角色描述"`
}

type RoleUpdateRes struct {
	RoleId uint `json:"role_id"`
}

//删除角色
type RoleDeleteReq struct {
	g.Meta `path:"/backend/role/delete" method:"delete"  desc:"删除角色" tags:"role" summary:"删除角色"`
	Id     uint `json:"id" v:"required#角色id不能为空" dc:"角色id"`
}

type RoleDeleteRes struct {
}

//获取角色列表
type RoleGetListCommonReq struct {
	g.Meta `path:"/backend/role/list" method:"get" tags:"role" summary:"获取角色列表"`
	CommonPaginationReq
}

type RoleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
