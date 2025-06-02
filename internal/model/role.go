package model

type Role struct {
	BaseModel
}

func (Role) TableName() string {
	return "roles"
}
	