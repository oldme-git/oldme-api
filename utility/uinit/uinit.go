// 此包可用作初始化一些配置

package uinit

import (
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

var SayingTagId uint32
var PoemTagId uint32

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

// poemTagId 获取诗词标签id
// 从配置文件中获取
func poemTagId() {
	cfg, _ := gcfg.New()
	idRaw, err := cfg.Get(gctx.New(), "poemTagId")
	if err != nil {
		PoemTagId = 0
	} else {
		PoemTagId = idRaw.Uint32()
	}
}
