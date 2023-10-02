package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hugovantighem/utrade/app"
	"github.com/hugovantighem/utrade/domain/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSaveTodo(t *testing.T) {
	t.Run("Errors", func(t *testing.T) {
		t.Run("Validation", func(t *testing.T) {
			// GIVEN a SaveTodoCmd with wrong message
			cmd := app.SaveTodoCmd{
				Msg:    "",
				UserID: uuid.NewString(),
			}

			// WHEN SaveTodo
			err := app.SaveTodo(context.Background(), nil, cmd)

			// THEN an application error is raised
			assert.Error(t, err)
			assert.ErrorAs(t, err, &app.ApplicationError{})
		})

		t.Run("Dependency", func(t *testing.T) {
			// GIVEN a SaveTodoCmd
			cmd := app.SaveTodoCmd{
				Msg:    "foobar",
				UserID: uuid.NewString(),
			}

			// AND an expected call to repo.Save
			repo := mocks.NewMockRepository(gomock.NewController(t))
			repo.EXPECT().Save(gomock.Any(), gomock.Any()).
				Return(fmt.Errorf("error"))

			// WHEN SaveTodo
			err := app.SaveTodo(context.Background(), repo, cmd)

			// THEN an application error is raised
			assert.Error(t, err)
			assert.ErrorAs(t, err, &app.ApplicationError{})
		})

	})

	t.Run("Success", func(t *testing.T) {
		// GIVEN a SaveTodoCmd
		cmd := app.SaveTodoCmd{
			Msg:    "FooBar",
			UserID: uuid.NewString(),
		}

		// AND an expected call to repo.Save
		repo := mocks.NewMockRepository(gomock.NewController(t))
		repo.EXPECT().Save(gomock.Any(), gomock.Any()).
			Return(nil)

		// WHEN SaveTodo
		err := app.SaveTodo(context.Background(), repo, cmd)

		// THEN no error is raised
		assert.NoError(t, err)
	})

}
