package repositories

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"fmt"
)

func UpdateUserDetails(data models.UserDetails, userID int) error {
	query := "UPDATE users SET gender=$1, date_of_birth=$2, phone_number=$3, address=$4 WHERE id=$5"

	res, err := database.DB.Exec(query, data.Gender, data.DateOfBirth, data.PhoneNumber, data.Address, userID)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", userID)
	}

	return nil
}
