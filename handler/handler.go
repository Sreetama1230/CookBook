package handler

import (
	"CookBook/db"
	"CookBook/models"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type ingredient struct {
	Name string `json:"name"`
}
type recipeIngredients struct {
	Quantity   float64    `json:"quantity"`
	Ingredient ingredient `json:"ingredient"`
}
type reqRecipe struct {
	Title             string              `json:"title"`
	RecipeIngredients []recipeIngredients `json:"recipe_ingredients"`
}

func CreateRecipe(c *gin.Context) {
	var input reqRecipe
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request playload"})
		return
	}
	var ris []models.RecipeIngredient
	for _, r := range input.RecipeIngredients {

		ri := models.RecipeIngredient{
			Quantity: r.Quantity,
			Ingredient: models.Ingredient{
				Name: r.Ingredient.Name,
			},
		}

		ris = append(ris, ri)
	}
	r := models.Recipe{
		Title:             input.Title,
		RecipeIngredients: ris,
	}
	if err := db.DB.Create(&r).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to save the recipe", "detail": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, r)
}

func GetRecipe(c *gin.Context) {
	providedId := c.Param("id")
	id, err := strconv.Atoi(providedId)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "provided id is invalid"})
		return
	}
	var recipe models.Recipe
	/*
		“Load the RecipeIngredients association on the Recipe model,
		and for each RecipeIngredient, also preload its associated Ingredient.”
	*/
	db.DB.Model(&models.Recipe{}).Preload("RecipeIngredients.Ingredient").Find(&recipe, id)

	c.JSON(http.StatusOK, recipe)
}
