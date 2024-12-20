package model

type AdminInput struct {
	Username string
	Password string
	Nickname string
	Avatar   string
}

type LoginInput struct {
	Username string
	Password string
}
