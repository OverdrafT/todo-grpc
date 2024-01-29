package todo

import (
	"context"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.todoRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
