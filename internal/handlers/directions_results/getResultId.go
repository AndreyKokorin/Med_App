package directions_results

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

func GetResultByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		helps.RespWithError(
			ctx,
			http.StatusBadRequest,
			"id is required",
			errors.New("id is required"),
		)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helps.RespWithError(
			ctx,
			http.StatusBadRequest,
			"id is invalid",
			errors.New("id is invalid"),
		)
		return
	}

	result, err := repositories.GetResultById(idInt, database.DB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helps.RespWithError(
				ctx,
				http.StatusNotFound,
				"",
				errors.New("results not found"),
			)
			return
		}

		helps.RespWithError(
			ctx,
			http.StatusInternalServerError,
			"",
			errors.New("id is invalid"),
		)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
