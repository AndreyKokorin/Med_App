package directions

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/models"
	repositories "awesomeProject/internal/repositories/directions"
	"awesomeProject/pkg/helps"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func CreateDirectionsHandler(ctx *gin.Context) {
	if !isDoctor(ctx) {
		return
	}

	currentDoctorId, ok := ctx.Get("user_id")
	if !ok {
		helps.RespWithError(ctx, http.StatusUnauthorized, "Your id was not received", errors.New("Your id was not received"))
		return
	}

	intDoctorId, ok := currentDoctorId.(int)
	if !ok {
		helps.RespWithError(ctx, http.StatusUnauthorized, "Your id was not received", errors.New("Your id was not received"))
		return
	}

	var direction models.Direction
	if err := ctx.ShouldBindJSON(&direction); err != nil {
		helps.RespWithError(
			ctx,
			http.StatusBadRequest,
			"Invalid request format",
			fmt.Errorf("failed to parse request body: %w", err),
		)
		slog.Error(fmt.Sprintf("failed to parse request body: %w", err))
		return
	}

	if intDoctorId == direction.DoctorID {
		helps.RespWithError(
			ctx,
			http.StatusConflict,
			"Cannot create referral to yourself",
			fmt.Errorf("doctor with ID %d attempted to create self-referral", intDoctorId),
		)
		return
	}

	direction, err := repositories.CreateNewDirection(database.DB, direction, intDoctorId)

	if err != nil {
		helps.RespWithError(
			ctx,
			http.StatusInternalServerError,
			"Invalid create direction",
			fmt.Errorf("invalid create direction: %w", err),
		)
		slog.Error(fmt.Sprintf("invalid create direction: %w", err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":    http.StatusCreated,
		"direction": direction,
	})
}

func isDoctor(ctx *gin.Context) bool {
	role, ok := ctx.Get("role")
	if !ok {
		helps.RespWithError(ctx, http.StatusUnauthorized, "Your role was not received", errors.New("Your role was not received"))
		return false
	}

	if role != "doctor" {
		return false
		helps.RespWithError(
			ctx,
			http.StatusForbidden,
			"Access denied: insufficient permissions",
			fmt.Errorf("user role '%s' has no access to this resource", role),
		)
	}

	return true
}
