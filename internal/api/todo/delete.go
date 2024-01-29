package todo

import (
	"context"
	desc "github.com/OverdrafT/todo-grpc/pkg/todo_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	err := i.todoService.Delete(ctx, req.GetId())
	if err != nil {
		return new(empty.Empty), err
	}

	return new(empty.Empty), nil
}
