package test

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/oldme-git/oldme-api/internal/logic/file"
	"github.com/oldme-git/oldme-api/internal/logic/rich"
	_ "github.com/oldme-git/oldme-api/internal/logic/rich"
)

func TestEdit(t *testing.T) {
	var (
		oldText = "http://127.0.0.1:8001/static/rich/202212/33.pnghttp://127.0.0.1:8001/static/rich/202212/44.png"
		newText = ""
	)
	rich.Edit(gctx.New(), &oldText, &newText)
	t.Log(newText)
}

func TestDel(t *testing.T) {
	var text = "http://127.0.0.1:8001/static/rich/202212/cp1b74s25srkibxxwu.pnghttp://127.0.0.1:8001/static/rich/202212/cp1b77mc9nycaffuo9.png"
	rich.Del(gctx.New(), &text)
}
