// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Saying is the golang structure for table saying.
type Saying struct {
	Id     uint   `json:"id"     orm:"id"     description:""`
	Saying string `json:"saying" orm:"saying" description:""`
	Author string `json:"author" orm:"author" description:""`
	Source string `json:"source" orm:"source" description:""`
}
