package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"oldme-api/internal/cmd"
	_ "oldme-api/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
}
