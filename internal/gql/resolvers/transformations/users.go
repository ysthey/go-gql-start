package transformations

import (
	"errors"
	"log"

	"github.com/google/uuid"
	gql "github.com/ysthey/go-gql-start/internal/gql/models"
	dbm "github.com/ysthey/go-gql-start/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		Email: i.Email,
		UUID:  i.UUID,
	}
	return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.UserInput, update bool, ids ...string) (o *dbm.User, err error) {
	o = &dbm.User{}

	if !update {
		//create new user
		if len(i.Email) == 0 {
			return nil, errors.New("Field [email] is required")
		}
		u, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		o.UUID = u.String()
		log.Println(o.UUID)
	} else {
		o.UUID = ids[0]

	}

	if len(i.Email) > 0 {
		o.Email = i.Email
		log.Println(o.Email)
	}

	return o, err
}
