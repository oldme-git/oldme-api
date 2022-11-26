package model

type AdminInput struct {
	Username string
	Password string
	Nickname string
	Avatar   string
}

type Login struct {
	Username string
	Password string
}
