package packed

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

type oCode struct {
}

var Oldme oCode

// GetMsg 获取code码对应的msg
func (o oCode) GetMsg(code int) string {
	var maps map[int]string
	maps = map[int]string{
		0:  "好耶",
		10: "你的请求参数不对哩",
	}
	return maps[code]
}

// SetErr 设置一个业务码的err
func (o oCode) SetErr(code int, msg ...string) (err error) {
	var text string
	if len(msg) == 0 {
		text = o.GetMsg(code)
	} else {
		text = strings.Join(msg, ",")
	}
	err = gerror.NewCode(gcode.New(code, "", nil), text)
	return
}
