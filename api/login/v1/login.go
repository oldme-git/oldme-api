package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"login" method:"post" sm:"登录" tags:"账户"`
	Username string `v:"required|length: 3,20"`
	Password string `v:"required|length: 8,30"`
}

type LoginRes struct {
	Token string `sm:"凭据" dc:"在需要鉴权的接口中header加入Authorization: token"`
}
