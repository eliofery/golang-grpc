package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// REST ...
type REST struct {
	mux  *runtime.ServeMux
	opts []grpc.DialOption
}

// NewREST ...
func NewREST() *REST {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	return &REST{
		mux:  mux,
		opts: opts,
	}
}

// Mux ...
func (s *REST) Mux() *runtime.ServeMux {
	return s.mux
}

// Opts ...
func (s *REST) Opts() []grpc.DialOption {
	return s.opts
}
