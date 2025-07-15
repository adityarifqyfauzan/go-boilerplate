package seeders

import (
	"github.com/adityarifqyfauzan/go-boilerplate/internal/helper/constant"
	"github.com/adityarifqyfauzan/go-boilerplate/internal/model"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s UserSeeder) Run(tx *gorm.DB) error {
	// super-admin and admin
	users := []model.User{
		{
			Email:        "admin@localhost.com",
			Password:     "admin",
			Name:         "Admin",
			UserStatusID: constant.USER_STATUS_ACTIVE_ID,
		},
		{
			Email:        "super-admin@localhost.com",
			Password:     "super-admin",
			Name:         "Super Admin",
			UserStatusID: constant.USER_STATUS_ACTIVE_ID,
		},
	}

	if err := tx.Create(&users).Error; err != nil {
		return err
	}

	// user detail
	userDetails := []model.UserDetail{
		{
			UserID: users[0].ID,
		},
		{
			UserID: users[1].ID,
		},
	}

	if err := tx.Create(&userDetails).Error; err != nil {
		return err
	}

	// user roles assign
	userRoles := []model.UserRole{
		{
			UserID: users[0].ID,
			RoleID: constant.ROLE_ADMIN_ID,
		},
		{
			UserID: users[1].ID,
			RoleID: constant.ROLE_SUPER_ADMIN_ID,
		},
	}

	if err := tx.Create(&userRoles).Error; err != nil {
		return err
	}

	// user status history
	userStatusHistory := []model.UserStatusHistory{
		{
			UserID:       users[0].ID,
			UserStatusID: constant.USER_STATUS_ACTIVE_ID,
		},
		{
			UserID:       users[1].ID,
			UserStatusID: constant.USER_STATUS_ACTIVE_ID,
		},
	}

	if err := tx.Create(&userStatusHistory).Error; err != nil {
		return err
	}

	return nil
}
