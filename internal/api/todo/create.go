package todo

import (
	"context"
	"github.com/OverdrafT/todo-grpc/internal/converter"
	desc "github.com/OverdrafT/todo-grpc/pkg/todo_v1"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.todoService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted todo with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
