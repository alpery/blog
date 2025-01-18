package model

import "github.com/google/uuid"

type BlogPost struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}
