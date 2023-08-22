package model

import "github.com/gogf/gf/v2/net/ghttp"

type FileUpLoadInput struct {
	File       *ghttp.UploadFile
	Name       string
	RandomName bool
}

type FileUpLoadOutput struct {
	Id   uint
	Name string
	Src  string
	Url  string
}

type FileDeleteInput struct {
	Url string
}
