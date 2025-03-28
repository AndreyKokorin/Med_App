package models

import "time"

type Direction struct {
	ID               int       `json:"id"`
	DoctorID         int       `json:"doctor_id"`
	PatientID        int       `json:"patient_id" binding:"required"`
	ExaminationType  int       `json:"examination_type" binding:"required"`
	ExecutorDoctorID int       `json:"executor_doctor_id" binding:"required"`
	CreatedAt        time.Time `json:"created_at"`
	Status           string    `json:"status"`
}

type DirectResult struct {
	Id          int       `json:"id"`
	DoctorID    int       `json:"doctor_id"`
	DirectionId int       `json:"direction_id"`
	CreatedAt   time.Time `json:"created_at"`
	FilePath    string    `json:"file_path"`
}
