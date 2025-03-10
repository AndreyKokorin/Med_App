package med_records

import (
	repositories "awesomeProject/internal/repositories/medical_records"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserRecords
// @Summary Получение медицинских записей пользователя
// @Description Возвращает список медицинских записей по ID пользователя
// @Tags medical_records
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "Идентификатор пользователя"
// @Success 200 {array} models.Record "Список медицинских записей"
// @Failure 400 {object} map[string]string "Ошибка валидации данных"
// @Failure 404 {object} map[string]string "Медицинские записи не найдены"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/users/{id}/records [get]
func GetUserRecords(ctx *gin.Context) {
	userid := ctx.Param("id")

	if userid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "userid is required"})
		return
	}

	userRecords, err := repositories.GetMedicalRecordsUserId(userid)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"err": "medical records not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "failed to get records"})
		return
	}

	ctx.JSON(http.StatusOK, userRecords)
}
