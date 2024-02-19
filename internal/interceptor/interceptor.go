package interceptor

import "google.golang.org/grpc"

// ChainUnaryInterceptors ...
func ChainUnaryInterceptors() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		Validate,
	)
}
