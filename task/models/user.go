package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Gender       string `json:"gender" form:"gender"`
	Telp         string `json:"telp" form:"telp"`
	Alamat       string `json:"alamat" form:"alamat"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Transactions []Transaction
}
