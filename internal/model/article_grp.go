package model

import "oldme-api/internal/model/entity"

type ArticleGrpList []entity.ArticleGrp

type ArticleGrpInput struct {
	Name        string
	Tags        string
	Description string
	Onshow      bool
}
