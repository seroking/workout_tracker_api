package config

import (
	"log"
	"os"
	"workout_tracker/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}

	dsn := os.Getenv("DB_DSN")
	// the DB is not automatically created automatically, create it befoore running the code
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Exercice{}, &models.WorkoutPlan{}, &models.WorkoutExercice{}, &models.Schedule{}, &models.WorkoutSession{})
	DB = db
}
