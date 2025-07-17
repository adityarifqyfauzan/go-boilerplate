package model

type UserDetail struct {
	BaseModel
	UserID      int    `json:"user_id"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

func (UserDetail) TableName() string {
	return "user_details"
}
