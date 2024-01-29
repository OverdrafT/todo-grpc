package todo

import (
	"context"

	"github.com/OverdrafT/todo-grpc/internal/model"
)

func (s *serv) Update(ctx context.Context, id int64, info *model.TodoNoteInfo) error {
	err := s.todoRepository.Update(ctx, id, info)
	if err != nil {
		return err
	}

	return nil
}
