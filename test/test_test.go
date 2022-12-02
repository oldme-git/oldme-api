package test

import (
	"github.com/gogf/gf/v2/os/gctx"
	_ "oldme-api/internal/logic/file"
	"oldme-api/internal/service"
	"testing"
)

func TestA(t *testing.T) {
	ctx := gctx.New()
	a, err := service.File().MoveTmp(ctx, "img", "cor7a3ls2w3gul65yy.png")
	if err != nil {
		panic(err)
	}
	t.Log(a)
	//service.File().MoveTmp(ctx, "img", "cor5s0e9miocwmjbr4.png")
}
