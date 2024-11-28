package interceptor

import (
	"context"
	"go-auth/pkg/access_v1"

	"go-chat-server/internal/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var authAddress string

// Set go-auth service address from .env
func SetAuthAddress(address string) {
	authAddress = address
}

// Check auth in go-auth service
func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	endpoint := info.FullMethod

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.ErrMetedataNotProvided
	}

	tokens := md["authorization"]
	if len(tokens) == 0 {
		return nil, errors.ErrMissingAccessToken
	}
	token := tokens[0]

	newMd := metadata.Pairs("endpoint", endpoint, "authorization", token)
	ctx = metadata.NewOutgoingContext(ctx, newMd)

	conn, err := grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		return nil, errors.ErrFailedConnectToService(err)
	}
	defer conn.Close()

	authClient := access_v1.NewAccessV1Client(conn)

	_, err = authClient.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: endpoint,
	})
	if err != nil {
		return nil, errors.ErrAccessDenied
	}

	return handler(ctx, req)
}
