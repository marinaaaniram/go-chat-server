package app

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/marinaaaniram/go-common-platform/pkg/closer"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/marinaaaniram/go-chat-server/internal/config"
	"github.com/marinaaaniram/go-chat-server/internal/interceptor"

	"github.com/marinaaaniram/go-chat-server/internal/logger"
	"github.com/marinaaaniram/go-chat-server/internal/tracing"
	descChat_v1 "github.com/marinaaaniram/go-chat-server/pkg/chat_v1"
)

var logLevel = flag.String("l", "info", "log level")

const (
	serviceName = "github.com/marinaaaniram/go-chat-server"
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

	logger.Init(getCore(getAtomicLevel()))
	tracing.Init(logger.Logger(), serviceName)

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
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				// otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
				// interceptor.ServerTracingInterceptor,
				interceptor.AuthInterceptor,
			)),
	)

	reflection.Register(a.grpcServer)

	descChat_v1.RegisterChatV1Server(a.grpcServer, a.serviceProvider.GetChatImpl(ctx))
	// descMessage_v1.RegisterMessageV1Server(a.grpcServer, a.serviceProvider.GetMessageImpl(ctx))

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

func getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     7, // days
	})

	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
	)
}

func getAtomicLevel() zap.AtomicLevel {
	var level zapcore.Level
	if err := level.Set(*logLevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
