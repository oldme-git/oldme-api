// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/oldme-git/oldme-api/internal/dao/internal"
)

// internalLinkDao is internal type for wrapping internal DAO implements.
type internalLinkDao = *internal.LinkDao

// linkDao is the data access object for table link.
// You can define custom methods on it to extend its functionality as you wish.
type linkDao struct {
	internalLinkDao
}

var (
	// Link is globally public accessible object for table link operations.
	Link = linkDao{
		internal.NewLinkDao(),
	}
)

// Fill with you ideas below.
