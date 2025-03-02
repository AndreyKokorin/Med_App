package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/helps"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log/slog"
	"net/http"
	"time"
)

func AddAppointment(ctx *gin.Context) {
	var apoint models.Appointment
	if err := ctx.ShouldBindJSON(&apoint); err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "Invalid request body format", err)
		return
	}

	// Начинаем транзакцию
	tx, err := database.DB.Begin()
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to start database transaction", err)
		return
	}

	var slotStatus string
	var slotStartTime time.Time
	err = tx.QueryRow(
		"SELECT status, start_time FROM time_slots WHERE id = $1",
		apoint.Slot_id,
	).Scan(&slotStatus, &slotStartTime)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			helps.RespWithError(ctx, http.StatusNotFound, "Time slot not found", errors.New("Time slot not found"))
		} else {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to check time slot", err)
		}
		return
	}

	// Проверка статуса и времени слота
	loc, _ := time.LoadLocation("Asia/Almaty") // GMT+5
	currentTime := time.Now().In(loc)

	if slotStatus == "booked" {
		tx.Rollback()
		helps.RespWithError(ctx, http.StatusConflict, "Time slot is already booked", errors.New("Time slot is already booked"))
		return
	}
	if slotStartTime.Before(currentTime) {
		tx.Rollback()
		helps.RespWithError(ctx, http.StatusGone, "Time slot has already passed", errors.New("Time slot has already passed"))
		return
	}

	// Обновляем статус слота
	const SlotStatusBooked = "booked"
	res, err := tx.Exec("UPDATE time_slots SET status = $1 WHERE id = $2 AND status != 'booked' AND start_time > NOW() AT TIME ZONE 'Asia/Almaty'", SlotStatusBooked, apoint.Slot_id)
	if err != nil {
		tx.Rollback()
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to update time slot status", err)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to retrieve updated rows count", err)
		slog.Error(err.Error())
		return
	}

	if rowsAffected == 0 {
		tx.Rollback()
		helps.RespWithError(ctx, http.StatusConflict, "Time slot is already booked or does not exist", nil)
		return
	}
	var appointmentID int
	// Вставляем запись в appointments
	err = tx.QueryRow("INSERT INTO appointments(patient_id, status, slot_id, schedule_id) VALUES ($1, $2 ,$3, $4) RETURNING id",
		apoint.PatientId, "Pending", apoint.Slot_id, apoint.Schedule_id).Scan(&appointmentID)
	if err != nil {
		tx.Rollback()
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code == "23505" {
			helps.RespWithError(ctx, http.StatusConflict, "Appointment already exists", err)
		} else {
			helps.RespWithError(ctx, http.StatusBadRequest, "Failed to create appointment record", err)
		}
		return
	}

	res, err = tx.Exec("UPDATE schedules SET booked_slots = booked_slots + 1 WHERE id = $1", apoint.Schedule_id)
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to retrieve updated rows count", err)
		return
	}

	if rowsAffected == 0 {
		tx.Rollback()
		helps.RespWithError(ctx, http.StatusConflict, "Schedule does not exist", nil)
		return
	}

	// Коммитим транзакцию
	err = tx.Commit()
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to commit transaction", err)
		return
	}

	ctx.JSON(http.StatusCreated, apoint)
}

/* docker exec -it f7dd40fb36c6 psql -U postgres -d fitness_api
 */
