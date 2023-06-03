package works

import (
	"net/http"

	"github.com/Armboy122/golang-Mytest/orm"
	"github.com/gin-gonic/gin"
)

func Creatework(c *gin.Context) {
	var json orm.Works // json เก็บข้อมูลของสิ่งที่ส่งเข้าไป
	c.Bind(&json)

	result := orm.Db.Create(&json) //กำหนดค่าไว้เช็ค err

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": json,
	})

}

func Readworks(c *gin.Context) {
	var json []orm.Works // json เก็บข้อมูลของDB เป็นแบบ array
	orm.Db.Find(&json)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Works Read success",
		"users":   json,
	})
}

func Listworks(c *gin.Context) {
	id := c.Param("id") // รับค่า id ที่ส่งมา
	var json orm.Works  // json เก็บข้อมูลของสิ่งที่ส่งเข้าไป
	orm.Db.First(&json, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Works Read success",
		"users":   json,
	})

}
func Deletworks(c *gin.Context) {
	id := c.Param("id") // รับค่า id ที่ส่งมา
	var json orm.Works  // json เก็บข้อมูลของสิ่งที่ส่งเข้าไป
	orm.Db.Delete(&json, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Delete success",
	})
}
func Updatework(c *gin.Context) {
	id := c.Param("id")                    // รับค่า id ที่ส่งมา
	var json orm.Works                     // ตัวแปรข้อมูลใน DB
	var work orm.Works                     // ตัวแปรข้อมูลใหม่ที่จะใส่ค่าเข้าไป
	c.Bind(&work)                          //เช็คค่าความต่างของ ข้อมูลใหม่เก่า
	orm.Db.First(&json, id)                //หาข้อมูล id จาก DB
	orm.Db.Model(&json).Updates(orm.Works{ // update ข้อมูลใส่เข้าไป
		Date:   work.Date,
		Detail: work.Detail,
	})
}
