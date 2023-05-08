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
		DB_Username string
		DB_Password string
		DB_Port     string
		DB_Host     string
		DB_Name     string
	}{
		DB_Username: "root",
		DB_Password: "",
		DB_Host:     "localhost",
		DB_Port:     "3306",
		DB_Name:     "mini_project",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Book{}, &models.Book_Details{}, &models.User{}, &models.Book_Return{}, &models.Transaction{})
}
