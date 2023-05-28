package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Armboy122/golang-Mytest/orm"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}
type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var hmacSampleSecret []byte

func Register(c *gin.Context) {
	var json RegisterBody // json เก็บข้อมูลของสิ่งที่ส่งเข้าไป
	//------------------ทดสอบว่าต่อ db ได้มั้ย ถ้าไม่ได้ return ออกเลย-----------------------
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//------------------เช็คก่อนว่าเคยregisterมารึยัง-----------------------
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Exists"})
		return
	}
	//------------------สร้างตัวเก็บข้อมูลลงDB-----------------------
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10) //keyเปลี่ยน password เป็นรหัสยากๆ
	user := orm.User{Username: json.Username,
		Password: string(encryptedPassword),
		Fullname: json.Fullname} // data ที่ได้มาให้เอามาเก็บใน user
	orm.Db.Create(&user) // เพิ่มข้อมูลเข้าไปใน table
	//------------------เช็คค่าที่ส่งกลับมา-----------------------
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Create Success", "userId": user.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Create Failed"})
	}
}

func Login(c *gin.Context) {
	var json LoginBody // json เก็บข้อมูลของสิ่งที่ส่งเข้าไป
	//------------------ทดสอบว่าต่อ db ได้มั้ย ถ้าไม่ได้ return ออกเลย-----------------------
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//------------------เช็คUser ว่ามี User รึป่าว-----------------------
	var userExist orm.User // Username ที่อยู่ใน DB
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Does Not Exists"})
		return
	}
	//------------------เช็คPassword ของ User ว่าถูกต้องรึป่าว-----------------------
	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": userExist.ID,
			"exp":    time.Now().Add(time.Minute * 1).Unix(),
		})

		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)

		c.JSON(http.StatusOK, gin.H{
			"status":       "ok",
			"message":      "login success",
			"token":        tokenString,
			"statusclaims": token.Claims.(jwt.MapClaims),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "login fail",
		})
	}
}
