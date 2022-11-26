package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"oldme-api/internal/controller/admin"
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
				group.Group("/admin", func(group *ghttp.RouterGroup) {
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Bind(admin.Login)
					})
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(service.Middleware().Auth)
						group.Group("/v1", func(group *ghttp.RouterGroup) {
							group.Bind(admin.Account)
							group.Bind(admin.ArticleGrp)
						})
					})
				})
			})
			s.Run()
			return nil
		},
	}
)
