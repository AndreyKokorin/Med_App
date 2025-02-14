package models

import "time"

type Appointment struct {
	Id              int       `json:"id"`
	PatientId       int       `json:"patient_id" validate:"required"`
	DoctorId        int       `json:"doctor_id" validate:"required"`
	AppointmentTime time.Time `json:"appointment_time" validate:"required"`
	Status          string    `json:"status" validate:"required"`
}
