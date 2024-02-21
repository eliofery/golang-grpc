package server

import (
	"go.uber.org/fx"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// Transport ...
type Transport struct {
	fx.Out

	credentials.TransportCredentials
}

// NewTransport ...
func NewTransport() Transport {
	return Transport{
		TransportCredentials: insecure.NewCredentials(),
	}
}
