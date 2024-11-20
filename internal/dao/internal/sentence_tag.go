// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SentenceTagDao is the data access object for table sentence_tag.
type SentenceTagDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns SentenceTagColumns // columns contains all the column names of Table for convenient usage.
}

// SentenceTagColumns defines and stores column names for table sentence_tag.
type SentenceTagColumns struct {
	SId string // 句子id
	TId string // tag id
}

// sentenceTagColumns holds the columns for table sentence_tag.
var sentenceTagColumns = SentenceTagColumns{
	SId: "s_id",
	TId: "t_id",
}

// NewSentenceTagDao creates and returns a new DAO object for table data access.
func NewSentenceTagDao() *SentenceTagDao {
	return &SentenceTagDao{
		group:   "default",
		table:   "sentence_tag",
		columns: sentenceTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SentenceTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SentenceTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SentenceTagDao) Columns() SentenceTagColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SentenceTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SentenceTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SentenceTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
