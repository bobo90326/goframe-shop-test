package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop-test/internal/consts"
	"goframe-shop-test/internal/dao"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/model/entity"
	"goframe-shop-test/internal/service"
	"time"
)

type sFile struct {
}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

func (s *sFile) Upload(ctx context.Context, in model.FileUpLoadInput) (out *model.FileUpLoadOutput, err error) {
	upLoadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if upLoadPath == "" {
		return nil, gerror.New("upload.path not found")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	//安全性检查，1分钟只能上传10次
	count, err := dao.FileInfo.Ctx(ctx).Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(time.Minute)).Count()
	if err != nil {
		return nil, err
	}

	if count > consts.FileMaxUploadCountMinute {
		return nil, gerror.New("1分钟只能上传10次")
	}
	dateDirName := gtime.Now().Format("Ymd")
	filename, err := in.File.Save(gfile.Join(upLoadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}
	data := entity.FileInfo{
		Name:   filename,
		Src:    gfile.Join(upLoadPath, dateDirName, filename),
		Url:    "/upload/" + dateDirName + "/" + filename, //和gfile.Join的效果一样
		UserId: gconv.Int(ctx.Value(consts.CtxAdminId)),
	} //保存到数据库
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.FileUpLoadOutput{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url,
	}, nil
}
