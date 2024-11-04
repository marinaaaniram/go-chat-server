package app

import (
	"context"
	"log"

	"github.com/marinaaaniram/go-chat-server/internal/api/chat"
	"github.com/marinaaaniram/go-chat-server/internal/api/message"
	"github.com/marinaaaniram/go-chat-server/internal/client/db"
	"github.com/marinaaaniram/go-chat-server/internal/client/db/pg"
	"github.com/marinaaaniram/go-chat-server/internal/client/db/transaction"
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

	dbClient          db.Client
	txManager         db.TxManager
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

// Get postgres config
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

// Get GRPC config
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

// Init db client
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.GetPGConfig().DSN())
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

// Init db transactions manager
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// Init Chat repository
func (s *serviceProvider) GetChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

// Init Chat service
func (s *serviceProvider) GetChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.GetChatRepository(ctx))
	}

	return s.chatService
}

// Init Chat implementaion
func (s *serviceProvider) GetChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.GetChatService(ctx))
	}

	return s.chatImpl
}

// Init Message repository
func (s *serviceProvider) GetMessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = messageRepository.NewRepository(s.DBClient(ctx))
	}

	return s.messageRepository
}

// Init Message service
func (s *serviceProvider) GetMessageService(ctx context.Context) service.MessageService {
	if s.messageService == nil {
		s.messageService = messageService.NewService(s.GetMessageRepository(ctx))
	}

	return s.messageService
}

// Init Message implementaion
func (s *serviceProvider) GetMessageImpl(ctx context.Context) *message.Implementation {
	if s.messageImpl == nil {
		s.messageImpl = message.NewImplementation(s.GetMessageService(ctx))
	}

	return s.messageImpl
}
