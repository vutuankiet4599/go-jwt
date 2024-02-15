package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vutuankiet4599/go-jwt/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		panic("Failed to load environment variable")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}

	db.AutoMigrate(&models.User{}, &models.Book{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
    dbSQL, err := db.DB()
    if err != nil {
        panic("Failed to close connection from database")
    }
    dbSQL.Close()
}
