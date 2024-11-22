package v1

import "github.com/oldme-git/oldme-api/internal/model"

type List struct {
	Id   model.Id `json:"id"`
	Name string   `json:"name"`
}
