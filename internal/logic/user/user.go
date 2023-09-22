package user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"goframe-shop-test/internal/consts"
	"goframe-shop-test/internal/dao"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/model/do"
	"goframe-shop-test/internal/service"
	"goframe-shop-test/utility"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Register 添加用户
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	//插入数据返回id
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

// 修改密码
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(ctx.Value(consts.CtxUserId)).Scan(&userInfo)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}

	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
		g.Dump("userInfo.SecretAnswer:", userInfo.SecretAnswer)
		g.Dump("in.SecretAnswer:", in.SecretAnswer)
		return out, errors.New(consts.ErrSecretAnswerMsg)
	}

	//处理加密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)

	//_, err = dao.UserInfo.Ctx(ctx).WherePri(ctx.Value(consts.CtxUserId)).Data(in).Update()
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	return model.UpdatePasswordOutput{Id: userId}, err
}
