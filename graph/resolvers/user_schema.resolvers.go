package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.73

import (
	"context"
	"fmt"

	"github.com/mauricio-pagarme/graphql-teste/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	fmt.Println(input)
	user, err := r.UserDB.Create(input.Name, input.Cnpj)
	if err != nil {
		panic(err)
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
		Cnpj: user.Cnpj,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserDB.FindAll()
	if err != nil {
		return nil, err
	}

	var usersModel []*model.User
	for _, user := range users {
		usersModel = append(usersModel, &model.User{
			ID:   user.ID,
			Name: user.Name,
			Cnpj: user.Cnpj,
		})
	}

	return usersModel, nil
}
