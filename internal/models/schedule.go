package models

type Schedule struct {
	BaseModel
	Days         []string
	Is_active    bool
	Is_reccuring bool
}
