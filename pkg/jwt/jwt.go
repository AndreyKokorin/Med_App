package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const AccessTokenTTL = 5 * time.Hour
const RefreshTokenTTL = 24 * time.Hour

func NewJWT(userID int, role string, exp time.Time) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp.Unix(),
		"role":    role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func ParseJWT(tokenString string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм совпадает
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем срок действия токена
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("exp claim is missing")
	}

	// Проверяем, не истек ли токен
	if int64(exp) < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	// Декодируем user_id и role
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid user_id claim")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid role claim")
	}

	return map[string]interface{}{
		"user_id": int(userID),
		"role":    role,
	}, nil
}
