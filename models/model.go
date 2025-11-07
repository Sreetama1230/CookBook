package models

import (
	"time"

	"gorm.io/gorm"
)

//Topic: Recipe ↔ Ingredient with a join table containing quantity — "CookBook"

type Recipe struct {
	ID                uint               `json:"id" gorm:"primaryKey"`
	Title             string             `json:"title"`
	RecipeIngredients []RecipeIngredient `json:"recipe_ingredients" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Ingredient struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

// join model stores metadata like Quantity
type RecipeIngredient struct {
	RecipeID     uint           `json:"-" gorm:"primaryKey"`
	IngredientID uint           `json:"-" gorm:"primaryKey"`
	Ingredient   Ingredient     `json:"ingredient" gorm:"foreignKey:IngredientID"`
	Quantity     float64        `json:"quantity"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}
