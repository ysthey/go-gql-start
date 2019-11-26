// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
	"log"
	"os"

	"github.com/ysthey/go-gql-start/internal/orm/migration"
	"github.com/ysthey/go-gql-start/pkg/config"

	//Imports the database dialect of choice
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

func init() {
	env := os.Getenv("APP_ENV")
	if env == "UNITTEST" {
		return
	}
	dialect = config.MustGet("gorm-dialect")
	dsn = config.MustGet("gorm-conn-dsn")
	seedDB = config.MustGetBool("gorm-seed-db")
	logMode = config.MustGetBool("gorm-log-mode")
	autoMigrate = config.MustGetBool("gorm-auto-migrate")
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(logMode)
	// Automigrate tables
	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Println("[ORM] Database connection initialized.")
	return orm, err
}
