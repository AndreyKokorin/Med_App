package middleware

import (
	"awesomeProject/pkg/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strings"
)

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.GetHeader("Authorization")

		splitedToken := strings.Split(authToken, " ")
		if len(splitedToken) != 2 {
			slog.Error("нету токена")
			unauthorized(ctx)
			return
		}
		if splitedToken[0] != "Bearer" {
			slog.Error("нету bearer token")
			unauthorized(ctx)
			return
		}

		dataFromToken, err := jwt.ParseJWT(splitedToken[1])
		if err != nil {
			slog.Error("не верный токен")
			unauthorized(ctx)
			return
		}

		userRole, ok := dataFromToken["role"].(string)
		if !ok || !isRoleAllowed(userRole, allowedRoles) {
			slog.Error("нету доступа роли")
			log.Println(userRole)
			log.Println(allowedRoles)
			unauthorized(ctx)
			return
		}

		userID, ok := dataFromToken["user_id"].(int) // или "user_id" в зависимости от структуры токена
		if !ok {
			slog.Error("User ID not found in token or invalid format")
			unauthorized(ctx)
			return
		}

		ctx.Set("user_id", userID)
		ctx.Set("role", userRole)

		ctx.Next()
	}
}

func isRoleAllowed(role string, allowedRoles []string) bool {
	for _, r := range allowedRoles {
		if r == role {
			return true
		}
	}
	return false
}
func unauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	ctx.Abort()
}
