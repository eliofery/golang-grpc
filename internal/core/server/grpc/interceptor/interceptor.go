package interceptor

import (
	"google.golang.org/grpc"
)

// New ...
func New(
	panicRecovery PanicRecovery,
	validator Validator,
	auth Auth,
) []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		panicRecovery.Recovery(),
		validator.Validate(),
		auth.Authorize(),
	}
}
