package todo

import (
	"context"
	"github.com/OverdrafT/todo-grpc/internal/converter"
	desc "github.com/OverdrafT/todo-grpc/pkg/todo_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	err := i.todoService.Update(ctx, req.GetId(), converter.ToUpdateNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return new(empty.Empty), nil
}
