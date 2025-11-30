package db

import (
	"CookBook/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbName)

	// dsn := "sreetama:password@tcp(127.0.0.1:3306)/cookbook?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open the db connection: %v", err)
	}

	if err := DB.AutoMigrate(&models.Recipe{}, &models.Ingredient{}, &models.RecipeIngredient{}); err != nil {
		log.Fatalf("auto-migrate error: %v", err)
	}

	log.Println("Database connected and migrated successfully!")
}
