package repositories

import (
	"awesomeProject/internal/models"
	"database/sql"
)

func NewUser(logUpdata models.LogUpUser, db *sql.DB, hashPassword string) error {
	query := "INSERT INTO users(name, age, email, password, roles) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(query, logUpdata.Name, logUpdata.Age, logUpdata.Email, hashPassword, logUpdata.Roles)

	if err != nil {
		return err
	}

	return nil
}
