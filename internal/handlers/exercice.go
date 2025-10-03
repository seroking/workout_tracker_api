package handlers

import (
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateData struct {
	Name        string
	Description string
	CategoryID  uint
}

func CreateExercice(c *gin.Context, db gorm.DB) {
	var input CreateData
	var checkExerciceExist models.Exercice
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := db.Where("name = ?", input.Name).First(&checkExerciceExist).Error; err == nil {
		c.JSON(409, gin.H{"error": "Exercice already exists"})
		return
	}

	exercice := models.Exercice{
		Name:        input.Name,
		Description: input.Description,
		CategoryID:  input.CategoryID,
	}

	if err := db.Create(&exercice).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to create exercice"})
		return
	}

	c.JSON(201, gin.H{"message": "Exercice Created sucessfully"})
}

func GetExercice(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var exercice models.Exercice
	if err := db.First(&exercice, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Exercice not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Exercice retrieved successfully", "data": exercice})
}

func ListExercices(c *gin.Context, db *gorm.DB) {
	var exercices []models.Exercice

	if err := db.Find(&exercices).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to retrieve exercices"})
		return
	}

	c.JSON(200, gin.H{"message": "Exercices retrieved successfully", "data": exercices})
}

func DeleteExercice(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var exercice models.Exercice

	if err := db.First(&exercice, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Exercice not found"})
		return
	}

	if err := db.Delete(&exercice).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete exercice"})
		return
	}
	c.JSON(200, gin.H{"message": "Exercice deleted successfully"})

}

func UpdateExercice(c *gin.Context, db gorm.DB) {
	type UpdateInput struct {
		Name        string
		Description string
	}
	id := c.Param("id")
	var input UpdateInput
	var exercice models.Exercice

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&exercice, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var updates = map[string]interface{}{
		"name":        input.Name,
		"description": input.Description,
	}

	if err := db.Model(&exercice).Updates(&updates).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to update exercice"})
		return
	}

	c.JSON(200, gin.H{"message": "Exercice updated successfully"})
}
