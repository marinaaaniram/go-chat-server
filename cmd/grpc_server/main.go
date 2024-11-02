package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chatAPI "github.com/marinaaaniram/go-chat-server/internal/api/chat"
	messageAPI "github.com/marinaaaniram/go-chat-server/internal/api/message"
	"github.com/marinaaaniram/go-chat-server/internal/config"
	chatRepository "github.com/marinaaaniram/go-chat-server/internal/repository/chat"
	messageRepository "github.com/marinaaaniram/go-chat-server/internal/repository/message"
	chatService "github.com/marinaaaniram/go-chat-server/internal/service/chat"
	messageService "github.com/marinaaaniram/go-chat-server/internal/service/message"
	descChat_v1 "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
	descMessage_v1 "github.com/marinaaaniram/go-chat-server/pkg/message_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

// type server struct {
// 	descChat_v1.UnimplementedChatV1Server
// 	descMessage_v1.UnimplementedMessageV1Server
// 	chatService    service.ChatService
// 	messageService service.MessageService
// }

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

	descChat_v1.RegisterChatV1Server(s, chatAPI.NewImplementation(chatSrv))
	descMessage_v1.RegisterMessageV1Server(s, messageAPI.NewImplementation(messageSrv))

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
