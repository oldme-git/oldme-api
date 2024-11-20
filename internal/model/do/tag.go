// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Tag is the golang structure of table tag for DAO operations like Where/Data.
type Tag struct {
	g.Meta `orm:"table:tag, do:true"`
	Id     interface{} //
	GrpId  interface{} // 分组id
	Name   interface{} // 标签名称
}
