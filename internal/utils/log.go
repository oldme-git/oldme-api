package utils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func Log(v interface{}) {
	l := g.Log()
	l.Info(gctx.New(), v)
}
