package models

type WorkoutExercice struct {
	BaseModel
	Set           int  `json:"set"`
	Rep           int  `json:"rep"`
	Weight        int  `json:"weight"`
	WorkoutPlanID uint `json:"workout_plan_id" gorm:"index:idx_plan_exercice,unique;not null"`
	ExerciceID    uint `json:"exercice_id" gorm:"index:idx_plan_exercice,unique;not null"`

	Exercice    *Exercice    `json:"exercice"`
	WorkoutPlan *WorkoutPlan `json:"workout_plan"`
}
