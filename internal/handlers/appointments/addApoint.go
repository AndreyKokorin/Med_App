package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"awesomeProject/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAppointment(ctx *gin.Context) {
	var apoint models.Appointment
	if err := ctx.ShouldBindJSON(&apoint); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate.ValidAndTrim(&apoint)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var apointId int
	query := "INSERT INTO Appointments(patient_id,doctor_id, appointment_time, status) VALUES ($1,$2,$3,$4) RETURNING id"
	err = database.DB.QueryRow(query, apoint.PatientId, apoint.DoctorId, apoint.AppointmentTime, apoint.Status).Scan(&apointId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"appointment_id": apointId})
}
