// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta      `orm:"table:article, do:true"`
	Id          interface{} //
	GrpId       interface{} // 分组id
	Title       interface{} // 标题
	Author      interface{} // 作者
	Thumb       interface{} // 图片地址
	Tags        interface{} // 标签，依英文逗号隔开
	Description interface{} // 简介
	Content     interface{} // 内容
	Order       interface{} // 排序，越大越靠前
	Ontop       interface{} // 是否置顶
	Onshow      interface{} // 是否显示
	Hist        interface{} // 点击数
	Post        interface{} // 评论数
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	LastedAt    *gtime.Time // 最后浏览时间
}
