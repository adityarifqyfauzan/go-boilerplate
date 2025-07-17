package model

type UserStatus struct {
	BaseModel
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (UserStatus) TableName() string {
	return "user_statuses"
}
