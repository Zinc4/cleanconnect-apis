package entities

import "gorm.io/gorm"

type AdditionalBill struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(255);not null"`
	Price int    `json:"price"`
}
