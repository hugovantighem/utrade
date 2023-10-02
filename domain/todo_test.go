package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hugovantighem/utrade/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTodo(t *testing.T) {
	t.Run("Errors", func(t *testing.T) {
		inputs := map[string]struct {
			msg    string
			userID string
		}{
			"InvalidMessage": {
				msg:    "",
				userID: uuid.NewString(),
			},
			"InvalidUserID": {
				msg:    "foo",
				userID: "",
			},
		}

		for name, input := range inputs {
			t.Run(name, func(t *testing.T) {
				// GIVEN a message
				msg := input.msg
				// AND a userID
				userID := input.userID
				// WHEN creating a Todo
				_, err := domain.NewTodo(msg, userID)
				// THEN a domain error is raised
				assert.Error(t, err)
				assert.ErrorAs(t, err, &domain.InvalidArgumentError{})
			})
		}
	})

	t.Run("Success", func(t *testing.T) {
		// GIVEN a message
		msg := "foo"
		// AND a userID
		userID := "bar"
		// WHEN creating a Todo
		result, err := domain.NewTodo(msg, userID)
		// THEN no error is raised
		require.NoError(t, err)
		// AND the created Todo hase a 'creating' status
		assert.Equal(t, domain.Created, result.Status())
	})

}
