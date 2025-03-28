package directions

import (
	"awesomeProject/internal/database"
	repositories "awesomeProject/internal/repositories/directions"
	"awesomeProject/pkg/helps"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDirectionId(ctx *gin.Context) {
	directionId := ctx.Param("id")
	if directionId == "" {
		helps.RespWithError(ctx, http.StatusBadRequest, "id parameter is required", nil)
		return
	}

	directionIdInt, err := strconv.Atoi(directionId)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error with convert id", err)
		return
	}
	if directionIdInt < 0 {
		helps.RespWithError(ctx, http.StatusBadRequest, "wrong id in request", errors.New("wrong id in request"))
		return
	}

	direction, err := repositories.GetDirectionId(directionIdInt, database.DB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helps.RespWithError(ctx, http.StatusNotFound, "direction with this id is not found", errors.New("direction with this id is not found"))
			return
		}
		helps.RespWithError(ctx, http.StatusInternalServerError, "error with get direction", err)
		return
	}

	ctx.JSON(http.StatusOK, direction)
}
