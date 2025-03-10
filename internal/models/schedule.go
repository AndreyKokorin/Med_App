package models

import "time"

type Schedule struct {
	Id          int       `json:"id"`
	DoctorId    int       `json:"doctor_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Capacity    int       `json:"capacity"`
	BookedCount int       `json:"booked_count"`
	Status      string    `json:"status"`
}
