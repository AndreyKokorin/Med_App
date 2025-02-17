package models

type Record struct {
	Id            int    `json:"id"`
	Patient_id    int    `json:"patient_id" validate:"required"`
	Doctor_id     int    `json:"doctor_id" validate:"required"`
	Diagnosis     string `json:"diagnosis" validate:"required,min=1,max=100"`
	Recomendation string `json:"recomendation" validate:"required,min=1,max=500"`
	CreateTime    string `json:"create_time"`
}
