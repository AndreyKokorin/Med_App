package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/helps"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAppointmentDetails — эндпоинт для фильтрованного запроса к appointment_details
func GetAppointmentDetails(ctx *gin.Context) {
	// Извлекаем параметры запроса
	doctorID := ctx.Query("doctor_id")
	patientID := ctx.Query("patient_id")
	appointmentStatus := ctx.Query("appointment_status")
	slotStatus := ctx.Query("slot_status")

	// Формируем базовый запрос
	query := "SELECT * FROM appointment_details WHERE 1=1"
	var args []interface{}
	paramIndex := 1

	// Добавляем фильтры, если параметры переданы
	if doctorID != "" {
		query += " AND doctor_id = $" + fmt.Sprint(paramIndex)
		args = append(args, doctorID)
		paramIndex++
	}
	if patientID != "" {
		query += " AND patient_id = $" + fmt.Sprint(paramIndex)
		args = append(args, patientID)
		paramIndex++
	}
	if appointmentStatus != "" {
		query += " AND appointment_status = $" + fmt.Sprint(paramIndex)
		args = append(args, appointmentStatus)
		paramIndex++
	}
	if slotStatus != "" {
		query += " AND slot_status = $" + fmt.Sprint(paramIndex)
		args = append(args, slotStatus)
		paramIndex++
	}

	// Выполняем запрос к базе данных
	rows, err := database.DB.Query(query, args...)
	if err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to query appointment details", err)
		return
	}
	defer rows.Close()

	// Собираем результаты
	var details []models.AppointmentDetail
	for rows.Next() {
		var d models.AppointmentDetail
		err := rows.Scan(
			&d.AppointmentID,
			&d.PatientID,
			&d.AppointmentStatus,
			&d.SlotID,
			&d.SlotStartTime,
			&d.SlotStatus,
			&d.ScheduleID,
			&d.DoctorID,
			&d.ScheduleStartTime,
			&d.ScheduleEndTime,
			&d.Capacity,
			&d.BookedSlots,
			&d.ScheduleStatus,
		)
		if err != nil {
			helps.RespWithError(ctx, http.StatusInternalServerError, "Failed to scan appointment details", err)
			return
		}
		details = append(details, d)
	}

	// Проверяем ошибки после итерации
	if err := rows.Err(); err != nil {
		helps.RespWithError(ctx, http.StatusInternalServerError, "Error occurred while reading rows", err)
		return
	}

	// Если ничего не найдено
	if len(details) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No appointments found matching the criteria",
			"data":    []models.AppointmentDetail{},
		})
		return
	}

	// Успешный ответ
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Appointments retrieved successfully",
		"data":    details,
	})
}
