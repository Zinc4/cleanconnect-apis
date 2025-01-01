package entities

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	ChangeType string `json:"change_type" gorm:"type:varchar(255);not null"`
	OldValue   string `json:"old_value" gorm:"type:text;not null"`
	NewValue   string `json:"new_value" gorm:"type:text;not null"`
	UserID     uint   `json:"user_id"`
	Status     string `json:"status" gorm:"type:varchar(50);not null"`
}
