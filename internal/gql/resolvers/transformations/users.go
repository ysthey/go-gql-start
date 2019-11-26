package transformations

import (
	"errors"

	"github.com/google/uuid"
	gql "github.com/ysthey/go-gql-start/internal/gql/models"
	dbm "github.com/ysthey/go-gql-start/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		Email:     i.Email,
		UUID:      i.UUID,
		Firstname: i.Firstname,
		Lastname:  i.Lastname,
	}
	return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.UserInput, update bool, ids ...string) (o *dbm.User, err error) {
	o = &dbm.User{
		Firstname: i.Firstname,
		Lastname:  i.Lastname,
	}

	if !update {
		//create new user
		if i.Email == nil {
			return nil, errors.New("Field [email] is required")
		}
		u, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		o.UUID = u.String()
	} else {
		o.UUID = ids[0]

	}

	if i.Email != nil {
		o.Email = *i.Email
	}

	return o, err
}
