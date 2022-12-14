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

func TestDel(t *testing.T) {
	var text = "http://127.0.0.1:8001/static/rich/202212/cp1b74s25srkibxxwu.pnghttp://127.0.0.1:8001/static/rich/202212/cp1b77mc9nycaffuo9.png"
	service.Rich().Del(gctx.New(), &text)
}
