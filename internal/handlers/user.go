package handlers

import (
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateData struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

func CreateUser(c *gin.Context, db *gorm.DB) {
	var input CreateData
	var UserExistCheck models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("username = ?", input.Username).First(&UserExistCheck).Error; err == nil {
		c.JSON(409, gin.H{"error": "Username already exists"})
		return
	}
	if err := db.Where("email = ?", input.Email).First(&UserExistCheck).Error; err == nil {
		c.JSON(409, gin.H{"error": "Email already exists"})
		return
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(HashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(201, gin.H{"message": "User created sucessfully"})
}

func GetUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	id := c.Param("id")

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"message": "User retrieved successfully", "data": user})

}

func DeleteUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	id := c.Param("id")

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User Not Found"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted sucessfully"})
}
