package main

import (
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/oldme-git/oldme-api/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/oldme-git/oldme-api/internal/logic"
	_ "github.com/oldme-git/oldme-api/internal/utility"
	_ "github.com/oldme-git/oldme-api/utility/uinit"
)

const (
	VERSION = "v1.6.0"
)

func main() {
	var err error

	fmt.Println(VERSION)

	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")

	// 全局进程时间为Asia/shanghai
	err = gtime.SetTimeZone("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	// 检查数据库是否能连接
	err = connData()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.New())
}

// 检查数据库连接是否正常
func connData() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("连接到数据库失败")
	}
	return nil
}
