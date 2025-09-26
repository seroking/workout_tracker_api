package handlers

import (
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateCategory(c *gin.Context, db *gorm.DB) {
	var input createData
	var CategoryExist models.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("name = ?", input.Name).First(&CategoryExist).Error; err == nil {
		c.JSON(409, gin.H{"error": "Category already exists"})
		return
	}

	category := models.Category{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := db.Create(&category).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to create category"})
		return
	}
	c.JSON(201, gin.H{"message": "Category created successfully"})
}

func GetCategory(c *gin.Context, db *gorm.DB) {
	var id = c.Param("id")
	var category models.Category

	if err := db.First(&category, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Category retrieved successfully", "data": category})
}

func ListCategories(c *gin.Context, db *gorm.DB) {
	var Categories []models.Category
	if err := db.Find(&Categories).Error; err != nil {
		c.JSON(500, gin.H{"error": "failed to retrieve categories"})
		return
	}

	c.JSON(200, gin.H{"message": "Categories retrieved successfully", "data": Categories})
}

func DeleteCategory(c *gin.Context, db *gorm.DB) {
	var id = c.Param("id")
	var category models.Category

	if err := db.First(&category, id); err != nil {
		c.JSON(404, gin.H{"error": "Category not found"})
		return
	}
	db.Delete(&category)

	c.JSON(200, gin.H{"message": "Category deleted successfully"})
}

type UpdateData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func UpdateCategory(c *gin.Context, db *gorm.DB) {
	var input UpdateData
	var checkCategory models.Category
	var category models.Category
	var id = c.Param("id")

	if err := db.First(&checkCategory, id); err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var updates = map[string]interface{}{
		"name":        "",
		"description": "",
	}

	if input.Name != "" {
		updates["name"] = input.Name
	}

	if input.Description != "" {
		updates["description"] = input.Description
	}

	if err := db.Model(&category).Updates(updates).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to update category"})
		return
	}
	c.JSON(200, gin.H{"message": "Category updated successfully"})
}
