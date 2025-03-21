package appointments

import (
	"awesomeProject/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeStatusAppointment(ctx *gin.Context) {
	appointId := ctx.Param("id")
	status := ctx.Param("status")

	query := "UPDATE appointments SET status = $1 WHERE id = $2"

	res, err := database.DB.Exec(query, status, appointId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении количества удаленных строк"})
		return
	}
	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return

	}

	ctx.JSON(http.StatusOK, gin.H{"status": status, "appointId": appointId})
}
