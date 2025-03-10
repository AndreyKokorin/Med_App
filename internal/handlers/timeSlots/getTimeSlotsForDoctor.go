package timeSlots

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/helps"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetActualTimeSlotsForDoctor
// @Summary Получение доступных слотов доктора
// @Description Возвращает список актуальных временных слотов для указанного доктора
// @Tags doctors
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "ID доктора"
// @Success 200 {object} map[string][]models.TimeSlot "Список доступных слотов"
// @Failure 400 {object} map[string]string "Ошибка валидации"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/doctors/{id}/slots [get]
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
	if len(slots) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"massage": "Not found slots", "data": []models.TimeSlot{}})
		return
	}

	if err = rows.Err(); err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Error occurred while processing time slots", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": slots})
}
