package todo

import (
	"github.com/OverdrafT/todo-grpc/internal/service"
	desc "github.com/OverdrafT/todo-grpc/pkg/todo_v1"
)

type Implementation struct {
	desc.UnimplementedTodoV1Server
	todoService service.TodoService
}

func NewImplementation(todoService service.TodoService) *Implementation {
	return &Implementation{
		todoService: todoService,
	}
}
