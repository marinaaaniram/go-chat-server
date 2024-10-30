package main

import (
	"context"
	"flag"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-chat-server/internal/config"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
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

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	serverInstance := &server{
		pool: pool,
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, serverInstance)

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// CreateChat - create chat
func (s *server) CreateChat(ctx context.Context, req *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	builderInsert := sq.Insert("chat").
		PlaceholderFormat(sq.Dollar).
		Columns("usernames").
		Values(pq.Array(req.Usernames)).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build query: %v", err)
	}

	var chatID int64
	err = s.pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert chat: %v", err)
	}

	log.Printf("Created chat with id: %d", chatID)

	return &desc.CreateChatResponse{
		Id: int64(chatID),
	}, nil
}

// DeleteChat - delete chat by id
func (s *server) DeleteChat(ctx context.Context, req *desc.DeleteChatRequest) (*emptypb.Empty, error) {
	builderSelect := sq.Select("COUNT(*)").
		From("chat").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	selectQuery, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	var count int
	err = s.pool.QueryRow(ctx, selectQuery, args...).Scan(&count)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to select chat: %v", err)
	}

	if count == 0 {
		return nil, status.Errorf(codes.NotFound, "Chat with id %d not found", req.GetId())
	}

	builderDelete := sq.Delete("chat").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build delete query: %v", err)
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete chat: %v", err)
	}

	log.Printf("Chat with id %d deleted", req.GetId())

	return &emptypb.Empty{}, nil
}

// SendMessage - send message to chat
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	builderInsert := sq.Insert("message").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "sent_by", "text").
		Values(req.GetChatId(), req.GetSentBy(), req.GetText()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build query: %v", err)
	}

	var messageId int
	err = s.pool.QueryRow(ctx, query, args...).Scan(&messageId)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23503" {
			return nil, status.Errorf(codes.InvalidArgument, "Chat with id %d not found", req.GetChatId())
		} else {
			return nil, status.Errorf(codes.Internal, "Failed to insert message: %v", err)
		}
	}

	log.Printf("Sent message with id: %d", messageId)

	return &emptypb.Empty{}, nil
}
