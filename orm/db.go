package orm

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
}

var Db *gorm.DB
var err error

func InitDB() {
	dsn := os.Getenv("MYSQL_DNS")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("filed to connect database")
	}
	Db.AutoMigrate(&User{})
}
