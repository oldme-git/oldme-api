// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/oldme-git/oldme-api/internal/dao/internal"
)

// internalSentenceTagDao is internal type for wrapping internal DAO implements.
type internalSentenceTagDao = *internal.SentenceTagDao

// sentenceTagDao is the data access object for table sentence_tag.
// You can define custom methods on it to extend its functionality as you wish.
type sentenceTagDao struct {
	internalSentenceTagDao
}

var (
	// SentenceTag is globally public accessible object for table sentence_tag operations.
	SentenceTag = sentenceTagDao{
		internal.NewSentenceTagDao(),
	}
)

// Fill with you ideas below.
