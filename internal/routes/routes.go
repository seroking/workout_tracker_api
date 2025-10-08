package routes

import (
	"workout_tracker/internal/handlers"
	"workout_tracker/internal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api")

	api.POST("/login", func(c *gin.Context) {
		handlers.SignIn(c, db)
	})

	api.POST("/register", func(c *gin.Context) {
		handlers.SignUp(c, db)
	})

	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware())

	// --- Workout Plans ---
	protected.GET("/workout-plan/:id", func(c *gin.Context) { handlers.GetWorkoutPlan(c, db) })
	protected.GET("/workout-plans", func(c *gin.Context) { handlers.ListWorkoutPlan(c, db) })
	protected.POST("/workout-plan", func(c *gin.Context) { handlers.CreateWorkoutPlan(c, db) })
	protected.DELETE("/workout-plan/:id", func(c *gin.Context) { handlers.DeleteWorkoutPlan(c, db) })
	protected.PUT("/workout-plan/:id", func(c *gin.Context) { handlers.UpdateWorkoutPlan(c, db) })

	// --- Workout Exercice ---
	protected.GET("/workout-exercice/:id", func(c *gin.Context) { handlers.GetWorkoutExercice(c, db) })
	protected.GET("/workout-exercices", func(c *gin.Context) { handlers.ListWorkoutExercices(c, db) })
	protected.POST("/workout-exercice", func(c *gin.Context) { handlers.CreateWorkoutExercice(c, db) })
	protected.DELETE("/workout-exercice/:id", func(c *gin.Context) { handlers.DeleteWorkoutExercice(c, db) })
	protected.PUT("/workout-exercice/:id", func(c *gin.Context) { handlers.UpdateWorkoutExercice(c, db) })

	// --- Schedule ---
	protected.GET("/schedule/:id", func(c *gin.Context) { handlers.GetSchedule(c, db) })
	protected.GET("/schedules", func(c *gin.Context) { handlers.ListSchedules(c, db) })
	protected.POST("/schedule", func(c *gin.Context) { handlers.CreateSchedule(c, db) })
	protected.DELETE("/schedule/:id", func(c *gin.Context) { handlers.DeleteSchedule(c, db) })
	protected.PUT("/schedule/:id", func(c *gin.Context) { handlers.UpdateSchedule(c, db) })

	admin := protected.Group("/admin")
	admin.Use(middlewares.AdminOnly(db))

	admin.GET("/users", func(c *gin.Context) { handlers.ListUsers(c, db) })
}
