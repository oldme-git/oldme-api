// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/oldme-git/oldme-api/internal/dao/internal"
)

// internalSayingDao is internal type for wrapping internal DAO implements.
type internalSayingDao = *internal.SayingDao

// sayingDao is the data access object for table saying.
// You can define custom methods on it to extend its functionality as you wish.
type sayingDao struct {
	internalSayingDao
}

var (
	// Saying is globally public accessible object for table saying operations.
	Saying = sayingDao{
		internal.NewSayingDao(),
	}
)

// Fill with you ideas below.
