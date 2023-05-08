package models

import (
	"time"

	"gorm.io/gorm"
)

type Book_Return struct {
	gorm.Model
	TransactionID uint      `json:"transaction_id" form:"transaction_id"`
	DateOfReturn  time.Time `json:"date_of_return"`
	Penalty       string    `json:"penalty" form:"penalty"`
}
