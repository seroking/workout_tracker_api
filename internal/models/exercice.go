package models

type Exercice struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description" gorm:"unique"`
	CategoryID  uint   `json:"category_id" gorm:"foreignKey"`

	Category *Category `json:"category"`
}
