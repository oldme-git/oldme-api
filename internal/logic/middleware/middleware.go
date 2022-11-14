package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"oldme-api/internal/service"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func (s sMiddleware) Before(r *ghttp.Request) {
	ctx := gctx.New()
	l := g.Log()
	l.Info(ctx, "前置中间件")
	r.Middleware.Next()
}

func (s sMiddleware) After(r *ghttp.Request) {
	r.Middleware.Next()
	ctx := gctx.New()
	l := g.Log()
	l.Info(ctx, "后置中间件")
}
