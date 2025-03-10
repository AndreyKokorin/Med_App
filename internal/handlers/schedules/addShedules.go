package schedules

import (
	"awesomeProject/internal/models"
	repositories "awesomeProject/internal/repositories/shedules"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddSchedules добавляет новое расписание (доступно врачам и администраторам)
// @Summary Добавление расписания
// @Description Добавляет новое расписание в систему (доступно врачам и администраторам)
// @Tags schedules
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param schedule body models.Schedule true "Данные расписания"
// @Success 201 {object} models.Schedule "Созданное расписание"
// @Failure 400 {object} map[string]string "Ошибка валидации запроса"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /schedules [post]
func AddSchedules(ctx *gin.Context) {
	schedule, err := getSchedules(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedule, err = repositories.AddShedules(schedule)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, schedule)
}

func getSchedules(ctx *gin.Context) (models.Schedule, error) {
	var schedule models.Schedule
	if err := ctx.ShouldBindJSON(&schedule); err != nil {
		return models.Schedule{}, err
	}

	return schedule, nil
}
