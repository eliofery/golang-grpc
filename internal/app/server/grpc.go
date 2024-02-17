package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPC ...
type GRPC struct {
	grpc *grpc.Server
}

// NewGRPC ...
func NewGRPC() *GRPC {
	server := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	return &GRPC{
		grpc: server,
	}
}

// GRPC ...
func (s *GRPC) GRPC() *grpc.Server {
	return s.grpc
}
