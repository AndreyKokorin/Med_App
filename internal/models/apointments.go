package models

import "time"

type Appointment struct {
	Id          int    `json:"id" `
	PatientId   int    `json:"patient_id" binding:"required"`
	Status      string `json:"status" default:"Pending"`
	Slot_id     int    `json:"slot_id" binding:"required"`
	Schedule_id int    `json:"schedule_id" binding:"required"`
}

type AppointmentsUpdateTime struct {
	AppointmentTime time.Time `json:"appointment_time"`
}

type AppointmentDetail struct {
	AppointmentID     int       `json:"appointment_id"`
	PatientID         int       `json:"patient_id"`
	AppointmentStatus string    `json:"appointment_status"`
	SlotID            int       `json:"slot_id"`
	SlotStartTime     time.Time `json:"slot_start_time"`
	SlotStatus        string    `json:"slot_status"`
	ScheduleID        int       `json:"schedule_id"`
	DoctorID          int       `json:"doctor_id"`
	ScheduleStartTime time.Time `json:"schedule_start_time"`
	ScheduleEndTime   time.Time `json:"schedule_end_time"`
	Capacity          int       `json:"capacity"`
	BookedSlots       int       `json:"booked_slots"`
	ScheduleStatus    string    `json:"schedule_status"`
}
