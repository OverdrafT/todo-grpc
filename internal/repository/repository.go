package repository

import (
	"context"

	"github.com/OverdrafT/todo-grpc/internal/model"
)

type TodoRepository interface {
	Create(ctx context.Context, info *model.TodoNoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.TodoNote, error)
}
