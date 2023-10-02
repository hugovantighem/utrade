package domain

import (
	"context"
)

//go:generate mockgen -destination=mocks/mock_Repository.go -package=mocks github.com/hugovantighem/utrade/domain Repository

type Repository interface {
	// Save persists the todo.
	Save(ctx context.Context, item Todo) error
	// Find return the Todo matching the given userID.
	Find(ctx context.Context, userID string) (Todo, error)
	// List returns a list of Todos matching the given status.
	List(ctx context.Context, status TodoStatus) ([]Todo, error)
}
