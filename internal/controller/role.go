package controller

import (
	"context"
	"goframe-shop-test/api/backend"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/service"
)

var Role = cRole{}

type cRole struct{}

// 新增
func (a *cRole) Create(ctx context.Context, req *backend.RoleReq) (res *backend.RoleRes, err error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleRes{RoleId: out.RoleId}, nil
}
