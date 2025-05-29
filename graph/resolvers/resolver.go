package resolvers

import "github.com/mauricio-pagarme/graphql-teste/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TransactionTypeDB *database.TransactionType
	UserDB            *database.User
	TransactionDB     *database.Transaction
}
