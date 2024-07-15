// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ArticleGrp is the golang structure for table article_grp.
type ArticleGrp struct {
	Id          uint   `json:"id"          orm:"id"          description:""`
	Name        string `json:"name"        orm:"name"        description:"名称"`
	Tags        string `json:"tags"        orm:"tags"        description:"标签，依英文逗号隔开"`
	Description string `json:"description" orm:"description" description:"简介"`
	Onshow      uint   `json:"onshow"      orm:"onshow"      description:"是否显示"`
	Order       int    `json:"order"       orm:"order"       description:"排序，越大越靠前"`
}
