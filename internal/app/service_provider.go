package app

import (
	"context"
	"github.com/OverdrafT/todo-grpc/internal/config"
	"log"

	"github.com/OverdrafT/todo-grpc/internal/api/todo"
	"github.com/OverdrafT/todo-grpc/internal/repository"
	todoRepository "github.com/OverdrafT/todo-grpc/internal/repository/todo"
	"github.com/OverdrafT/todo-grpc/internal/service"
	todoService "github.com/OverdrafT/todo-grpc/internal/service/todo"
	"github.com/OverdrafT/todo-grpc/pkg/closer"
	"github.com/OverdrafT/todo-grpc/pkg/db"
	"github.com/OverdrafT/todo-grpc/pkg/db/pg"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	noteRepository repository.TodoRepository

	noteService service.TodoService

	noteImpl *todo.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TodoRepository(ctx context.Context) repository.TodoRepository {
	if s.noteRepository == nil {
		s.noteRepository = todoRepository.NewRepository(s.DBClient(ctx))
	}

	return s.noteRepository
}

func (s *serviceProvider) TodoService(ctx context.Context) service.TodoService {
	if s.noteService == nil {
		s.noteService = todoService.NewService(
			s.TodoRepository(ctx),
		)
	}

	return s.noteService
}

func (s *serviceProvider) NoteImpl(ctx context.Context) *todo.Implementation {
	if s.noteImpl == nil {
		s.noteImpl = todo.NewImplementation(s.TodoService(ctx))
	}

	return s.noteImpl
}
