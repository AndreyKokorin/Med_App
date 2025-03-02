package timeSlots

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/helps"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetActualTimeSlotsForDoctor(ctx *gin.Context) {
	doctorId := ctx.Param("id")

	if doctorId == "" {
		helps.RespWithError(ctx, http.StatusBadRequest, "Doctor ID is required", nil)
		return
	}

	query := `SELECT
        s.doctor_id,
        ts.id AS slot_id,
        ts.start_time AS slot_start_time,
        ts.end_time AS slot_end_time,
        ts.status AS slot_status
    FROM schedules s
    INNER JOIN time_slots ts ON s.id = ts.schedule_id
    WHERE s.doctor_id = $1
    AND s.status = 'active'
    ORDER BY ts.start_time`

	rows, err := database.DB.Query(query, doctorId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helps.RespWithError(ctx, http.StatusNotFound, "No available time slots found for the doctor", nil)
			return
		}
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to retrieve time slots from database", err)
		return
	}
	defer rows.Close()

	var slots []models.TimeSlot
	for rows.Next() {
		var slot models.TimeSlot
		err = rows.Scan(&slot.DoctorId, &slot.Id, &slot.StartTime, &slot.EndTime, &slot.Status)
		if err != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to scan time slot data", err)
			return
		}
		slots = append(slots, slot)
	}

	// Check for any errors that occurred during row iteration
	if err = rows.Err(); err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Error occurred while processing time slots", err)
		return
	}

	ctx.JSON(http.StatusOK, slots)
}
