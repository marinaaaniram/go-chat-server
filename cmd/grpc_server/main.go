package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-chat-server/internal/config"
	"github.com/marinaaaniram/go-chat-server/internal/converter"
	chatRepository "github.com/marinaaaniram/go-chat-server/internal/repository/chat"
	messageRepository "github.com/marinaaaniram/go-chat-server/internal/repository/message"
	"github.com/marinaaaniram/go-chat-server/internal/service"
	chatService "github.com/marinaaaniram/go-chat-server/internal/service/chat"
	messageService "github.com/marinaaaniram/go-chat-server/internal/service/message"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedChatV1Server
	chatService    service.ChatService
	messageService service.MessageService
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	chatRepo := chatRepository.NewRepository(pool)
	chatSrv := chatService.NewService(chatRepo)

	messageRepo := messageRepository.NewRepository(pool)
	messageSrv := messageService.NewService(messageRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{chatService: chatSrv, messageService: messageSrv})

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// CreateChat - create chat
func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	chatDesc, err := s.chatService.Create(ctx, converter.FromDescCreateToChat(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateChatResponse{
		Chat: chatDesc,
	}, nil
}

// DeleteChat - delete chat by id
func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	err := s.chatService.Delete(ctx, converter.FromDescDeleteToChat(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// SendMessage - send message to chat
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	err := s.messageService.SendMessage(ctx, converter.FromDescToMessage(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
