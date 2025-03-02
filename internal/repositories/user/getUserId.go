package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
)

func GetUserId(db *sql.DB, id int) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, email, age, roles FROM users WHERE id=$1", id).Scan(&user.Id, &user.Email, &user.Age, &user.Name, &user.Roles)
	return user, err
}
