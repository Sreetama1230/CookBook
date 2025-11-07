package db

import (
	"CookBook/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "sreetama:password@tcp(127.0.0.1:3306)/cookbook?charset=utf8mb4&parseTime=True&loc=Local"
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
