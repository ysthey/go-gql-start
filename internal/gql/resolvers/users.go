package resolvers

import (
	"context"

	"github.com/ysthey/go-gql-start/internal/gql/models"
	tf "github.com/ysthey/go-gql-start/internal/gql/resolvers/transformations"
	dbm "github.com/ysthey/go-gql-start/internal/orm/models"
	"log"
)

// CreateUser creates a record
func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, false)
}

// UpdateUser updates a record
func (r *mutationResolver) UpdateUser(ctx context.Context, uuid string, input models.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, true, uuid)
}

// DeleteUser deletes a record
func (r *mutationResolver) DeleteUser(ctx context.Context, uuid string) (bool, error) {
	return userDelete(r, uuid)
}

// Users lists records
func (r *queryResolver) Users(ctx context.Context, uuid *string) (*models.Users, error) {
	return userList(r, uuid)
}

// ## Helper functions

func userCreateUpdate(r *mutationResolver, input models.UserInput, update bool, ids ...string) (*models.User, error) {
	dbo, err := tf.GQLInputUserToDBUser(&input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the user
	} else {
		db = db.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	gql, err := tf.DBUserToGQLUser(dbo)
	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()
	return gql, db.Error
}

func userDelete(r *mutationResolver, uuid string) (bool, error) {
	return false, nil
}

func userList(r *queryResolver, uuid *string) (*models.Users, error) {
	entity := "users"
	whereID := "uuid = ?"
	record := &models.Users{}
	dbRecords := []*dbm.User{}
	db := r.ORM.DB.New()
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
