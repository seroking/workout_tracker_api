package handlers

import(
	"github.com/gin-gonic/gin",
	"gorm.io/gorm"
	"workout_tracker/internal/models"
)

type createData struct{
	name string	`json:"name"`
	description string	`json:"description"`
}

func CreateCategory(c *gin.Context, db *gorm.DB){
	var input createData

	if err := c.ShouldBindJSON(&input).Error; err != nil{
		c.JSON(400, gin.H{"error": err})
		return
	}
}
