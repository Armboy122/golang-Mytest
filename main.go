package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID          string `json:"id"` // json: แปลงชื่อให้อยู่ในรูปแบบที่ต้องการจากตัวใหญ่ก็เป็นตัวเล็กจากคำยาวเป็นสั้น
	Name        string `json:"name"`
	Description string `json:"desc"`
}

// สร้างข้อมูลมาลองDB
var courses = []Course{ //Couse บอกว่า ตัวแปลนี้มีชนิดข้อมูลทั้สรา้งตาม type ที่กล่าวถึง
	{ID: "1", Name: "TDD", Description: "let Go"},
	{ID: "2", Name: "Arm", Description: "let armboy Go"},
}

func main() {
	fmt.Println("hello")
	r := gin.Default()

	//สร้าง path ที่ใช้งาน
	r.GET("/courses", listCourses)
	r.GET("/courses/:id", getCourses)

	r.Run(":8000")
}

// func เรียกดูข้อมูลที่มี
func listCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}

func getCourses(c *gin.Context) {
	id := c.Param("id")
	for _, course := range courses {
		if course.ID == id {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "course not found",
	})
}
