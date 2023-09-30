package grpc

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
	"github.com/fbriansyah/micro-session-service/internal/port"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServerAdapter struct {
	server              *grpc.Server
	service             port.SessionServicePort
	grpcPort            int
	accessTokenDuration time.Duration
	refeshTokenDuration time.Duration

	session.UnimplementedSessionServiceServer
}

type GrpcServerConfig struct {
	GrpcPort            int
	AccessTokenDuration time.Duration
	RefeshTokenDuration time.Duration
}

// NewGrpcServerAdapter create GrpcServerAdapter server
func NewGrpcServerAdapter(service port.SessionServicePort, config GrpcServerConfig) *GrpcServerAdapter {
	return &GrpcServerAdapter{
		grpcPort:            config.GrpcPort,
		accessTokenDuration: config.AccessTokenDuration,
		refeshTokenDuration: config.RefeshTokenDuration,
		service:             service,
	}
}

// Run grpc server
func (a *GrpcServerAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf("%d", a.grpcPort))

	if err != nil {
		log.Fatalf("failed to listen on port %d: %v\n", a.grpcPort, err)
	}
	log.Printf("Server listen on port %d \n", a.grpcPort)
	grpcServer := grpc.NewServer()

	a.server = grpcServer

	session.RegisterSessionServiceServer(grpcServer, a)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server grpc on port %d: %v\n", a.grpcPort, err)
	}
}

// Stop the grpc server
func (a *GrpcServerAdapter) Stop() {
	a.server.Stop()
}

func generateError(code codes.Code, msg string) error {
	s := status.New(code, msg)
	s, _ = s.WithDetails(&errdetails.ErrorInfo{})
	return s.Err()
}
