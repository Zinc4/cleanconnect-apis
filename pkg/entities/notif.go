package entities

import "gorm.io/gorm"

type Notif struct {
	gorm.Model
	Notification string `json:"notification" gorm:"type:text;not null"`
	UserID       uint   `json:"user_id"`
	Username     string `json:"username" gorm:"type:varchar(255);not null"`
	Amount       int    `json:"amount"`
}
