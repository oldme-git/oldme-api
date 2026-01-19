package model

type ReadingStatus uint

const (
	// 10-弃读 15-特殊书类 21-粗读完结 29-正常完结 95-在读
	ReadingStatusAbandoned  ReadingStatus = 10
	ReadingStatusSpecial                  = 15
	ReadingStatusRoughDone                = 21
	ReadingStatusNormalDone               = 29
	ReadingStatusCurrently                = 95
)
