package seeders

import (
	"gorm.io/gorm"
)

// register all seeders here 👇
func RegisterSeeders() {
	Register(RoleSeeder{})
	Register(UserStatusSeeder{})
	Register(UserSeeder{})
}

type Seeder interface {
	Run(tx *gorm.DB) error
}

var seeders []Seeder

func Register(s Seeder) {
	seeders = append(seeders, s)
}

func GetSeeders() []Seeder {
	return seeders
}
