package v1

import "github.com/gogf/gf/v2/frame/g"

type AdminLoginReq struct {
	g.Meta   `path:"/admin/login" method:"post" sm:"管理员登录" tags:"登录"`
	Username string `v:"required|length: 3,20"`
	Password string `v:"required|length: 8,30"`
}

type AdminLoginRes struct {
	Result bool
}
