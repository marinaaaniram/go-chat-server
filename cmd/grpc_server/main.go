package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("Server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// Create - create chat
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Chat usernames: %s", req.GetUsernames())

	randomID, err := rand.Int(rand.Reader, big.NewInt(1<<63-1))
	if err != nil {
		return nil, err
	}

	log.Printf("Chat id: %d", randomID)

	return &desc.CreateResponse{
		Id: randomID.Int64(),
	}, nil
}

// Delete - delete chat by id
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat id: %d", req.GetId())
	return &emptypb.Empty{}, nil
}

// SendMessage - send message to chat
func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf(
		"Chat from: %s, Chat text: %s, Chat timestamp: %s",
		req.GetFrom(), req.GetText(), req.GetTimestamp())

	return &emptypb.Empty{}, nil
}
