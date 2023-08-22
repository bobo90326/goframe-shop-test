package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-shop-test/api/backend"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/service"
)

type cFile struct {
}

var File = cFile{}

func (c *cFile) FileUpLoad(ctx context.Context, req *backend.FileUploadReq) (res *backend.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}
	upload, err := service.File().Upload(ctx, model.FileUpLoadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	return &backend.FileUploadRes{
		Name: upload.Name,
		Url:  upload.Url,
	}, nil
}
