package app

import (
	"context"
	"fmt"

	"github.com/hugovantighem/utrade/domain"
	"github.com/sirupsen/logrus"
)

// List retrieves todos according that are currently in the 'created' status.
func ListCreated(ctx context.Context, repo domain.Repository) ([]domain.Todo, error) {
	logrus.Infof("list todos")

	// retrieve todos
	result, err := repo.List(ctx, domain.Created)
	if err != nil {
		return nil, ApplicationError{Message: fmt.Sprintf("cannot list todos: %s", err.Error())}
	}

	return result, nil
}
