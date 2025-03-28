package directions

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	repositories "awesomeProject/internal/repositories/directions"
	"awesomeProject/pkg/helps"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFilterDirections(ctx *gin.Context) {
	filter := map[string]interface{}{
		"doctor_id":           ctx.Query("doctor_id"),
		"patient_id":          ctx.Query("patient_id"),
		"examination_type_id": ctx.Query("examination_type_id"),
		"status":              ctx.Query("status"),
	}

	directions, err := repositories.GetFilterDirections(database.DB, filter)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusOK, []models.Direction{})
			return
		}
		helps.RespWithError(ctx, http.StatusInternalServerError, "error with get directions", err)
		return
	}

	if len(directions) == 0 {
		ctx.JSON(http.StatusOK, []models.Direction{})
		return
	}

	ctx.JSON(http.StatusOK, directions)
}
