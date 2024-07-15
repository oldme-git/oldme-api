// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Reading is the golang structure of table reading for DAO operations like Where/Data.
type Reading struct {
	g.Meta     `orm:"table:reading, do:true"`
	Id         interface{} //
	Name       interface{} // 书名
	Author     interface{} // 作者
	Status     interface{} // 状态: 1完结 2在读 3弃读
	FinishedAt *gtime.Time // 读完时间
}
