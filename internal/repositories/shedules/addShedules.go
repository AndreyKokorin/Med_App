package repositories

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
)

func AddShedules(shedule models.Schedule) (models.Schedule, error) {
	query := "INSERT INTO schedules( doctor_id, start_time, end_time, capacity) VALUES ($1, $2, $3, $4) RETURNING id, doctor_id, start_time, end_time, capacity"

	var chedule models.Schedule
	err := database.DB.QueryRow(query, shedule.DoctorId, shedule.StartTime, shedule.EndTime, shedule.Capacity).Scan(&chedule.Id, &chedule.DoctorId, &chedule.StartTime, &chedule.EndTime, &chedule.Capacity)
	if err != nil {
		return models.Schedule{}, err
	}

	return chedule, nil
}
