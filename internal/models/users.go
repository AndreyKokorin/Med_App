package models

import "time"

type User struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Roles       string    `json:"roles"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Gender      string    `json:"gender"`
	DateOfBirth string    `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Avatar_url  string    `json:"avatar_url"`
}

type UserDetails struct {
	Gender      string    `json:"gender" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Address     string    `json:"address" binding:"required"`
}

type LogUpUser struct {
	Name      string `json:"name"  validate:"required,min=2"`
	Age       int    `json:"age" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6" `
	Roles     string
	RoleToken string `json:"role_token"`
}

type LogInUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUser struct {
	Name        string `json:"name" validate:"omitempty,min=2"`
	Email       string `json:"email" validate:"omitempty,email"`
	Age         int    `json:"age" validate:"omitempty,min=0"`
	Gender      string `json:"gender" validate:"omitempty,oneof=male female other"`
	DateOfBirth string `json:"date_of_birth" validate:"omitempty"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,min=10"`
	Address     string `json:"address" validate:"omitempty,min=5"`
}
type ChangeData struct {
	Code        string `json:"code"`
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
}
