package backend

import "github.com/gogf/gf/v2/frame/g"

type PermissionCreateBase struct {
	Name string `json:"name" v:"required#权限名称不能为空" dc:"权限名称"`
	Path string `json:"desc"  dc:"权限描述"`
}

type PermissionReq struct {
	g.Meta `path:"/backend/permission/add" method:"post"  desc:"添加权限" tags:"permission"`
	PermissionCreateBase
}

type PermissionRes struct {
	PermissionId int `json:"permission_id"`
}

//修改权限
type PermissionUpdateReq struct {
	g.Meta `path:"/backend/permission/update" method:"post"  desc:"修改权限" tags:"permission"`
	Id     uint `json:"id" v:"required#权限id不能为空" dc:"权限id"`
	PermissionCreateBase
}

type PermissionUpdateRes struct {
	Id uint `json:"id"`
}

//删除权限
type PermissionDeleteReq struct {
	g.Meta `path:"/backend/permission/delete" method:"delete"  desc:"删除权限" tags:"permission" summary:"删除权限"`
	Id     uint `json:"id" v:"required#权限id不能为空" dc:"权限id"`
}

type PermissionDeleteRes struct {
}

//获取权限列表
type PermissionGetListCommonReq struct {
	g.Meta `path:"/backend/permission/list" method:"get" tags:"permission" summary:"获取权限列表"`
	CommonPaginationReq
}

type PermissionGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
