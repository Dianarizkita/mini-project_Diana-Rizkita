package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	LoanDate     string `json:"loan_date" form:"loan_date"`
	DateOfReturn string `json:"date_of_return" form:"date_of_return"`
	Status       string `json:"status" form:"status"`
}
