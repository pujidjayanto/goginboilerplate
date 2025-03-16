package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pujidjayanto/goginboilerplate/internal/config"
	"github.com/pujidjayanto/goginboilerplate/pkg/delivery"
)

func Authenticate() gin.HandlerFunc {
	jwtSecret := []byte(config.GetSecretKey())

	return func(ginCtx *gin.Context) {
		authHeader := ginCtx.GetHeader("Authorization")
		if authHeader == "" {
			delivery.Failed(ginCtx, http.StatusUnauthorized, "Authorization header is required")
			ginCtx.Abort()
			return
		}

		// Split the header to get the token part
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			delivery.Failed(ginCtx, http.StatusUnauthorized, "Authorization header format must be Bearer {token}")
			ginCtx.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			delivery.Failed(ginCtx, http.StatusUnauthorized, "invalid token")
			ginCtx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userId, ok := claims["userId"].(float64); ok {
				ginCtx.Set("userId", uint(userId))
			}
		}

		ginCtx.Next()
	}
}
