package user

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Armboy122/golang-Mytest/orm"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ReadAll(c *gin.Context) {
	//------------------เช็ค token -----------------------
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
	header := c.Request.Header.Get("Authorization")          //ค่าของ token ใส่ในตัวแปล header
	tokenString := strings.Replace(header, "Bearer ", "", 1) //แยกค่า token ออกจาก bearer
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//------------------ข้อมูลที่จะถูกเรียกหลังlogin-----------------------
		// fmt.Println(claims["userId"])
		var users []orm.User
		orm.Db.Find(&users)
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "User Read success",
			"users":   users,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

}
