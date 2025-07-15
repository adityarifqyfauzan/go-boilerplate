package model

import "time"

type BaseModel struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SoftDelete struct {
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
