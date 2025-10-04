package handlers

import (
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSchedule(c *gin.Context, db *gorm.DB) {
	type createData struct {
		Days          []string `json:"days"`
		IsActive      bool     `json:"is_active"`
		IsRecurring   bool     `json:"is_recurring"`
		WorkoutPlanID uint     `json:"workout_plan_id"`
	}
	var input createData
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	schedule = models.Schedule{
		Days:          input.Days,
		IsActive:      input.IsActive,
		IsRecurring:   input.IsRecurring,
		WorkoutPlanID: input.WorkoutPlanID,
	}
	if err := db.Create(&schedule).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to create schedule"})
		return
	}
	c.JSON(200, gin.H{"message": "Schedule created sucessfully"})
}

func GetSchedule(c *gin.Context, db *gorm.DB) {
	var id = c.Param("id")
	var schedule models.Schedule

	if err := db.First(&schedule, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Schedule not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Schedule retrieved successfully", "data": schedule})
}

func ListSchedule(c *gin.Context, db *gorm.DB) {
	var schedules []models.Schedule

	if err := db.Find(&schedules).Error; err != nil {
		c.JSON(404, gin.H{"error": "Schedules not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Schedules retrieved successfully", "data": schedules})
}

func DeleteSchedule(c *gin.Context, db *gorm.DB) {
	var id = c.Param("id")
	var schedule models.Schedule
	if err := db.First(&schedule, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Schedule not found"})
		return
	}

	if err := db.Delete(&schedule).Error; err != nil {
		c.JSON(400, gin.H{"error": "Failed to delete schedule"})
		return
	}
	c.JSON(200, gin.H{"message": "Schedule deleted successfully"})
}

func UpdateSchedule(c *gin.Context, db *gorm.DB) {
	type updateData struct {
		Days        []string `json:"days"`
		IsActive    bool     `json:"is_active"`
		IsRecurring bool     `json:"is_recurring"`
	}
	id := c.Param("id")
	var input updateData
	var Schedule models.Schedule

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.First(&Schedule, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Schedule not found"})
		return
	}

	updates := map[string]interface{}{
		"days":         input.Days,
		"is_active":    input.IsActive,
		"is_recurring": input.IsRecurring,
	}

	if err := db.Model(&Schedule).Updates(updates).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Schedule updated successfully"})
}
