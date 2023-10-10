package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type LogoutReq struct {
	g.Meta `path:"logout" method:"post" sm:"登出" tags:"账户"`
	Token  string
}

type LogoutRes struct {
	Result bool
}

type InfoReq struct {
	g.Meta `path:"info" method:"get" sm:"获取登录信息" tags:"账户"`
}

type InfoRes struct {
	Username  string      `json:"username" dc:"用户名"`
	Nickname  string      `json:"nickname" dc:"昵称"`
	Avatar    string      `json:"avatar" dc:"头像"`
	Register  *gtime.Time `json:"register" dc:"注册时间"`
	LastLogin *gtime.Time `json:"lastLogin" dc:"最后登录时间"`
}
