package resolvers

import (
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/ysthey/go-gql-start/internal/gql/models"
	"github.com/ysthey/go-gql-start/internal/orm/migration"
)

func createTestDb() *gorm.DB {
	ggdb, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Panic("[Test DB] err: ", err)
	}
	err = migration.ServiceAutoMigration(ggdb)
	if err != nil {
		log.Panic("[Test DB] err: ", err)
	}
	return ggdb

}

func TestCreate(t *testing.T) {
	ggdb := createTestDb()
	defer ggdb.Close()
	email := "hello@hello.com"
	gqlUser, err := userCreateUpdate(ggdb, models.UserInput{Email: &email}, false)
	if err != nil {
		t.Errorf("Failed to create new user: %s", err.Error())
	}

	if gqlUser.Email != email {
		t.Fail()
	}
	if len(gqlUser.UUID) == 0 {
		t.Fail()
	}
}

func TestUpdate(t *testing.T) {
	ggdb := createTestDb()
	defer ggdb.Close()
	email := "hello@hello.com"
	gqlUser, err := userCreateUpdate(ggdb, models.UserInput{Email: &email}, false)
	if err != nil {
		t.Errorf("Failed to create new user: %s", err.Error())
	}
	if gqlUser.Email != email {
		t.Fail()
	}
	if len(gqlUser.UUID) == 0 {
		t.Fail()
	}
	newemail := "new@hello.com"

	gqlUser, err = userCreateUpdate(ggdb, models.UserInput{Email: &newemail}, true, gqlUser.UUID)
	if err != nil {
		t.Errorf("Failed to update new email : %s", err.Error())
	}
	if gqlUser.Email != newemail {
		t.Fail()
	}
}

func TestList(t *testing.T) {
	//97e3e85e-6ad6-425c-b70b-2eb1275d8d1e
	ggdb := createTestDb()
	defer ggdb.Close()
	uuid := "97e3e85e-6ad6-425c-b70b-2eb1275d8d1e"
	users, err := userList(ggdb, &uuid)
	if err != nil {
		t.Errorf("Failed to get seed user: %s", err.Error())
	}
	if users.Count != 1 {
		t.Fail()
	}
	if users.List[0].UUID != uuid {
		t.Fail()
	}
	if users.List[0].Email != "test@test.com" {
		t.Fail()
	}

	if *(users.List[0].Firstname) != "test" {
		t.Fail()
	}

	if *(users.List[0].Lastname) != "test" {
		t.Fail()
	}

}

func TestDelete(t *testing.T) {
	//97e3e85e-6ad6-425c-b70b-2eb1275d8d1e
	ggdb := createTestDb()
	defer ggdb.Close()
	uuid := "97e3e85e-6ad6-425c-b70b-2eb1275d8d1e"
	users, err := userList(ggdb, &uuid)
	if err != nil {
		t.Errorf("Failed to get seed user: %s", err.Error())
	}
	if users.Count != 1 {
		log.Println("wrong user count before delete")
		t.Fail()
	}

	_, err = userDelete(ggdb, uuid)
	if err != nil {
		t.Errorf("Failed to delete seed user: %s", err.Error())
	}

	users, err = userList(ggdb, &uuid)
	if err != nil {
		t.Errorf("Failed to get seed user: %s", err.Error())
	}
	if users.Count != 0 {
		log.Println("wrong user count after delete")
		t.Fail()
	}

}
