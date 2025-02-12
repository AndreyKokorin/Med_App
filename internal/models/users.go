package models

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    string `json:"roles"`
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
