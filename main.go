package main

import (
	AuthController "github.com/Armboy122/golang-Mytest/controller"
	"github.com/Armboy122/golang-Mytest/orm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Register struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
}

func main() {
	orm.InitDB()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
