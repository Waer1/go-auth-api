package config

import (
	"api-auth/pkg/user/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.DbUser,
		Config.DbPassword,
		Config.DbHost,
		Config.DbPort,
		Config.DbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database

	// Automigrate the schema
	if err := DB.AutoMigrate(&models.User{}); err != nil { // Ensure you migrate all necessary models
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database connected and migrated")
}
