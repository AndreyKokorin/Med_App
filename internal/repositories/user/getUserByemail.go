package repositories

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := "SELECT id, roles,password FROM users WHERE email=$1"
	err := database.DB.QueryRow(query, email).Scan(&user.Id, &user.Roles, &user.Password)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
