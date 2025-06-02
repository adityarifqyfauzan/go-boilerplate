package model

type UserRole struct {
	BaseModel
}

func (UserRole) TableName() string {
	return "user_roles"
}
