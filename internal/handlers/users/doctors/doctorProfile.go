package doctors

import (
	doctorRep "awesomeProject/internal/repositories/doctor"
	"awesomeProject/pkg/helps"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

const errNotDoctor = "Error not found doctor with this id"

// GetDoctorProfile
// @Summary Получение профиля доктора
// @Description Возвращает профиль доктора по указанному ID
// @Tags doctors
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID доктора"
// @Success 200 {object} map[string]models.DoctorProfile "Профиль доктора"
// @Failure 400 {object} map[string]string "Ошибка валидации"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/doctors/{id}/profile [get]
func GetDoctorProfile(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Error convert user_id", err)
	}

	doctorProfile, err := doctorRep.GetFullDoctorProfile(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helps.RespWithError(ctx, http.StatusBadRequest, "Error not found doctor with this id", err)
			return
		}
		if errors.Is(err, errors.New(errNotDoctor)) {
			helps.RespWithError(ctx, http.StatusBadRequest, "Error not found doctor with this id", err)
			return
		}
		slog.Error(err.Error())
		helps.RespWithError(ctx, http.StatusInternalServerError, "Error query db", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": doctorProfile})
}
