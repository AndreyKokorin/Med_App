package doctors

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/helps"
	"awesomeProject/pkg/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"strings"
)

// UpdateDoctorsData обновляет данные врача (доступно только врачам)
// @Summary Обновление данных врача
// @Description Позволяет врачу обновить свою информацию в системе
// @Tags doctors
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param data body models.DoctorProfile true "Новые данные врача"
// @Success 200 {object} map[string]interface{} "Обновленные данные врача"
// @Failure 400 {object} map[string]string "Ошибка валидации данных"
// @Failure 401 {object} map[string]string "Пользователь не найден или не авторизован"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /doctors/update [put]
func UpdateDoctorsData(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			helps.RespWithError(ctx, http.StatusUnauthorized, "user_id не найден", nil)
		}
	}()

	userId := ctx.MustGet("user_id").(int)

	var doctorNewData models.DoctorProfile
	if err := ctx.ShouldBindJSON(&doctorNewData); err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Error binding data", err)
		return
	}

	err := validate.ValidAndTrim(&doctorNewData)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Error invalid data", err)
		return
	}

	query, args := buildQuery(doctorNewData, userId)

	_, err = database.DB.Exec(query, args...)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Error updating doctor data", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": doctorNewData, "message": "Doctor data updated successfully"})
}

func buildQuery(doctorNewData models.DoctorProfile, userId int) (string, []interface{}) {
	var args []interface{}
	var parts []string
	argsCount := 1

	if doctorNewData.Experience != 0 {
		parts = append(parts, fmt.Sprintf("experience=$%d", argsCount))
		args = append(args, doctorNewData.Experience)
		argsCount++
	}

	if doctorNewData.Education != "" {
		parts = append(parts, (fmt.Sprintf("education=$%d", argsCount)))
		args = append(args, doctorNewData.Education)
		argsCount++
	}

	if doctorNewData.Languages != nil && len(doctorNewData.Languages) > 0 {
		parts = append(parts, (fmt.Sprintf("languages=$%d", argsCount)))
		args = append(args, pq.Array(doctorNewData.Languages))
		argsCount++
	}

	if len(parts) == 0 {
		return "", nil
	}

	query := fmt.Sprintf("UPDATE doctors SET %s WHERE user_id=$%d", strings.Join(parts, ", "), argsCount)
	args = append(args, userId)

	return query, args
}
