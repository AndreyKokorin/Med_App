package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
)

func GetUserId(db *sql.DB, id int) (models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, email, age, name, roles, date_of_birth, phone_number, address, gender, avatar_url FROM users WHERE id=$1", id).
		Scan(&user.Id, &user.Email, &user.Age, &user.Name, &user.Roles, &user.DateOfBirth, &user.PhoneNumber, &user.Address, &user.Gender, &user.Avatar_url)
	return user, err
}
