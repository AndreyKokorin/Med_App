package med_records

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecordId(ctx *gin.Context) {
	userid := ctx.Param("id")
	if userid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "userid is required"})
		return
	}

	var record models.Record
	query := "SELECT id, patient_id,doctor_id,diagnosis,recomendation,created_time FROM Medical_Records WHERE id = $1"
	err := database.DB.QueryRow(query, userid).Scan(&record.Id, &record.Patient_id, &record.Doctor_id, &record.Diagnosis, &record.Recomendation, &record.CreateTime)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, record)
}
