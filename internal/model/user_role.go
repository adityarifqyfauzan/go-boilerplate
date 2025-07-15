package model

type UserRole struct {
	BaseModel
	UserID int
	RoleID int
}

func (UserRole) TableName() string {
	return "user_roles"
}
