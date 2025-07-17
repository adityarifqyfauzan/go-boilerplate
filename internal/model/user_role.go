package model

type UserRole struct {
	BaseModel
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
