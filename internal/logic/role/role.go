package role

import (
	"context"
	"goframe-shop-test/internal/dao"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/service"
)

//角色管理
type sRole struct{}

func init() {
	service.RegisterRole(New())
}
func New() *sRole {
	return &sRole{}
}

//创建
func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out *model.RoleCreateOutput, err error) {
	lastInsertId, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return &model.RoleCreateOutput{
		RoleId: int(lastInsertId),
	}, nil
}
