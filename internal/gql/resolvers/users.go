package resolvers

import (
	"context"
	"errors"

	"log"

	"github.com/jinzhu/gorm"
	"github.com/ysthey/go-gql-start/internal/gql/models"
	tf "github.com/ysthey/go-gql-start/internal/gql/resolvers/transformations"
	dbm "github.com/ysthey/go-gql-start/internal/orm/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user models.UserInput) (*models.User, error) {
	return userCreateUpdate(r.ORM.DB, user, false)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, uuid string, user models.UserInput) (*models.User, error) {
	return userCreateUpdate(r.ORM.DB, user, true, uuid)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, uuid string) (bool, error) {
	return userDelete(r.ORM.DB, uuid)
}

func (r *queryResolver) Users(ctx context.Context, uuid *string) (*models.Users, error) {
	return userList(r.ORM.DB, uuid)
}

// ## Helper functions

func userCreateUpdate(gdb *gorm.DB, user models.UserInput, update bool, ids ...string) (*models.User, error) {
	dbo, err := tf.GQLInputUserToDBUser(&user, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := gdb.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the user
	} else {
		db = db.Model(&dbo).Where("uuid=?", ids[0]).Update(dbo).First(dbo) // Or update it
	}
	gql, err := tf.DBUserToGQLUser(dbo)
	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()
	return gql, db.Error
}

func userDelete(gdb *gorm.DB, uuid string) (bool, error) {
	db := gdb.New()
	wtm := &models.User{}
	db = db.Where("uuid= ?", uuid).First(wtm)
	rerr := db.RecordNotFound()
	if rerr {
		log.Println("error while deleting", uuid, "user record not found")
		return false, errors.New("user record not found")
	}
	err := db.Delete(wtm).Error
	if err != nil {
		log.Println("error while deleting", uuid, err)
		return false, err
	}

	return true, nil
}

func userList(gdb *gorm.DB, uuid *string) (*models.Users, error) {
	entity := "users"
	whereID := "uuid = ?"
	record := &models.Users{}
	dbRecords := []*dbm.User{}
	db := gdb.New()
	if uuid != nil {
		db = db.Where(whereID, *uuid)
	}
	db = db.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(dbRec); err != nil {
			log.Println(entity, err)
		} else {
			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}
