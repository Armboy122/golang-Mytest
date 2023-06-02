package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/Armboy122/golang-Mytest/controller/auth" //เรียกใช้func
	"github.com/Armboy122/golang-Mytest/controller/user" //เรียกใช้func
	"github.com/Armboy122/golang-Mytest/middleware"
	"github.com/Armboy122/golang-Mytest/orm" //เรียกใช้  func init จากไฟล์ orm
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//---------------------------เช็คไฟล์ .env-----------------------------
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	//---------------------------ต่อ db-----------------------------
	orm.InitDB()
	//---------------------------path ต่างๆของDB-----------------------------
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", auth.Register) // สมัครใช้งาน
	r.POST("/login", auth.Login)       // login เข้าระบบ
	authorized := r.Group("/user", middleware.Logger())
	authorized.GET("/readall", user.ReadAll) // สำหรับแอดมินไว้ดูงานทุกคน
	authorized.GET("/profile", user.Profile) // ให้ User ดูงานและแก้ไขห้ามลบ

	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
