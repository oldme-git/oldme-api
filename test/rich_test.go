package test

import (
	"github.com/gogf/gf/v2/os/gctx"
	_ "oldme-api/internal/logic/file"
	_ "oldme-api/internal/logic/rich"
	"oldme-api/internal/service"
	"testing"
)

func TestEdit(t *testing.T) {
	var (
		oldText = "http://127.0.0.1:8001/static/rich/202212/33.pnghttp://127.0.0.1:8001/static/rich/202212/44.png"
		newText = ""
	)
	service.Rich().Edit(gctx.New(), &oldText, &newText)
	t.Log(newText)
}
