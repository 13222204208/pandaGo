package main

import (
	"github.com/gogf/gf/v2/frame/g"
	_ "server/internal/packed"

	_ "server/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"
	"server/internal/cmd"
)

func main() {
	g.I18n().SetLanguage("cn")
	cmd.Main.Run(gctx.GetInitCtx())
}
