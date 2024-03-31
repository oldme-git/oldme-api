package packed

import (
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type pErr struct {
	maps map[int]string
}

const CodeOk = 0
const CodeErrSys = 99999

var Err = &pErr{
	maps: map[int]string{
		CodeOk:     "好耶",
		CodeErrSys: "未知错误",
		10100:      "sorry,数据走丢了",
		10101:      "你所在文章分类不存在",
		10201:      "你所在的文章不存在",
		10301:      "网站地址需要http协议奥",
		10302:      "回复的内容不在同一文章下",
		10303:      "你要回复的内容不见了",
		10501:      "图片类型不正确",
		10503:      "文件不存在",
		20100:      "你的账号或者密码不对哩",
		20101:      "不正确的或者过期的token，请重新获取",
		20102:      "┗|｀O′|┛ 嗷~~，获取你信息失败了",
	},
}

// GetMsg 获取code码对应的msg
func (c *pErr) GetMsg(code int) string {
	return c.maps[code]
}

// Skip 抛出一个业务级别的错误，不会打印错误堆栈信息
func (c *pErr) Skip(code int, msg ...string) (err error) {
	var msgStr string
	if len(msg) == 0 {
		msgStr = c.GetMsg(code)
	} else {
		msg = append([]string{c.GetMsg(code)}, msg...)
		msgStr = strings.Join(msg, ", ")
	}
	return gerror.NewWithOption(gerror.Option{
		Stack: false,
		Text:  msgStr,
		Code:  gcode.New(code, "", nil),
	})
}

// Sys 抛出一个系统级别的错误，使用特殊的code码：99999
// !!! 使用该方法时，它会打印错误堆栈信息到日志，但是一定不要把任何错误信息抛出到客户端，防止泄露系统信息
// !!! 推荐做法是在后置中间件中捕获 code 99999 的错误，然后返回给客户端一个统一的错误提示
func (c *pErr) Sys(err error) error {
	return gerror.NewCode(gcode.New(CodeErrSys, "", nil), err.Error())
}

// GetOkMsg 获取成功的msg
func (c *pErr) GetOkMsg() string {
	return c.GetMsg(CodeOk)
}

// GetSysMsg 获取系统错误的msg
func (c *pErr) GetSysMsg() string {
	return c.GetMsg(CodeErrSys)
}
