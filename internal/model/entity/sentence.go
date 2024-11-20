// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Sentence is the golang structure for table sentence.
type Sentence struct {
	Id       uint   `json:"id"       orm:"id"       description:""`
	BookId   uint   `json:"bookId"   orm:"book_id"  description:""`
	Sentence string `json:"sentence" orm:"sentence" description:""`
}
