package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"goframe-shop-test/internal/cmd"
	_ "goframe-shop-test/internal/logic"
	_ "goframe-shop-test/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
