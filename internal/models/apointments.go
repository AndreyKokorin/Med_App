package models

import "time"

type Appointment struct {
	Id              int       `json:"id"`
	PatientId       int       `json:"patient_id"`
	DoctorId        int       `json:"doctor_id"`
	AppointmentTime time.Time `json:"appointment_time"`
	Status          string    `json:"status"`
	ShedulesID      int       `json:"shedules_id"`
}

type AppointmentsUpdateTime struct {
	AppointmentTime time.Time `json:"appointment_time"`
}
