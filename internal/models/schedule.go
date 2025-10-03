package models

// type Schedule struct {
// 	BaseModel
// 	Days         []string
// 	Is_active    bool
// 	Is_reccuring bool
// }

type Schedule struct {
	BaseModel
	WorkoutPlanID uint         `json:"workout_plan_id"`
	WorkoutPlan   *WorkoutPlan `json:"workout_plan"`

	Days        []string `gorm:"type:json" json:"days"`
	IsActive    bool     `json:"is_active"`
	IsRecurring bool     `json:"is_recurring"`
}
