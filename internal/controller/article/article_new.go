// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package article

import (
	"github.com/oldme-git/oldme-api/api/article"
)

type ControllerApp struct{}

func NewApp() article.IArticleApp {
	return &ControllerApp{}
}

type ControllerV1 struct{}

func NewV1() article.IArticleV1 {
	return &ControllerV1{}
}
