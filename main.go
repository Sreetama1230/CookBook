package main

import (
	"CookBook/db"
	"CookBook/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	r := gin.Default()

	r.POST("/recipes", handler.CreateRecipe)
	r.GET("/recipes/:id", handler.GetRecipe)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
