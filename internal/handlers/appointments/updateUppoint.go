package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func UpdateAppointDate(ctx *gin.Context) {
	// Получаем и проверяем ID
	apointId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		slog.Error("Invalid user ID", slog.Any("error", err))
		return
	}

	// Читаем JSON
	var updateData models.AppointmentsUpdateTime
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		slog.Error("User request failed", slog.Any("error", err))
		return
	}

	// Проверяем, передано ли новое время
	if updateData.AppointmentTime.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "appointment_time is required"})
		return
	}

	// Обновляем запись
	var updateTime time.Time
	query := `UPDATE Appointments 
              SET appointment_time = $1  
              WHERE id = $2 
              RETURNING appointment_time`
	err = database.DB.QueryRow(query, updateData.AppointmentTime, apointId).Scan(&updateTime)

	// Проверяем ошибки
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		slog.Error("User SQL failed", slog.Any("error", err))
		return
	}

	// Отправляем успешный ответ
	ctx.JSON(http.StatusOK, gin.H{"newApointTime": updateTime})
}
