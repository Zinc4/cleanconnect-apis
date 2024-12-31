package entities

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Status     string   `json:"status" gorm:"type:varchar(255);not null"`
	Image      string   `json:"image" gorm:"type:varchar(255);not null"`
	BillID     uint     `json:"bill_id"`
	Bill       Bill     `json:"bill" gorm:"foreignKey:BillID"`
	CustomerID uint     `json:"customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignKey:CustomerID"`
}
