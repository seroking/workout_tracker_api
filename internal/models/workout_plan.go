package models

type WorkoutPlan struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	User        *User  `json:"user"`

	WorkoutExercices []WorkoutExercice `json:"workout_exercices" gorm:"foreignKey:WorkoutPlanID"`
}
