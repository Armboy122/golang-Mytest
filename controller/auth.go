package auth

import (
	"net/http"

	"github.com/Armboy122/golang-Mytest/orm"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// สร้าง register
type RegisterBody struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Check user
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "fail",
		})
		return
	}

	//Create User
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := orm.User{Username: json.Username,
		Password: string(encryptedPassword),
		Fullname: json.Fullname}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "success",
			"userID":  user.ID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "fail",
		})
	}
}

// สรา้ง login

type LoginBody struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Check user
	var userExist orm.User
	orm.Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Dose not exists",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(json.Password))
	if err == nil {
		//Check user
		var userExist orm.User
		orm.Db.Where("username = ?", json.Username).First(&userExist)
		if userExist.ID > 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "success",
			})
			return
		}
	} else {
		//Check user
		var userExist orm.User
		orm.Db.Where("username = ?", json.Username).First(&userExist)
		if userExist.ID > 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "error",
				"message": "fail",
			})
			return
		}
	}
}
