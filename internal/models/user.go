package models

import "time"

type UserDTO struct {
}

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	LastModifiedAt time.Time `json:"last_modified_at"`
}
