package main

import (
	"github.com/gogf/gf/v2/os/gtime"
	_ "oldme-api/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"oldme-api/internal/cmd"
)

func main() {
	// 全局进程时间为Asia/shanghai
	err := gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.New())
}
