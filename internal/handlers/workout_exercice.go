package handlers

import (
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWorkoutExercice(c *gin.Context, db *gorm.DB) {
	type createData struct {
		Set           int  `json:"set"`
		Rep           int  `json:"rep"`
		Weight        int  `json:"weight"`
		WorkoutPlanID uint `json:"workout_plan_id"`
		ExerciceID    uint `json:"exercice_id"`
	}

	var input createData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var workoutExercice = models.WorkoutExercice{
		Set:           input.Set,
		Rep:           input.Rep,
		Weight:        input.Weight,
		WorkoutPlanID: input.WorkoutPlanID,
		ExerciceID:    input.ExerciceID,
	}

	if err := db.Create(&workoutExercice).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to create workout exercice"})
		return
	}
	c.JSON(200, gin.H{"message": "workout exercice created successfully"})
}

func GetWorkoutExercice(c *gin.Context, db *gorm.DB) {
	var id = c.Param("id")
	var workoutExercice models.WorkoutExercice

	if err := db.First(&workoutExercice, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout exercice not found"})
		return
	}

	c.JSON(200, gin.H{"message": "workout Plan retrieved successfully", "data": workoutExercice})
}

func ListWorkoutExercices(c *gin.Context, db *gorm.DB) {

	var workoutExercices []models.WorkoutExercice

	if err := db.Find(&workoutExercices).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout exercice not found"})
		return
	}

	c.JSON(200, gin.H{"message": "workout exercices retrieved successfully", "data": workoutExercices})
}

func DeleteWorkoutExercice(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var workoutExercice models.WorkoutExercice

	if err := db.First(&workoutExercice, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout exercice not found"})
		return
	}

	if err := db.Delete(&workoutExercice).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to delete workout exercice"})
		return
	}
	c.JSON(200, gin.H{"message": "workout exercice deleted successfully"})
}

func UpdateWorkoutExercice(c *gin.Context, db *gorm.DB) {
	type updateData struct {
		Set    int `json:"set"`
		Rep    int `json:"rep"`
		Weight int `json:"weight"`
	}
	var input updateData
	var id = c.Param("id")
	var workoutExercice models.WorkoutExercice

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&workoutExercice, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout exercice not found"})
		return
	}

	var update = map[string]interface{}{
		"set":    input.Set,
		"rep":    input.Rep,
		"weight": input.Weight,
	}

	if err := db.Model(&workoutExercice).Updates(update).Error; err != nil {
		c.JSON(400, gin.H{"error": "failed to update workout exercice"})
		return
	}

	c.JSON(200, gin.H{"message": "workout exercice updated successfully"})
}
