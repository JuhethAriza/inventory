package dto

import "time"

type UserDTO struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}
