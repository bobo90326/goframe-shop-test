package main

import (
	"goframe-shop-test/internal/cmd"
	_ "goframe-shop-test/internal/logic"
	_ "goframe-shop-test/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

// 启动///////////
func main() {
	cmd.Main.Run(gctx.New())
}
