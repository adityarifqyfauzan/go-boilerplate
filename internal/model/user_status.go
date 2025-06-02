package model

type UserStatus struct {
	BaseModel
}

func (UserStatus) TableName() string {
	return "user_statuses"
}
