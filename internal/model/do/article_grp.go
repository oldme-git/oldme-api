// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleGrp is the golang structure of table article_grp for DAO operations like Where/Data.
type ArticleGrp struct {
	g.Meta      `orm:"table:article_grp, do:true"`
	Id          interface{} //
	Name        interface{} // 名称
	Tags        interface{} // 标签，依英文逗号隔开
	Description interface{} // 简介
	Onshow      interface{} // 是否显示
	Order       interface{} // 排序，越大越靠前
}
