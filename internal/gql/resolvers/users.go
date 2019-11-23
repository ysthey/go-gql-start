package resolvers

import (
	"context"

	"github.com/ysthey/go-gql-start/internal/gql/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, uuid string) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Users(ctx context.Context, uuid string) ([]*models.User, error) {
	panic("not implemented")
}
