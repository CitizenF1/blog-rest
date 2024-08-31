package models

import "time"

type Post struct {
	ID        int       `json:"id"`
	Subject   string    `json:"subject"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
}
