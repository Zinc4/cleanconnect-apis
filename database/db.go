package database

import (
	"clean-connect/pkg/entities"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error

	// dsn := "host=pgsql user=postgres password=Zeto.2003 dbname=ceco port=5432 sslmode=disable"
	dsn := "host=ep-green-base-a2cboio4.eu-central-1.pg.koyeb.app user=koyeb-adm password=fHuCSk0PiFa7 dbname=koyebdb"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	DB.Logger = logger.Default.LogMode(logger.Info)

	err = DB.AutoMigrate(&entities.Customer{}, &entities.Bill{}, &entities.Payment{}, &entities.AdditionalBill{})
	if err != nil {
		log.Fatal("Database migration failed", err)
	}

}
