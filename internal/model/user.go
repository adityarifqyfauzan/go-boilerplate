package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	SoftDelete
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"-"`
	UserStatusID  int    `json:"user_status_id"`
	RememberToken string `json:"remember_token"`
}

func (User) TableName() string {
	return "users"
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	m.Password = string(hashedPassword)
	return nil
}
