package schedules

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/helps"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetFilterSchedules
// @Summary Получение расписаний с фильтрацией
// @Description Позволяет получить список расписаний с возможностью фильтрации по врачу, дате, времени и статусу
// @Tags schedules
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param doctor_id query int false "ID врача"
// @Param date query string false "Дата в формате YYYY-MM-DD"
// @Param time query string false "Время в формате HH:MM:SS"
// @Param status query string false "Статус расписания (active, archived)"
// @Success 200 {object} map[string]interface{} "Список найденных расписаний"
// @Failure 400 {object} map[string]string "Ошибка валидации параметров"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /schedules/filter [get]
func GetFilterSchedules(ctx *gin.Context) {
	doctorId := ctx.Query("doctor_id")
	date := ctx.Query("date")
	timeQuery := ctx.Query("time")
	status := ctx.Query("status")

	query := strings.Builder{}
	query.WriteString("select id, doctor_id, start_time, end_time, capacity, booked_slots, status from schedules where 1 = 1")

	var parameters []interface{}
	count := 1

	// Валидация doctorId (должен быть числом)
	if doctorId != "" {
		doctorIdInt, err := strconv.Atoi(doctorId)
		if err != nil {
			helps.RespWithError(ctx, http.StatusBadRequest, "Invalid doctor_id format, must be an integer", err)
			return
		}
		query.WriteString(" AND doctor_id = $" + strconv.Itoa(count))
		parameters = append(parameters, doctorIdInt)
		count++
	}

	// Валидация date (должен быть в формате YYYY-MM-DD)
	if date != "" {
		if _, err := time.Parse("2006-01-02", date); err != nil {
			helps.RespWithError(ctx, http.StatusBadRequest, "Invalid date format, expected YYYY-MM-DD", err)
			return
		}
		query.WriteString(" AND DATE(start_time) = $" + strconv.Itoa(count))
		parameters = append(parameters, date)
		count++
	}

	// Валидация time (должен быть в формате HH:MM:SS)
	if timeQuery != "" {
		if _, err := time.Parse("15:04:05", timeQuery); err != nil {
			helps.RespWithError(ctx, http.StatusBadRequest, "Invalid time format, expected HH:MM:SS", err)
			return
		}
		query.WriteString(" AND start_time::time = $" + strconv.Itoa(count))
		parameters = append(parameters, timeQuery)
		count++
	}

	if status != "" {
		validStatuses := map[string]bool{
			"archived": true,
			"active":   true,
		}
		if !validStatuses[status] {
			helps.RespWithError(ctx, http.StatusBadRequest, "Invalid status value, must be one of: active, archived", nil)
			return
		}
		query.WriteString(" AND status = $" + strconv.Itoa(count))
		parameters = append(parameters, status)
		count++
	}

	var schedules []models.Schedule
	rows, err := database.DB.Query(query.String(), parameters...)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Error occurred while reading rows", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var schedule models.Schedule
		err = rows.Scan(&schedule.Id, &schedule.DoctorId, &schedule.StartTime, &schedule.EndTime, &schedule.Capacity, &schedule.BookedCount, &schedule.Status)
		if err != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to scan appointment details", err)
			return
		}
		schedules = append(schedules, schedule)
	}

	if err = rows.Err(); err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Error iterating over rows", err)
		return
	}

	if len(schedules) == 0 {
		helps.RespWithError(ctx, http.StatusNotFound, "No schedules found", nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": schedules,
		"status":  "success",
		"message": "Records retrieved successfully",
		"code":    http.StatusOK})
}
