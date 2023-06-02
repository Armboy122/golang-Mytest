package works

import (
	"net/http"

	"github.com/Armboy122/golang-Mytest/orm"
	"github.com/gin-gonic/gin"
)

func Creatework(c *gin.Context) {
	var json orm.Works // json เก็บข้อมูลของสิ่งที่ส่งเข้าไป

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	orm.Db.Create(&json)
	c.JSON(http.StatusOK, gin.H{"data": json})
}
