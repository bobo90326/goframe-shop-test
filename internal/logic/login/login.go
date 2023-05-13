package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-shop-test/internal/dao"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/model/entity"
	"goframe-shop-test/internal/service"
	"goframe-shop-test/utility"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())

}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("账号或密码不正确！！！")
	}
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}
