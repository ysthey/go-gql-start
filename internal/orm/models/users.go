package models

import "github.com/jinzhu/gorm"

// User defines a user for the app
type User struct {
	gorm.Model
	// We don't to actually delete the users, maybe audit if we want to hard delete them? or wait x days to purge from the table, also
	Email string `gorm:"not null;unique_index:idx_email"`
	UUID  string `gorm:"not null;unique_index:idx_uuid"` // External UUID
}
