// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReadingDao is the data access object for table reading.
type ReadingDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns ReadingColumns // columns contains all the column names of Table for convenient usage.
}

// ReadingColumns defines and stores column names for table reading.
type ReadingColumns struct {
	Id         string //
	Name       string // 书名
	Author     string // 作者
	Status     string // 状态: 1完结 2在读 3弃读
	FinishedAt string // 读完时间
}

// readingColumns holds the columns for table reading.
var readingColumns = ReadingColumns{
	Id:         "id",
	Name:       "name",
	Author:     "author",
	Status:     "status",
	FinishedAt: "finished_at",
}

// NewReadingDao creates and returns a new DAO object for table data access.
func NewReadingDao() *ReadingDao {
	return &ReadingDao{
		group:   "default",
		table:   "reading",
		columns: readingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ReadingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ReadingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ReadingDao) Columns() ReadingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ReadingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ReadingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ReadingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
