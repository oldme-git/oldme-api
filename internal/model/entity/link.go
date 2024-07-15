// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Link is the golang structure for table link.
type Link struct {
	Id          uint   `json:"id"          orm:"id"          description:""`
	Name        string `json:"name"        orm:"name"        description:"名称"`
	Description string `json:"description" orm:"description" description:"描述"`
	Link        string `json:"link"        orm:"link"        description:"链接"`
}
