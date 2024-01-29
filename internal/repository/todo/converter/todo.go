package converter

import (
	"github.com/OverdrafT/todo-grpc/internal/model"
	modelRepo "github.com/OverdrafT/todo-grpc/internal/repository/todo/model"
)

func ToTodoFromRepo(note *modelRepo.TodoNote) *model.TodoNote {
	return &model.TodoNote{
		ID:        note.ID,
		Info:      ToTodoInfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ToTodoInfoFromRepo(info modelRepo.TodoNoteInfo) model.TodoNoteInfo {
	return model.TodoNoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
