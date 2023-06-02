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
	Role     string
}

var Db *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DB")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}
	Db.AutoMigrate(&User{})
}
