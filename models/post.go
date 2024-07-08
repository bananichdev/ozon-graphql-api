package models

import "time"

type Post struct {
	ID               int       `json:"id"`
	Author           string    `json:"author"`
	Title            string    `json:"title"`
	Content          string    `json:"content"`
	CommentsDisabled bool      `json:"comments_disabled"`
	CreatedAt        time.Time `json:"created_at"`
}
