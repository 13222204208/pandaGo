package cmd

import (
	"context"
	"server/internal/router"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			router.InitAdminRouter()
			s := g.Server()
			s.Run()
			return nil
		},
	}
)
