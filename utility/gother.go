// 这里面的函数依赖与gf

package utility

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

func Log(v interface{}) {
	l := g.Log()
	l.Info(gctx.New(), v)
}

func GetUrl(r *ghttp.Request) string {
	var proto = "http://"
	if r.TLS != nil {
		proto = "https://"
	}
	return proto + r.Host
}
