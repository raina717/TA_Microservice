package models

//connect database disini
import (
	"Backend/core"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(mysql.Open(core.BaseDatabase), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = database
}
