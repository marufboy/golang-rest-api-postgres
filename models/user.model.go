package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primary_key"`
	Name     string    `gorm:"type: varchar(255); not null"`
	Email    string    `gorm:"uniqueIndex; not null"`
	Password string    `gorm:"not null"`
	Photo    string    `gorm:"not null"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required, min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Photo           string `json:"photo" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name"`
	Photo string    `json:"photo,omitempty"`
}
