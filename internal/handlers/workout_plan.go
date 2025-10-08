package handlers

import (
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWorkoutPlan(c *gin.Context, db *gorm.DB) {
	type CreateData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		UserID      uint   `json:"user_id"`
	}
	var input CreateData
	var checkDataExist models.WorkoutPlan
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("name = ? AND user_id = ?", input.Name, input.UserID).First(&checkDataExist).Error; err == nil {
		c.JSON(409, gin.H{"error": "Workout Plan already exist"})
		return
	}
	workoutPlan := models.WorkoutPlan{
		Name:        input.Name,
		Description: input.Description,
		UserID:      input.UserID,
	}

	if err := db.Create(&workoutPlan).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to create the workout plan"})
		return
	}
	c.JSON(201, gin.H{"message": "Workout Plan created successfully"})
}

func GetWorkoutPlan(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var workoutPlan models.WorkoutPlan

	if err := db.Preload("WorkoutExercices.Exercice.Category").First(&workoutPlan, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout plan not found"})
		return
	}
	c.JSON(200, gin.H{"message": "workout plan retrieved sucessfully", "data": workoutPlan})
}

func ListWorkoutPlan(c *gin.Context, db *gorm.DB) {
	var workoutPlans []models.WorkoutPlan

	if err := db.Preload("User").Find(&workoutPlans).Error; err != nil {
		c.JSON(404, gin.H{"error": "failed to retrieve workout plans"})
		return
	}
	c.JSON(200, gin.H{"message": "workout plans retrieved successfully", "data": workoutPlans})
}

func DeleteWorkoutPlan(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var workoutPlan models.WorkoutPlan

	if err := db.First(&workoutPlan, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout plan not found"})
		return
	}

	if err := db.Delete(&workoutPlan).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to delete workout plan"})
		return
	}

	c.JSON(200, gin.H{"message": "Workout plan deleted sucessfully"})
}

func UpdateWorkoutPlan(c *gin.Context, db *gorm.DB) {
	type updateData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	id := c.Param("id")
	var input updateData
	var workoutPlan models.WorkoutPlan

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&workoutPlan, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout plan not found"})
		return
	}

	var updates = map[string]interface{}{
		"name":        input.Name,
		"description": input.Description,
	}

	if err := db.Model(&workoutPlan).Updates(updates).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to update workout plan"})
		return
	}
	c.JSON(200, gin.H{"message": "Workout plan updated successfully"})
}
