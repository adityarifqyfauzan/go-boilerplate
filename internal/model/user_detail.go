package model

type UserDetail struct {
	BaseModel
	UserID      int
	Address     string
	PhoneNumber string
}

func (UserDetail) TableName() string {
	return "user_details"
}
