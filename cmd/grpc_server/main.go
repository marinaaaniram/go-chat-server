package main

import (
	"context"
	"fmt"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jackc/pgx/v4/pgxpool"
	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

const (
	grpcPort = 50051
	dbDSN    = "host=localhost port=54321 dbname=postgres user=postgres password=postgres sslmode=disable"
)

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	serverInstance := &server{
		pool: pool,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
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

// Create - create chat
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
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

	return &desc.CreateResponse{
		Id: int64(chatID),
	}, nil
}

// Delete - delete chat by id
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
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
		Columns("text", "sent_by").
		Values(req.GetText(), req.GetSentBy()).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build query: %v", err)
	}

	var messageId int
	err = s.pool.QueryRow(ctx, query, args...).Scan(&messageId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert message: %v", err)
	}

	log.Printf("Sent message with id: %d", messageId)

	return &emptypb.Empty{}, nil
}
