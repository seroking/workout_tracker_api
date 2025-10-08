package handlers

import (
	"time"
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateWorkoutSession(c *gin.Context, db *gorm.DB) {

	type createData struct {
		Date          time.Time `json:"date"`
		ScheduleID    uint      `json:"schedule_id"`
		WorkoutPlanID uint      `json:"workout_plan_id"`
	}
	var input createData
	var workoutSession models.WorkoutSession
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	workoutSession = models.WorkoutSession{
		Date:          input.Date,
		ScheduleID:    input.ScheduleID,
		WorkoutPlanID: input.WorkoutPlanID,
	}

	if err := db.Create(&workoutSession).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to create workout session"})
		return
	}
	c.JSON(200, gin.H{"message": "Workout session created successfully"})
}

func GetWorkoutSession(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var workoutSession models.WorkoutSession

	if err := db.Preload("WorkoutPlan.WorkoutExercices.Exercice.Category").First(&workoutSession, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout session not found"})
		return
	}

	c.JSON(200, gin.H{"message": "workout session retrieved sucessfully", "workout_session": workoutSession})
}

func ListWorkoutSessions(c *gin.Context, db *gorm.DB) {
	var workoutSessions []models.WorkoutSession

	if err := db.Preload("WorkoutPlan").Find(&workoutSessions).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to retrieve workout sessions"})
		return
	}
	c.JSON(200, gin.H{"message": "workout sessions retrieved sucessfully", "workout_sessions": workoutSessions})
}

func DeleteWorkoutSession(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var workoutSession models.WorkoutSession

	if err := db.First(&workoutSession, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout session not found"})
		return
	}
	if err := db.Delete(&workoutSession).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete workout session"})
		return
	}
	c.JSON(200, gin.H{"message": "Workout session deleted sucessfully"})
}

func UpdateWorkoutSession(c *gin.Context, db *gorm.DB) {
	type updateData struct {
		Date time.Time `json:"date"`
	}
	id := c.Param("id")
	var workoutSession models.WorkoutSession
	var input updateData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&workoutSession, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Workout session not found"})
		return
	}

	var updates = map[string]interface{}{
		"date": input.Date,
	}

	if err := db.Model(&workoutSession).Updates(updates).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Workout session updated sucessfully"})
}
