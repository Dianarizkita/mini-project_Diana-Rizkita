package models

import (
	"gorm.io/gorm"
)

type Book_Details struct {
	gorm.Model
	BookCode string `json:"book_code" form:"book_code"`
	Book     []Book `json:"book" form:"book"`
	Status   string `json:"status" form:"status"`
}
