package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/validate"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func AddAppointment(ctx *gin.Context) {
	var apoint models.Appointment
	if err := ctx.ShouldBindJSON(&apoint); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apoint.Status = "Pending"

	err := validate.ValidAndTrim(&apoint)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slog.Info("apoint: ", apoint)

	checkQuery := `
SELECT id, booked_slots, capacity
FROM schedules
WHERE doctor_id = $1
  AND start_time <= $2
  AND end_time > $2
  AND status = 'active'
  AND booked_slots < capacity;
`

	row := database.DB.QueryRow(checkQuery, apoint.DoctorId, apoint.AppointmentTime)

	var id, bookedSlots, capacity int
	err = row.Scan(&id, &bookedSlots, &capacity)

	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no available schedule for this appointment"})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	var apointId int
	query := "INSERT INTO Appointments(patient_id,doctor_id, appointment_time, status,schedule_id) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	err = database.DB.QueryRow(query, apoint.PatientId, apoint.DoctorId, apoint.AppointmentTime, apoint.Status, apoint.ShedulesID).Scan(&apointId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"appointment_id": apointId})
}
