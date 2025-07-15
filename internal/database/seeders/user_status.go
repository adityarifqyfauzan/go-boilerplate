package seeders

import (
	"github.com/adityarifqyfauzan/go-boilerplate/internal/model"
	"gorm.io/gorm"
)

type UserStatusSeeder struct{}

func (s UserStatusSeeder) Run(tx *gorm.DB) error {
	userStatuses := []model.UserStatus{
		{
			Name: "Pending",
			Slug: "pending",
		},
		{
			Name: "Active",
			Slug: "active",
		},
		{
			Name: "Inactive",
			Slug: "inactive",
		},
	}

	if err := tx.Create(&userStatuses).Error; err != nil {
		return err
	}

	return nil
}
