package models

import "time"

type DoctorProfile struct {
	ID         int      `json:"id"`
	UserID     int      `json:"user_id"`
	Specialty  string   `json:"specialty" validate:"max=100"`
	Experience int      `json:"experience" validate:"gte=0"`
	Education  string   `json:"education" validate:"max=100"`
	Languages  []string `json:"languages"`
}

type FullDoctorProfile struct {
	UserID      int       `json:"user_id"`
	Specialty   string    `json:"specialty" validate:"max=100"`
	Experience  int       `json:"experience" validate:"gte=0"`
	Education   string    `json:"education" validate:"max=100"`
	Languages   []string  `json:"languages"`
	Name        string    `json:"name"`
	Age         int       `json:"age"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Roles       string    `json:"roles"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Avatar_url  string    `json:"avatar_url"`
}
