package models

import (
	"gorm.io/gorm"
)

type Book_Return struct {
	gorm.Model
	DateOfReturn string `json:"date_of_return" form:"date_of_return"`
	Penalty      string `json:"penalty" form:"penalty"`
}
