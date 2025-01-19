package model

import (
	"time"

	"github.com/google/uuid"
)

type BlogPost struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
