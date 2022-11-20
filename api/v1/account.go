package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type LoginReq struct {
	g.Meta   `path:"login" method:"post" sm:"登录" tags:"账户"`
	Username string `v:"required|length: 3,20"`
	Password string `v:"required|length: 8,30"`
}

type LoginRes struct {
	Token string `sm:"凭据" dc:"在需要鉴权的接口中header加入Authorization: token"`
}

type LogoutReq struct {
	g.Meta `path:"logout" method:"post" sm:"登出" tags:"账户"`
	Token  string
}

type LogoutRes struct {
	Result bool
}

type AccountReq struct {
	g.Meta `path:"info" method:"get" sm:"获取登录信息" tags:"账户"`
}

type AccountRes struct {
	Username  string      `json:"username" dc:"用户名"`
	Nickname  string      `json:"nickname" dc:"昵称"`
	Avatar    string      `json:"avatar" dc:"头像"`
	Register  *gtime.Time `json:"register" dc:"注册时间"`
	LastLogin *gtime.Time `json:"lastLogin" dc:"最后登录时间"`
}
