package role

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"goframe-shop-test/internal/dao"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/model/entity"
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

//删除
func (s *sRole) Delete(ctx context.Context, id uint) error {
	_, err := dao.RoleInfo.Ctx(ctx).Where(g.Map{dao.RoleInfo.Columns().Id: id}).Unscoped().Delete()
	return err
}

//修改
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) error {
	_, err := dao.RoleInfo.Ctx(ctx).
		Data(in).
		FieldsEx(dao.RoleInfo.Columns().Id).
		Where(g.Map{dao.RoleInfo.Columns().Id: in.Id}).
		Update()
	return err
}

//获取角色列表
func (s *sRole) GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error) {
	var (
		m = dao.RoleInfo.Ctx(ctx)
	)
	out = &model.RoleGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.RoleInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	//不指定item的键名用：Scan
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
