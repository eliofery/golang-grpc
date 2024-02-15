package server

import (
	"log/slog"
	"net"

	userServiceV1 "github.com/eliofery/golang-fullstack/internal/app/api/microservice/user/v1"
	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	desc "github.com/eliofery/golang-fullstack/pkg/microservice/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// GRPC ...
type GRPC struct {
	server *grpc.Server
	config *config.Config
	logger *eslog.Logger
}

// NewGRPC ...
func NewGRPC(config *config.Config, logger *eslog.Logger) *GRPC {
	return &GRPC{
		config: config,
		logger: logger,
	}
}

// Init ...
func (g *GRPC) Init() {
	g.server = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(g.server)

	desc.RegisterUserV1ServiceServer(g.server, userServiceV1.New(
	// Services List
	))
}

// Run ...
func (g *GRPC) Run() error {
	list, err := net.Listen("tcp", g.config.GRPCAddress())
	if err != nil {
		return err
	}
	g.logger.Info("GRPC server is running on", slog.String("address", g.config.GRPCAddress()))

	if err = g.server.Serve(list); err != nil {
		return err
	}

	return nil
}
