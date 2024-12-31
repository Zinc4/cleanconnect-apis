package entities

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName         string `json:"first_name" gorm:"type:varchar(255);not null"`
	LastName          string `json:"last_name" gorm:"type:varchar(255);not null"`
	Email             string `json:"email" gorm:"type:varchar(255);not null"`
	Password          string `json:"password" gorm:"type:varchar(255);not null"`
	NIK               string `json:"nik" gorm:"type:varchar(255);not null"`
	Address           string `json:"address" gorm:"type:varchar(255)"`
	Verified          bool   `json:"verified" gorm:"type:bool" default:"false"`
	Role              string `json:"role" gorm:"type:varchar(255);not null;default:customer"`
	Kategori          string `json:"kategori" gorm:"type:varchar(255);not null"`
	Image             string `json:"image" gorm:"type:varchar(255)"`
	NoHP              string `json:"no_hp" gorm:"type:varchar(255)"`
	VerificationToken string `json:"verification_token" gorm:"type:varchar(255);unique" `
}
