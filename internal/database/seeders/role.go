package seeders

import (
	"github.com/adityarifqyfauzan/go-boilerplate/internal/model"
	"gorm.io/gorm"
)

type RoleSeeder struct{}

func (s RoleSeeder) Run(tx *gorm.DB) error {
	roles := []model.Role{
		{
			Name:     "Super Admin",
			Slug:     "super-admin",
			IsActive: true,
		},
		{
			Name:     "Admin",
			Slug:     "admin",
			IsActive: true,
		},
		{
			Name:     "User",
			Slug:     "user",
			IsActive: true,
		},
	}

	if err := tx.Create(&roles).Error; err != nil {
		return err
	}

	return nil
}
