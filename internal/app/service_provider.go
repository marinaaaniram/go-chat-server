package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/marinaaaniram/go-chat-server/internal/api/chat"
	"github.com/marinaaaniram/go-chat-server/internal/api/message"
	"github.com/marinaaaniram/go-chat-server/internal/closer"
	"github.com/marinaaaniram/go-chat-server/internal/config"
	"github.com/marinaaaniram/go-chat-server/internal/repository"
	chatRepository "github.com/marinaaaniram/go-chat-server/internal/repository/chat"
	messageRepository "github.com/marinaaaniram/go-chat-server/internal/repository/message"
	"github.com/marinaaaniram/go-chat-server/internal/service"
	chatService "github.com/marinaaaniram/go-chat-server/internal/service/chat"
	messageService "github.com/marinaaaniram/go-chat-server/internal/service/message"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool            *pgxpool.Pool
	chatRepository    repository.ChatRepository
	messageRepository repository.MessageRepository

	chatService    service.ChatService
	messageService service.MessageService

	chatImpl    *chat.Implementation
	messageImpl *message.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GetPGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) GetPgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.GetPGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) GetChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.GetPgPool(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) GetChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.GetChatRepository(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) GetChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.GetChatService(ctx))
	}

	return s.chatImpl
}

func (s *serviceProvider) GetMessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepository.NewRepository(s.GetPgPool(ctx))
	}

	return s.messageRepository
}

func (s *serviceProvider) GetMessageService(ctx context.Context) service.MessageService {
	if s.messageService == nil {
		s.messageService = messageService.NewService(s.GetMessageRepository(ctx))
	}

	return s.messageService
}

func (s *serviceProvider) GetMessageImpl(ctx context.Context) *message.Implementation {
	if s.messageImpl == nil {
		s.messageImpl = message.NewImplementation(s.GetMessageService(ctx))
	}

	return s.messageImpl
}
