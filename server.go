package main

import (
	"github.com/vutuankiet4599/go-jwt/config"
	"github.com/vutuankiet4599/go-jwt/routes"
	"gorm.io/gorm"
)


var (
	db *gorm.DB = config.SetupDatabaseConnection()
	router = routes.InitApiRouter()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router.Run()
}