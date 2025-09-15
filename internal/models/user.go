package models

type User struct {
	BaseModel
	Username     string `json:"username" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
}
