package model

type SentenceQuery struct {
	Paging
	BookId Id
	// 根据指定id查询
	Ids []Id
}

type SentenceInput struct {
	Id       Id
	BookId   Id
	TagIds   []Id
	Sentence string
}
