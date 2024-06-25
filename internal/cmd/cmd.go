package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/oldme-git/oldme-api/internal/controller/account"
	"github.com/oldme-git/oldme-api/internal/controller/article"
	"github.com/oldme-git/oldme-api/internal/controller/article_grp"
	"github.com/oldme-git/oldme-api/internal/controller/file"
	"github.com/oldme-git/oldme-api/internal/controller/link"
	"github.com/oldme-git/oldme-api/internal/controller/login"
	"github.com/oldme-git/oldme-api/internal/controller/other"
	"github.com/oldme-git/oldme-api/internal/controller/reply"
	"github.com/oldme-git/oldme-api/internal/controller/saying"
	"github.com/oldme-git/oldme-api/internal/logic/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 允许跨域
				// group.Middleware(ghttp.MiddlewareCORS)
				// admin路由
				// group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.Response)
				group.Group("/admin", func(group *ghttp.RouterGroup) {
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Bind(login.NewV1())
					})
					group.Group("/", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.Auth)
						group.Group("/v1", func(group *ghttp.RouterGroup) {
							group.Bind(
								account.NewV1(),
								article.NewV1(),
								article_grp.NewV1(),
								file.NewV1(),
								saying.NewV1(),
								link.NewV1(),
								reply.NewV1(),
							)
						})
					})
				})

				// app路由
				group.Group("/app", func(group *ghttp.RouterGroup) {
					group.Bind(
						article.NewApp(),
						article_grp.NewApp(),
						reply.NewApp(),
						other.NewApp(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
