package config

import (
	"fmt"

	models "api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnetcion() {
	dbConfig, err := LoadEnv(".")
	if err != nil {
		fmt.Println(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.IP, dbConfig.Port, dbConfig.DB_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserAccounts{})
	db.AutoMigrate(&models.UserDetails{})
	DB = db
}
