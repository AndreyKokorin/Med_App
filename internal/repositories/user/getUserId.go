package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
)

func GetUserId(db *sql.DB, id int) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, email, age, name FROM users WHERE id=$1", id).Scan(&user.Id, &user.Email, &user.Age, &user.Name)
	return user, err
}
