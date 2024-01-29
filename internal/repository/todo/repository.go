package todo

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/OverdrafT/todo-grpc/pkg/db"

	"github.com/OverdrafT/todo-grpc/internal/model"
	"github.com/OverdrafT/todo-grpc/internal/repository"
	"github.com/OverdrafT/todo-grpc/internal/repository/todo/converter"
	modelRepo "github.com/OverdrafT/todo-grpc/internal/repository/todo/model"
)

const (
	tableName = "todo"

	idColumn        = "id"
	titleColumn     = "title"
	contentColumn   = "content"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.TodoRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.TodoNoteInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(titleColumn, contentColumn).
		Values(info.Title, info.Content).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "todo_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.TodoNote, error) {
	builder := sq.Select(idColumn, titleColumn, contentColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "todo_repository.Get",
		QueryRaw: query,
	}

	var note modelRepo.TodoNote
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&note.ID, &note.Info.Title, &note.Info.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToTodoFromRepo(&note), nil
}

func (r *repo) Update(ctx context.Context, id int64, info *model.TodoNoteInfo) error {
	builder := sq.Update(tableName).Set(contentColumn, info.Content).Set(titleColumn, info.Title).Set(updatedAtColumn, time.Now()).Where(sq.Eq{idColumn: id}).PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "todo_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "todo_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) List(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
