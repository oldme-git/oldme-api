// 此包可用作初始化一些配置

package uinit

import (
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

var SayingTagId uint32

func init() {
	sayingTagId()
}

// sayingTagId 获取短句标签id，用于首页展示
// 从配置文件中获取
func sayingTagId() {
	cfg, _ := gcfg.New()
	idRaw, err := cfg.Get(gctx.New(), "sayingTagId")
	if err != nil {
		SayingTagId = 0
	} else {
		SayingTagId = idRaw.Uint32()
	}
}
