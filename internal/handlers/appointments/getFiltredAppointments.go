package appointments

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

func GetFilterAppointments(ctx *gin.Context) {
	doctorId := ctx.Query("doctorId")
	patientId := ctx.Query("patientId")
	date := ctx.Query("date")

	query := strings.Builder{}
	query.WriteString("SELECT id, patient_id, doctor_id, appointment_time, status FROM appointments WHERE 1=1")
	args := []interface{}{}
	argId := 1

	if doctorId != "" {
		query.WriteString(fmt.Sprintf(" AND doctor_id = $%d", argId))
		args = append(args, doctorId)
		argId++
		slog.Info(doctorId)
	}

	if patientId != "" {
		query.WriteString(fmt.Sprintf(" AND patient_id = $%d", argId))
		id, _ := strconv.Atoi(patientId)
		args = append(args, id)
		argId++
		slog.Info(patientId)
	}

	if date != "" {
		slog.Info(date)
		query.WriteString(fmt.Sprintf(" AND DATE(appointment_time) = $%d", argId))
		args = append(args, date)
		argId++
	}

	slog.Info(query.String())

	rows, err := database.DB.Query(query.String(), args...)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	var appoints []models.Appointment
	for rows.Next() {
		var appoint models.Appointment

		err = rows.Scan(&appoint.Id, &appoint.PatientId, &appoint.DoctorId, &appoint.AppointmentTime, &appoint.Status)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}

		appoints = append(appoints, appoint)
	}

	if len(appoints) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"err": "No appointments found"})
		return
	}

	ctx.JSON(http.StatusOK, appoints)
}
