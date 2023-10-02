package db

import (
	"context"
	"sync"

	"github.com/hugovantighem/utrade/domain"
	"github.com/sirupsen/logrus"
)

type InMemoryRepository struct {
	mu    sync.Mutex
	items map[string]domain.Todo
}

func NewInMemoryRepository() domain.Repository {
	return &InMemoryRepository{
		items: make(map[string]domain.Todo),
	}
}

func (x *InMemoryRepository) Save(ctx context.Context, item domain.Todo) error {
	logrus.Debugf("repo Save todo: msg=%q, userID=%q", item.Message(), item.UserID())

	x.mu.Lock()
	x.items[item.UserID()] = item
	x.mu.Unlock()

	return nil
}

func (x *InMemoryRepository) Find(ctx context.Context, userID string) (domain.Todo, error) {
	logrus.Debugf("repo Find todo: userID=%q", userID)

	x.mu.Lock()
	result, ok := x.items[userID]
	x.mu.Unlock()

	if !ok {
		return domain.Todo{}, domain.NotFoundError{Target: "Todo", LookupKey: "userID", LookupValue: userID}
	}

	return result, nil

}

func (x *InMemoryRepository) List(ctx context.Context, status domain.TodoStatus) ([]domain.Todo, error) {
	logrus.Debugf("repo List todos: status=%q", status)

	if !status.IsValid() {
		return []domain.Todo{}, nil
	}

	result := make([]domain.Todo, 0)
	x.mu.Lock()

	for _, item := range x.items {
		if item.Status() == status {
			result = append(result, item)
		}
	}

	x.mu.Unlock()

	return result, nil
}
