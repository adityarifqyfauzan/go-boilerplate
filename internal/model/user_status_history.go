package model

type UserStatusHistory struct {
	BaseModel
	UserID       int
	UserStatusID int
	CreatedBy    int
}

func (UserStatusHistory) TableName() string {
	return "user_status_histories"
}
