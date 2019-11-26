package jobs

import (
	"github.com/jinzhu/gorm"
	"github.com/ysthey/go-gql-start/internal/orm/models"
	"gopkg.in/gormigrate.v1"
)

var (
	fname     string       = "test"
	lname     string       = "test"
	firstUser *models.User = &models.User{
		Email:     "test@test.com",
		UUID:      "97e3e85e-6ad6-425c-b70b-2eb1275d8d1e",
		Firstname: &fname,
		Lastname:  &lname,
	}
)

// SeedUsers inserts the first users
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		return db.Create(&firstUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(&firstUser).Error
	},
}
