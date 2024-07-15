// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Reading is the golang structure for table reading.
type Reading struct {
	Id         uint        `json:"id"         orm:"id"          description:""`
	Name       string      `json:"name"       orm:"name"        description:"书名"`
	Author     string      `json:"author"     orm:"author"      description:"作者"`
	Status     string      `json:"status"     orm:"status"      description:"状态: 1完结 2在读 3弃读"`
	FinishedAt *gtime.Time `json:"finishedAt" orm:"finished_at" description:"读完时间"`
}
