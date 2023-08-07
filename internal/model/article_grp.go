package model

import "oldme-api/internal/model/entity"

type ArticleGrpInput struct {
	Name        string `v:"required|length:2, 30"`
	Tags        string `v:"length:1, 200"`
	Description string `v:"length:2, 200"`
	Onshow      bool   `v:"required"`
	Order       int    `json:"order" v:"integer|between:-9999,9999"`
}

type ArticleGrpListSafe struct {
	entity.ArticleGrp
	ArticleCount uint `json:"article_count"`
	Onshow       Out  `json:"onshow,omitempty"`
	Order        Out  `json:"order,omitempty"`
}
