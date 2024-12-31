package entities

import (
	"time"

	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	CustomerID       uint           `json:"customer_id"`
	Customer         Customer       `json:"customer" gorm:"foreignKey:CustomerID"`
	AdditionalBillID uint           `json:"additional_bill_id"`
	AdditionalBill   AdditionalBill `json:"additional_bill" gorm:"foreignKey:AdditionalBillID"`
	Description      string         `json:"description"`
	Amount           int            `json:"amount"`
	BillDate         time.Time      `json:"bill_date"`
	BillDue          time.Time      `json:"bill_due"`
	Status           string         `json:"status" default:"pending"`
	Category         string         `json:"category"`
	QrUrl            string         `json:"url"`
}
