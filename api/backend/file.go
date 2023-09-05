package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file" method:"post" mime:"mu" tags:"file" summary:"文件上传"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required#文件不能为空"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url"  dc:"图片地址"`
}
