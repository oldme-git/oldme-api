// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta    `orm:"table:admin, do:true"`
	Id        interface{} //
	Username  interface{} // 用户名
	Password  interface{} // 密码
	Nickname  interface{} // 昵称
	Avatar    interface{} // 头像，base64
	Register  *gtime.Time // 注册时间
	Salt      interface{} // 随机盐值
	LastLogin *gtime.Time // 最后登录时间
}
