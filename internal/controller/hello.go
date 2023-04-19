package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"goframe-shop-test/api/hello/v1"
)

var (
	Hello = cHello{}
)

type cHello struct {
}

func (c *cHello) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
