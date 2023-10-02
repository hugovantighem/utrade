package app

import (
	"context"
	"fmt"

	"github.com/hugovantighem/utrade/domain"
	"github.com/sirupsen/logrus"
)

// GetTodo retrieves a todo according to the given query.
func GetTodo(ctx context.Context, repo domain.Repository, query GetTodoQuery) (domain.Todo, error) {
	logrus.Infof("get todo: userID=%q", query.UserID)

	// retrieve todo
	result, err := repo.Find(ctx, query.UserID)
	if err != nil {
		return domain.Todo{}, ApplicationError{Message: fmt.Sprintf("cannot get a todo: %s", err.Error())}
	}

	return result, nil
}
