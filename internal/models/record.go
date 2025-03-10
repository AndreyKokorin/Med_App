package models

import "time"

type Record struct {
	Id            int    `json:"id"`
	Patient_id    int    `json:"patient_id" validate:"required"`
	Doctor_id     int    `json:"doctor_id" validate:"required"`
	Diagnosis     string `json:"diagnosis" validate:"required,min=1,max=100"`
	Recomendation string `json:"recomendation" validate:"required,min=1,max=500"`
	CreateTime    string `json:"create_time"`
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
