package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/pkg/helps"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

const (
	SlotStatusAvailable        = "available"
	AppointmentStatusCancelled = "Cancel"
)

// CancelAppointment
// @Summary Отмена записи на прием
// @Description Позволяет отменить запись на прием, освобождая временной слот
// @Tags appointments
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "ID записи на прием"
// @Success 200 {object} map[string]interface{} "Запись успешно отменена"
// @Failure 400 {object} map[string]string "Некорректный идентификатор записи"
// @Failure 404 {object} map[string]string "Запись на прием или временной слот не найдены"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /appointments/{id} [delete]
func CancelAppointment(ctx *gin.Context) {
	// Получаем ID записи из параметров запроса
	appointmentID := ctx.Param("id")
	if appointmentID == "" {
		helps.RespWithError(ctx, http.StatusBadRequest, "Appointment ID is required", nil)
		return
	}

	// Начинаем транзакцию
	tx, err := database.DB.Begin()
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to start database transaction", err)
		return
	}
	defer tx.Rollback() // Откат в случае паники или незавершённой транзакции

	// Проверяем существование записи и получаем slot_id
	var slotID int
	err = tx.QueryRow(
		"SELECT slot_id FROM appointments WHERE id = $1",
		appointmentID,
	).Scan(&slotID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to rollback transaction after query error", rollbackErr)
			return
		}
		if err == sql.ErrNoRows {
			helps.RespWithError(ctx, http.StatusNotFound, "Appointment not found", errors.New("Appointment not found"))
		} else {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to check appointment", err)
		}
		return
	}

	// Обновляем статус слота на "available"
	res, err := tx.Exec(
		"UPDATE time_slots SET status = $1 WHERE id = $2",
		SlotStatusAvailable,
		slotID,
	)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to rollback transaction after slot update error", rollbackErr)
			return
		}
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to update time slot status", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to rollback transaction after rows affected error", rollbackErr)
			return
		}
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to retrieve updated rows count for time slot", err)
		slog.Error(err.Error())
		return
	}

	if rowsAffected == 0 {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to rollback transaction after no rows affected", rollbackErr)
			return
		}
		helps.RespWithError(ctx, http.StatusNotFound, "Time slot not found or already available", nil)
		return
	}

	// Обновляем статус записи на "Cancelled"
	res, err = tx.Exec(
		"UPDATE appointments SET status = $1 WHERE id = $2",
		AppointmentStatusCancelled,
		appointmentID,
	)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to rollback transaction after appointment update error", rollbackErr)
			return
		}
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to update appointment status", err)
		return
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to rollback transaction after rows affected error", rollbackErr)
			return
		}
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to retrieve updated rows count for appointment", err)
		slog.Error(err.Error())
		return
	}

	if rowsAffected == 0 {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to rollback transaction after no rows affected", rollbackErr)
			return
		}
		helps.RespWithError(ctx, http.StatusNotFound, "Appointment not found or already cancelled", nil)
		return
	}

	// Коммитим транзакцию
	if err := tx.Commit(); err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to commit transaction", err)
		return
	}

	// Успешный ответ
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Appointment cancelled successfully",
		"id":      appointmentID,
	})
}
