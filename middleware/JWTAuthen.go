package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		hms := []byte(os.Getenv("JWT_SECRET_KEY"))
		header := c.Request.Header.Get("Authorization")          //ค่าของ token ใส่ในตัวแปล header
		tokenString := strings.Replace(header, "Bearer ", "", 1) //แยกค่า token ออกจาก bearer
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hms, nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
		} else {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
		}
		c.Next()
	}
}
