package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"oldme-api/internal/cmd"
	_ "oldme-api/internal/logic"
	_ "oldme-api/internal/packed"
)

func main() {
	// 全局进程时间为Asia/shanghai
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")

	cmd.Main.Run(gctx.New())
}
