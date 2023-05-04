package models

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Gender   string `json:"gender" form:"gender"`
	Telp     string `json:"telp" form:"telp"`
	Alamat   string `json:"alamat" form:"alamat"`
	Email    string `json:"email" form:"email"`
}
