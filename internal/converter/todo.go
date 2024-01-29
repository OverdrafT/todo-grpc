package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/OverdrafT/todo-grpc/internal/model"
	desc "github.com/OverdrafT/todo-grpc/pkg/todo_v1"
)

func ToNoteFromService(todo *model.TodoNote) *desc.TodoNote {
	var updatedAt *timestamppb.Timestamp
	if todo.UpdatedAt.Valid {
		updatedAt = timestamppb.New(todo.UpdatedAt.Time)
	}

	return &desc.TodoNote{
		Id:        todo.ID,
		Info:      ToNoteInfoFromService(todo.Info),
		CreatedAt: timestamppb.New(todo.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToNoteInfoFromService(info model.TodoNoteInfo) *desc.TodoNoteInfo {
	return &desc.TodoNoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

func ToNoteInfoFromDesc(info *desc.TodoNoteInfo) *model.TodoNoteInfo {
	return &model.TodoNoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}

func ToUpdateNoteInfoFromDesc(info *desc.UpdateTodoNoteInfo) *model.TodoNoteInfo {
	return &model.TodoNoteInfo{
		Title:   info.Title.Value,
		Content: info.Content.Value,
	}
}
