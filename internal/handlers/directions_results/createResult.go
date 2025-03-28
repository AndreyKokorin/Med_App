package directions_results

import (
	"awesomeProject/internal/database"
	repositories "awesomeProject/internal/repositories/directions"
	"awesomeProject/pkg/helps"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func CreateResultHandler(ctx *gin.Context) {
	doctorId, err := helps.GetIdFromContext(ctx)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "failed to get doctor id from ctx", err)
		return
	}

	formFile, err := ctx.FormFile("result_file")
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "failed to get result file", err)
		return
	}

	slog.Info("Request Content-Type: " + ctx.ContentType())
	form, _ := ctx.MultipartForm()
	if form != nil {
		slog.Info("Form files: ", form.File)
	}

	directionId := ctx.PostForm("id")
	if directionId == "" {
		helps.RespWithError(ctx, http.StatusBadRequest, "direction_id is required", errors.New("direction_id is required"))
		return
	}

	directIdInt, err := strconv.Atoi(directionId)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "invalid direction_id format", err)
		return
	}

	fileUrl, err := LoadFile(*formFile)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "failed to load file", err)
		return
	}

	result, err := repositories.CreateResultExamination(directIdInt, doctorId, fileUrl, database.DB)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "failed to create result examination", err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
