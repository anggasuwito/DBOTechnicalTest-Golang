package config

import (
	"DBOTechnicalTest-Golang/api/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func ConnDB() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SCHEMA"),
	))
	if err != nil {
		log.Fatal("Error connect DB 1 : ", err.Error())
	}

	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Order{})
	db.Model(&models.Order{}).AddForeignKey("customer_id", "customer(id)", "RESTRICT", "RESTRICT")
	return db
}
