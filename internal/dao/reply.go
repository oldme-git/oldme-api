// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/oldme-git/oldme-api/internal/dao/internal"
)

// internalReplyDao is internal type for wrapping internal DAO implements.
type internalReplyDao = *internal.ReplyDao

// replyDao is the data access object for table reply.
// You can define custom methods on it to extend its functionality as you wish.
type replyDao struct {
	internalReplyDao
}

var (
	// Reply is globally public accessible object for table reply operations.
	Reply = replyDao{
		internal.NewReplyDao(),
	}
)

// Fill with you ideas below.
