package config

import (
	"fmt"
	"task/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitMigrate()
}

func InitDB() {

	config := struct {
		DB_USERNAME string
		DB_PASSWORD string
		DB_PORT     string
		DB_HOST     string
		DB_NAME     string
	}{
		DB_USERNAME: "root",
		DB_PASSWORD: "",
		DB_HOST:     "localhost",
		DB_PORT:     "3306",
		DB_NAME:     "mini_project",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Book{}, &models.Book_Details{}, &models.User{}, models.Transaction{}, &models.Book_Return{})
}
