package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string `json:"title" form:"title"`
	Author          string `json:"author" form:"author"`
	Publisher       string `json:"publisher" form:"publisher"`
	PublicationYear string `json:"publication_year" form:"publication_year"`
	Genre           string `json:"genre" form:"genre"`
	BookStock       int32  `json:"book_stock" form:"book_stock"`
	Book_Detailss   []Book_Details
}
