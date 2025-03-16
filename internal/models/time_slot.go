package models

import "time"

type TimeSlot struct {
	Id         int       `json:"id" binding:"required"`
	DoctorId   int       `json:"doctor_id" binding:"required"`
	StartTime  time.Time `json:"start_time" binding:"required"`
	EndTime    time.Time `json:"end_time" binding:"required"`
	ScheduleId int       `json:"schedule_id" binding:"required"`
	Status     string    `json:"status" binding:"required"`
}
