package model

type User struct {
	Base
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

func (User) TableName() string {
	return "users"
}
