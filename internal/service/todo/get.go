package todo

import (
	"context"

	"github.com/OverdrafT/todo-grpc/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.TodoNote, error) {
	note, err := s.todoRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return note, nil
}
