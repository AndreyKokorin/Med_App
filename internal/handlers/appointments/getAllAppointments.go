package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func GetAllAppointment(ctx *gin.Context) {
	var apoints []models.Appointment

	query := "SELECT id, patient_id, doctor_id, appointment_time, status from Appointments "
	rows, err := database.DB.Query(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var apoint models.Appointment
		err := rows.Scan(&apoint.Id, &apoint.PatientId, &apoint.DoctorId, &apoint.AppointmentTime, &apoint.Status)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			slog.Error(err.Error())
			return
		}
		apoints = append(apoints, apoint)
	}

	if err := rows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		slog.Error("Error during iteration: " + err.Error())
		return
	}

	if len(apoints) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "No appointments found"})
		return
	}

	ctx.JSON(http.StatusOK, apoints)
}
