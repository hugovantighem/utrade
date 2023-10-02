package app

import (
	"context"
	"fmt"

	"github.com/hugovantighem/utrade/domain"
	"github.com/sirupsen/logrus"
)

// SaveTodo persists a todo according to the given command.
func SaveTodo(ctx context.Context, repo domain.Repository, cmd SaveTodoCmd) error {
	logrus.Infof("save todo: msg=%q, userID=%q", cmd.Msg, cmd.UserID)

	// instanciate todo
	item, err := domain.NewTodo(cmd.Msg, cmd.UserID)
	if err != nil {
		return ApplicationError{Message: fmt.Sprintf("cannot create a todo: %s", err.Error())}
	}

	// persist
	err = repo.Save(ctx, item)
	if err != nil {
		return ApplicationError{Message: fmt.Sprintf("cannot save a todo: %s", err.Error())}
	}

	return nil
}
