package middlewares

import (
	// "net/http"

	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("JWT_SECRET")
		if secretKey == "" {
			c.AbortWithStatusJSON(500, gin.H{"error": "Secret key Don't exist"})
		}
		authHeader := c.GetHeader("Authorization")
		splittedHeader := strings.Split(authHeader, "Bearer ")

		if len(splittedHeader) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Malformed Token"})
		}
		jwtToken := splittedHeader[1]

		token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected SignIn method: %v", t.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid Token"})
			return
		}

		if claim, ok := token.Claims.(jwt.MapClaims); ok {
			if UserID, exist := claim["user_id"]; exist {
				c.Set("user_id", UserID)
			}
		}
		c.Next()
	}
}
