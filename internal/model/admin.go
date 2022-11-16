package model

type AdminInput struct {
	Username string // 用户名
	Password string // 密码
	Nickname string // 昵称
	Avatar   string // 头像，base64
}

type Login struct {
	Username string
	Password string
}
