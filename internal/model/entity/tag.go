// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Tag is the golang structure for table tag.
type Tag struct {
	Id    uint   `json:"id"    orm:"id"     description:""`
	GrpId uint   `json:"grpId" orm:"grp_id" description:"分组id"`
	Name  string `json:"name"  orm:"name"   description:"标签名称"`
}
