package models

import "encoding/json"

type Article struct {
	Id     json.Number `json:"id"`
	Title  string      `json:"title"`
	Author string      `json:"author"`
}
type Articles []Article
