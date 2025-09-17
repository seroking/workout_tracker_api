package handlers

import (
	"os"
	"time"
	"workout_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SignUpInp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUp(c *gin.Context, db *gorm.DB) {
	var input SignUpInp
	var ExistingUser models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if input.Username == "" || input.Email == "" || input.Password == "" {
		c.JSON(400, gin.H{"error": "Invalid Input"})
		return
	}

	if err := db.Where("username = ?", input.Username).First(&ExistingUser).Error; err == nil {
		c.JSON(409, gin.H{"error": "this username already exists"})
		return
	}
	if err := db.Where("email = ?", input.Email).First(&ExistingUser).Error; err == nil {
		c.JSON(409, gin.H{"error": "this email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": "failed to hash Password"})
		return
	}

	User := models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}
	if err := db.Create(&User).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(200, gin.H{
		"message": "User created successfully",
	})
}

type SignInInp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context, db *gorm.DB) {
	var input SignInInp
	var user models.User

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		c.JSON(500, gin.H{"error": "secret key not configured"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid Input"})
		return
	}

	err := db.Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid crendentials"})
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to generate token"})
		return
	}
	c.JSON(200, gin.H{"message": "Login successful!", "token": signedToken})

}
