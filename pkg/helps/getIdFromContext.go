package helps

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetIdFromContext(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get("user_id")
	if !ok {
		return 0, errors.New("error getting user_id from context")
	}

	// Проверка типа userId
	userIdInt, ok := userId.(int)
	if !ok {
		return 0, errors.New("user_id is not int")
	}

	return userIdInt, nil
}
