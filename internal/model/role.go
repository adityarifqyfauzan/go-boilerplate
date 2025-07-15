package model

type Role struct {
	BaseModel
	Name     string
	Slug     string
	IsActive bool
}

func (Role) TableName() string {
	return "roles"
}
