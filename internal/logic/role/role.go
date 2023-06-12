package role

import (
	"context"
	"goframe-shop-test/internal/dao"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())

}

func New() *sRole {
	return &sRole{}
}

// Create 创建内容
func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	// 不允许HTML代码
	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: int(lastInsertID)}, err
}
