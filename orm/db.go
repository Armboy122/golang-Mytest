package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	dsn := "host=localhost user=peagolang password=supersecret dbname=peagolang port=5432 sslmode=disable"
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("filed to connect database")
	}

	Db.AutoMigrate(&User{})
}
