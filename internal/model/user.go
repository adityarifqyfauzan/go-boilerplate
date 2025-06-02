package model

type User struct {
	BaseModel
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (User) TableName() string {
	return "users"
}
