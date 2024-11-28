package app

import (
	"context"
	"log"
	"net"

	"github.com/marinaaaniram/go-common-platform/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"go-chat-server/internal/config"
	"go-chat-server/internal/interceptor"
	descChat_v1 "go-chat-server/pkg/chat_v1"
	descMessage_v1 "go-chat-server/pkg/message_v1"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// Create app
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

// Init deps
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// Init config
func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

// Init service provider
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

// Init GRPC server
func (a *App) initGRPCServer(ctx context.Context) error {

	authAddress := a.serviceProvider.GetGRPCConfig().AuthServiceAddress()
	interceptor.SetAuthAddress(authAddress)

	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(interceptor.AuthInterceptor),
	)

	reflection.Register(a.grpcServer)

	descChat_v1.RegisterChatV1Server(a.grpcServer, a.serviceProvider.GetChatImpl(ctx))
	descMessage_v1.RegisterMessageV1Server(a.grpcServer, a.serviceProvider.GetMessageImpl(ctx))

	return nil
}

// Run GRPC server
func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GetGRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GetGRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
