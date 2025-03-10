package med_records

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRecordId получает медицинскую запись по ID (для докторов и администраторов)
// @Summary Получение медицинской записи
// @Description Возвращает медицинскую запись по её ID (доступно докторам и администраторам)
// @Tags medical_records
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID медицинской записи"
// @Success 200 {object} models.Record "Данные медицинской записи"
// @Failure 400 {object} map[string]string "Ошибка валидации запроса"
// @Failure 404 {object} map[string]string "Запись не найдена"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /doctors/records/{id} [get]
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
