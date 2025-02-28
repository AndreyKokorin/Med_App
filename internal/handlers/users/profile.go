package users

import (
	"awesomeProject/internal/database"
	repositories "awesomeProject/internal/repositories/user"
	"awesomeProject/pkg/helps"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func GetProfile(ctx *gin.Context) {
	id, ok := ctx.Get("user_id")

	if !ok {
		helps.RespWithError(ctx, http.StatusUnauthorized, "user_id not found", errors.New("user_id not found"))
		return
	}

	slog.Info("user_id: ", id)

	user, err := repositories.GetUserId(database.DB, id.(int))

	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
		}
		helps.RespWithError(ctx, status, "failed to fetch user", err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
