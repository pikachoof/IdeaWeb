package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type LoginRequest struct {
	Email	 string `json:"email" binding:"reequired,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role string `json:"role" binding:"oneof=user admin"`
}

type AuthResponse struct {
	User *User `json:"user"`
	AccessToken string `json:"access_token"`
	ExpiresAt int64 `json:"expires_at"`
}
