package todo

import (
	"github.com/OverdrafT/todo-grpc/internal/repository"
	"github.com/OverdrafT/todo-grpc/internal/service"
)

type serv struct {
	todoRepository repository.TodoRepository
}

func NewService(
	todoRepository repository.TodoRepository,
) service.TodoService {
	return &serv{
		todoRepository: todoRepository,
	}
}

func NewMockService(deps ...interface{}) service.TodoService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.TodoRepository:
			srv.todoRepository = s
		}
	}

	return &srv
}
