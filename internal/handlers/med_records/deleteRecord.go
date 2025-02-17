package med_records

import (
	"awesomeProject/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteRecord(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	query := "DELETE FROM users WHERE id=$1"
	res, err := database.DB.Exec(query, userId)
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
	ctx.JSON(http.StatusOK, gin.H{"record deleted": userId})
}
