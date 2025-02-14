package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func GetAppointment(ctx *gin.Context) {
	ApointId := ctx.Param("id")

	var apoint models.Appointment
	query := "SELECT id, patient_id, doctor_id, appointment_time, status from Appointments WHERE id = $1"
	err := database.DB.QueryRow(query, ApointId).Scan(&apoint.Id, &apoint.PatientId, &apoint.DoctorId, &apoint.AppointmentTime, &apoint.Status)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erroe": err.Error()})
		slog.Error(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, apoint)
}
