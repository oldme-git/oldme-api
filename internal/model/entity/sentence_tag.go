// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SentenceTag is the golang structure for table sentence_tag.
type SentenceTag struct {
	SId uint `json:"sId" orm:"s_id" description:"句子id"`
	TId uint `json:"tId" orm:"t_id" description:"tag id"`
}
