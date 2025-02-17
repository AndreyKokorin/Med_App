package appointments

import (
	"awesomeProject/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteAppointment(ctx *gin.Context) {
	idAppoint, err := strconv.Atoi(ctx.Param("id"))

	query := "DELETE FROM appointments WHERE id_appointment = $1"
	res, err := database.DB.Exec(query, idAppoint)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении количества удаленных строк"})
		return
	}
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"userId": idAppoint})
}
