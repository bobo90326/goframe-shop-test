package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleReq struct {
	g.Meta `path:"/article/add" tags:"文章" method:"post" summary:"添加文章"`
	ArticleCommonAddUpdate
}

type ArticleCommonAddUpdate struct {
	//UserId  uint   `json:"user_id" v:"required#用户id必填" dc:"用户id"`
	Title   string `json:"title" v:"required#标题必填" dc:"标题"`
	PicUrl  string `json:"pic_url"           dc:"图片"`
	Desc    string `json:"desc"              description:"文章概要描述"`
	IsAdmin uint   `d:"1" dc:"1 后台管理员发布 2前台用户发布"`
	Detail  string `json:"detail" v:"required#文章详情必填" dc:"文章详情"`
	Praise  uint   `json:"praise" v:"required#点赞数量必填" dc:"点赞数量"`
}

type ArticleRes struct {
	Id uint `json:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" method:"delete" tags:"文章" summary:"删除文章接口"`
	Id     uint `v:"min:1#请选择需要删除的文章" dc:"文章id"`
}
type ArticleDeleteRes struct{}

type ArticleUpdateReq struct {
	g.Meta `path:"/article/update/" method:"post" tags:"文章" summary:"修改文章接口"`
	Id     uint `json:"id"      v:"min:1#请选择需要修改的文章" dc:"文章Id"`
	ArticleCommonAddUpdate
}
type ArticleUpdateRes struct {
	Id uint `json:"id"`
}
type ArticleGetListCommonReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"文章" summary:"文章列表接口"`
	CommonPaginationReq
}
type ArticleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
