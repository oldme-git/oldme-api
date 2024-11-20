// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SentenceTag is the golang structure of table sentence_tag for DAO operations like Where/Data.
type SentenceTag struct {
	g.Meta `orm:"table:sentence_tag, do:true"`
	SId    interface{} // 句子id
	TId    interface{} // tag id
}
