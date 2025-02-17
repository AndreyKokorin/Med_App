package med_records

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRecord(ctx *gin.Context) {
	var newRecord models.Record
	if err := ctx.ShouldBindJSON(&newRecord); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate.ValidAndTrim(&newRecord)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var recordId int
	query := "INSERT INTO Medical_Records(patient_id, doctor_id, diagnosis, recomendation) VALUES ($1,$2,$3,$4) RETURNING id"
	err = database.DB.QueryRow(query, newRecord.Patient_id, newRecord.Doctor_id, newRecord.Diagnosis, newRecord.Recomendation).Scan(&recordId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"recordId": recordId})
}
