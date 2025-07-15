package model

type UserStatus struct {
	BaseModel
	Name string
	Slug string
}

func (UserStatus) TableName() string {
	return "user_statuses"
}
