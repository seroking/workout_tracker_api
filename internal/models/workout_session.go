package models

import "time"

type WorkoutSession struct {
	BaseModel
	Date          time.Time `json:"date"`
	ScheduleID    uint      `json:"schedule_id"`
	WorkoutPlanID uint      `json:"workout_plan_id"`
}
