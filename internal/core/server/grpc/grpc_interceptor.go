package grpc

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

// Interceptor ...
type Interceptor struct {
	fx.Out

	Interceptors []grpc.UnaryServerInterceptor
}

// NewInterceptor ...
func NewInterceptor() Interceptor {
	useInterceptors := []grpc.UnaryServerInterceptor{
		//Validate,
	}

	return Interceptor{
		Interceptors: useInterceptors,
	}
}
