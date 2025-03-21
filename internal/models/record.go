package models

import "time"

type Record struct {
	Id            int    `json:"id"`
	Patient_id    int    `json:"patient_id" validate:"required"`
	Doctor_id     int    `json:"doctor_id" validate:"required"`
	Diagnosis     string `json:"diagnosis" validate:"required"`
	Recomendation string `json:"recomendation" validate:"required"`
	CreateTime    string `json:"create_time" validate:"required"`
	TimeSlotsId   int    `json:"time_slots" validate:"required"`
	Anamnesis     string `json:"anamnesis" validate:"required"`
	SlotId        int    `json:"timeslot_id" validate:"required"`
}

type MinimalUserData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FullMedicalRecord struct {
	ID             int             `json:"id"`
	Patient        MinimalUserData `json:"patient"`
	Doctor         MinimalUserData `json:"doctor"`
	Diagnosis      string          `json:"diagnosis"`
	Recommendation string          `json:"recomendation"`
	CreatedTime    time.Time       `json:"created_time"`
}
