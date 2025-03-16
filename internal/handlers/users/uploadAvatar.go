package users

import (
	"awesomeProject/internal/database"
	"awesomeProject/pkg/byteScale"
	"awesomeProject/pkg/helps"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"
	"time"
)

// UploadUserAvatar uploads a user's avatar.
// @Summary Загрузка аватара пользователя
// @Description Загружает новый аватар для текущего пользователя
// @Tags users
// @Security ApiKeyAuth
// @Accept multipart/form-data
// @Produce json
// @Param avatar formData file true "Файл изображения (jpg, jpeg, png)"
// @Success 200 {object} map[string]string "Аватар успешно загружен"
// @Failure 400 {object} map[string]string "Ошибка загрузки файла"
// @Failure 401 {object} map[string]string "Пользователь не авторизован"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /shared/users/avatar [put]
func UploadUserAvatar(ctx *gin.Context) {
	userId, err := helps.GetIdFromContext(ctx)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error with getting user_id from contexts", err)
		return
	}

	file, err := ctx.FormFile("avatar")
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error uploading file from form", err)
		return
	}

	ext := filepath.Ext(file.Filename)
	var contentType string
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	default:
		helps.RespWithError(ctx, http.StatusBadRequest, "invalid file type", nil)
		return
	}

	newFileName := fmt.Sprintf("%d_%d%s", userId, time.Now().UnixNano(), ext)
	filePath := newFileName

	fileOpened, err := file.Open()
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error with opening file", err)
		return
	}
	defer fileOpened.Close()

	fileByte, err := io.ReadAll(fileOpened)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error with reading file", err)
		return
	}

	avatarUrl, err := byteScale.UploadFile(fileByte, contentType, filePath)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error uploading file to api", err)
		return
	}

	query := "UPDATE users SET avatar_url = $1 WHERE id = $2"
	res, err := database.DB.Exec(query, avatarUrl, userId)
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error with adding url to db", err)
		return
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil {
		helps.RespWithError(ctx, http.StatusBadRequest, "error with RowsAffected", err)
		return
	}

	if RowsAffected == 0 {
		helps.RespWithError(ctx, http.StatusBadRequest, "error not one line has been changed", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "avatar uploaded successfully", "avatar_url": avatarUrl})
}
