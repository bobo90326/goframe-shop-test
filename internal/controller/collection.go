package controller

import (
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop-test/api/frontend"
	"goframe-shop-test/internal/model"
	"goframe-shop-test/internal/service"
	"golang.org/x/net/context"
)

var Collection = cCollection{}

type cCollection struct{}

func (a cCollection) Add(ctx context.Context, req *frontend.AddCollectionReq) (res *frontend.AddCollectionRes, err error) {
	data := model.AddCollectionInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().AddCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.AddCollectionRes{
		Id: out.Id,
	}, nil
}

func (a cCollection) Delete(ctx context.Context, req *frontend.DeleteCollectionReq) (res *frontend.DeleteCollectionRes, err error) {
	data := model.DeleteCollectionInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().DeleteCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeleteCollectionRes{
		Id: out.Id,
	}, nil
}

func (a *cCollection) List(ctx context.Context, req *frontend.ListCollectionReq) (res *frontend.ListCollectionRes, err error) {
	getListRes, err := service.Collection().GetList(ctx, model.CollectionListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}

	return &frontend.ListCollectionRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
