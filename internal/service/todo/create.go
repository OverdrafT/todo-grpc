package todo

import (
	"context"

	"github.com/OverdrafT/todo-grpc/internal/model"
)

func (s *serv) Create(ctx context.Context, info *model.TodoNoteInfo) (int64, error) {
	id, err := s.todoRepository.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, nil
}
