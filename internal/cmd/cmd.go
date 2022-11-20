package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"oldme-api/internal/controller"
	"oldme-api/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(service.Middleware().Response)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Bind(controller.Account)
				})
				group.Bind(controller.Demo)
			})
			s.Run()
			return nil
		},
	}
)
