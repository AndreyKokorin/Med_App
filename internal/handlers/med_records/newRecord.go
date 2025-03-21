package med_records

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/validate"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

// NewRecord создаёт новую медицинскую запись (для докторов и администраторов)
// @Summary Создание медицинской записи
// @Description Добавляет новую медицинскую запись в базу данных (доступно докторам и администраторам)
// @Tags medical_records
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param record body models.Record true "Данные медицинской записи"
// @Success 201 {object} map[string]int "ID созданной записи"
// @Failure 400 {object} map[string]string "Ошибка валидации"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /doctors/records [post]
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

	slog.Info("Record obj:", newRecord)

	query := "INSERT INTO Medical_Records(patient_id, doctor_id, diagnosis, recomendation, anamnesis, timeslot_id) VALUES ($1,$2,$3,$4, $5) RETURNING id"
	err = database.DB.QueryRow(query, newRecord.Patient_id, newRecord.Doctor_id, newRecord.Diagnosis, newRecord.Recomendation, newRecord.Anamnesis, newRecord.TimeSlotsId).Scan(&newRecord.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"record": newRecord, "massage": "Med record created successfully!"})
}
