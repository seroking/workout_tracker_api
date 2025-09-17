package models

type User struct {
	BaseModel
	Username     string `json:"username" gorm:"uniqueIndex;type:varchar(50)"`
	Email        string `json:"email" gorm:"uniqueIndex;type:varchar(255)"`
	PasswordHash string `json:"-"`
	Role         string `json:"role" gorm:"type:enum('user','admin');default:'user'"`
}
