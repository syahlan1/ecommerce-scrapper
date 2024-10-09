package database

import (
    "fmt"
    "os"

    "github.com/syahlan1/ecommerce-scrapper.git/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

)

var DB *gorm.DB

func Connect() {
    // Cek apakah DATABASE_URL ada di environment (untuk Heroku)
    databaseURL := os.Getenv("DATABASE_URL")
    var dsn string

    if databaseURL != "" {
        // Gunakan DATABASE_URL di Heroku
        dsn = databaseURL
    } else {
        // Gunakan koneksi lokal jika DATABASE_URL tidak ada (untuk pengembangan lokal)
        host := "localhost"
        username := "postgres"
        password := "password"
        dbname := "ecommerce"
        port := "5432"

        dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)
    }

    // Buka koneksi ke database
    connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    DB = connection

    // Migrasi model
    connection.AutoMigrate(
        &models.Product{},
        &models.UserHistory{},
        &models.User{},
    )
}
