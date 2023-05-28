package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/Armboy122/golang-Mytest/controller/auth" //เรียกใช้func
	"github.com/Armboy122/golang-Mytest/controller/user" //เรียกใช้func
	"github.com/Armboy122/golang-Mytest/orm"             //เรียกใช้  func init จากไฟล์ orm
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	orm.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.GET("/user/readall", user.ReadAll)

	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
