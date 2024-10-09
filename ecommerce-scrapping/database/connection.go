// database/db.go
package database

import (
	"fmt"

	"github.com/syahlan1/ecommerce-scrapper.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

var DB *gorm.DB

func Connect() {
	host := "localhost"
	username := "postgres"
	password := "password"
	dbname := "ecommerce"
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	DB = connection

	connection.AutoMigrate(
		&models.Product{},
		&models.UserHistory{},
		&models.User{},
	)
}
