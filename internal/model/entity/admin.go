// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id        uint        `json:"id"        orm:"id"         description:""`
	Username  string      `json:"username"  orm:"username"   description:"用户名"`
	Password  string      `json:"password"  orm:"password"   description:"密码"`
	Nickname  string      `json:"nickname"  orm:"nickname"   description:"昵称"`
	Avatar    string      `json:"avatar"    orm:"avatar"     description:"头像，base64"`
	Register  *gtime.Time `json:"register"  orm:"register"   description:"注册时间"`
	Salt      string      `json:"salt"      orm:"salt"       description:"随机盐值"`
	LastLogin *gtime.Time `json:"lastLogin" orm:"last_login" description:"最后登录时间"`
}
