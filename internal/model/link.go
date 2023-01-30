package model

type LinkInput struct {
	Name        string `json:"name" v:"required|length:1, 20"`
	Description string `json:"description" v:"length:2, 200"`
	Link        string `json:"link" v:"required|length:1, 200"`
}
