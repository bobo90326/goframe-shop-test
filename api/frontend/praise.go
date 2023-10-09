package frontend

import "github.com/gogf/gf/v2/frame/g"

type AddPraiseReq struct {
	g.Meta   `path:"/add/praise" tags:"前台点赞" method:"post" summary:"添加点赞"`
	ObjectId uint  `json:"object_id"  description:"对象id" v:"required#点赞对象id不能数为空"`
	Type     uint8 `json:"type"      description:"点赞类型：1商品 2文章" v:"in:1,2"`
}

type AddPraiseRes struct {
	Id uint `json:"id"`
}

// 删除点赞
type DeletePraiseReq struct {
	g.Meta   `path:"/delete/praise" tags:"前台点赞" method:"post" summary:"删除点赞"`
	Id       uint  `json:"id"    `
	ObjectId uint  `json:"object_id"  `
	Type     uint8 `json:"type"      `
}

type DeletePraiseRes struct {
	Id uint `json:"id"`
}

type ListPraiseReq struct {
	g.Meta `path:"/praise/list" method:"post" tags:"前台点赞" summary:"点赞列表"`
	Type   uint8 `json:"type"  v:"in:1,2" description:"点赞类型：1商品 2文章"`
	CommonPaginationReq
}

type ListPraiseRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type ListPraiseItem struct {
	Id       int         `json:"id"        description:""`
	UserId   int         `json:"user_id"    description:"用户id"`
	ObjectId int         `json:"object_id"  description:"对象id"`
	Type     int         `json:"type"      description:"点赞类型：1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
