package med_records

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserRecords(ctx *gin.Context) {
	userid := ctx.Param("id")
	if userid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "userid is required"})
		return
	}

	var userRecords []models.Record
	query := "SELECT patient_id,doctor_id,diagnosis,recomendation,created_time FROM Medical_Records WHERE patient_id = $1"
	rows, err := database.DB.Query(query, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var record models.Record
		err := rows.Scan(&record.Patient_id, &record.Doctor_id, &record.Diagnosis, &record.Recomendation, &record.CreateTime)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		userRecords = append(userRecords, record)
	}

	ctx.JSON(http.StatusOK, userRecords)
}
