package model

type Role struct {
	BaseModel
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	IsActive bool   `json:"is_active"`
}

func (Role) TableName() string {
	return "roles"
}
