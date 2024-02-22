package database

import (
	"cmd/blog-website-backend/main.go/models"
	"cmd/blog-website-backend/main.go/pkg/config"
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"), "?parseTime=true")
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	if err = DB.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		panic(err)
	}
}
