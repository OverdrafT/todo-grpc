package service

import (
	"context"

	"github.com/OverdrafT/todo-grpc/internal/model"
)

type TodoService interface {
	Create(ctx context.Context, info *model.TodoNoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.TodoNote, error)
}
