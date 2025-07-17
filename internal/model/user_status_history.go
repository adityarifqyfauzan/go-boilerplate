package model

type UserStatusHistory struct {
	BaseModel
	UserID       int `json:"user_id"`
	UserStatusID int `json:"user_status_id"`
	CreatedBy    int `json:"created_by"`
}

func (UserStatusHistory) TableName() string {
	return "user_status_histories"
}
