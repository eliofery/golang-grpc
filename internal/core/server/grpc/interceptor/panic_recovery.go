package interceptor

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"strings"

	"github.com/eliofery/golang-grpc/pkg/eslog"
	"google.golang.org/grpc"
)

// PanicRecovery ...
type PanicRecovery interface {
	Recovery() grpc.UnaryServerInterceptor
}

type panicRecovery struct {
	logger eslog.Logger
}

// NewPanicRecovery ...
func NewPanicRecovery(logger eslog.Logger) PanicRecovery {
	return &panicRecovery{
		logger: logger,
	}
}

// Recovery ...
func (p *panicRecovery) Recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		defer func() {
			if r := recover(); r != nil {
				p.logger.Error("Recovered from panic", slog.Any("err", r))

				stackStr := string(debug.Stack())
				stackStr = strings.ReplaceAll(stackStr, "\t", "  ")
				stacks := strings.Split(stackStr, "\n")

				p.logger.Error("Stack trace")
				for _, line := range stacks {
					fmt.Println(line)
				}
			}
		}()

		return handler(ctx, req)
	}
}
