package repository

import (
	"context"

	"github.com/OverdrafT/todo-grpc/internal/model"
)

type TodoRepository interface {
	Create(ctx context.Context, info *model.TodoNoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.TodoNote, error)
	Update(ctx context.Context, id int64, info *model.TodoNoteInfo) error
	Delete(ctx context.Context, id int64) error
	//List(ctx context.Context) error
}
