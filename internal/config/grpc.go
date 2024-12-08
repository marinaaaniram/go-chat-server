package config

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	grpcHostEnvName     = "GRPC_HOST"
	grpcPortEnvName     = "GRPC_PORT"
	grpcAuthHostEnvName = "GRPC_AUTH_HOST"
	grpcAuthPortEnvName = "GRPC_AUTH_PORT"
)

type GRPCConfig interface {
	Address() string
	AuthServiceAddress() string
}

type grpcConfig struct {
	host      string
	port      string
	auth_host string
	auth_port string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	auth_host := os.Getenv(grpcAuthHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc auth service host not found")
	}

	auth_port := os.Getenv(grpcAuthPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc auth service port not found")
	}

	return &grpcConfig{
		host:      host,
		port:      port,
		auth_host: auth_host,
		auth_port: auth_port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *grpcConfig) AuthServiceAddress() string {
	return net.JoinHostPort(cfg.auth_host, cfg.auth_port)
}
