package consts

import (
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

var JwtKey []byte

func init() {
	jwtInit()
}

func jwtInit() {
	cfg, _ := gcfg.New()
	Jwt, err := cfg.Get(gctx.New(), "jwtKey")
	if err != nil {
		panic(err)
	}
	JwtKey = Jwt.Bytes()
}
