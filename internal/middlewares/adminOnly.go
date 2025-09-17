package middlewares

import (
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminOnly(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var User models.User
		userID, exist := c.Get("user_id")
		if !exist {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid User ID"})
			return
		}
		err := db.Where("ID = ?", userID).First(&User).Error
		if err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": "User Not found"})
			return
		}

		if User.Role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Admins Only"})
			return
		}

		c.Next()
	}
}
