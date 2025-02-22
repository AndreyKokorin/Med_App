package repositories

import (
	"database/sql"
)

func UserExists(email string, Db *sql.DB) (bool, error) {
	var exists bool
	err := Db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", email).Scan(&exists)
	return exists, err
}
