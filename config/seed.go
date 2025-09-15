package config

import (
	"errors"
	"log"
	"os"
	"workout_tracker/internal/models"

	// "github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var admin models.User
	// load the env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("failed to load the .env file")
	// }
	//verify if the data already exist
	result := db.First(&admin, "role = ?", "admin")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		name := os.Getenv("ADMIN_NAME")
		email := os.Getenv("ADMIN_MAIL")
		password := os.Getenv("ADMIN_PASSWORD")
		role := "admin"
		if name == "" || email == "" || password == "" {
			log.Fatal("Missing admin credentials")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("Password check failed", err)
		}

		admin = models.User{
			Username:     name,
			Email:        email,
			PasswordHash: string(hashedPassword),
			Role:         role,
		}
		if err := db.Create(&admin).Error; err != nil {
			log.Fatal("Failed to create admin user", err)
		}
		println("Admin User Seeded successfully.")
	}

}
