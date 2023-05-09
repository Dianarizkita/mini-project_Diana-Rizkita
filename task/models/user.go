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
	Password     string `json:"password" form:"password"`
	Email        string `json:"email" form:"email"`
	Token        string `gorm:"-"`
	Transactions []Transaction
}
