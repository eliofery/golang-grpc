package app

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	"github.com/eliofery/golang-fullstack/internal/config"
	"github.com/eliofery/golang-fullstack/pkg/database/postgres"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	pb "github.com/eliofery/golang-fullstack/pkg/microservice/v1"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// MicroserviceServer ...
type MicroserviceServer struct {
	pb.UnimplementedMicroServiceServer
}

// GetUser ...
func (s *MicroserviceServer) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Result: &wrappers.StringValue{Value: "success"},
	}, nil
}

// App ...
type App struct {
	config *config.Config
	logger *eslog.Logger
}

// New ...
func New(cfg *config.Config) *App {
	return &App{
		config: cfg,
		logger: cfg.NewLogger(),
	}
}

// Run ...
func Run(a *App) error {
	ctx := context.Background()

	if err := a.PostgresInit(ctx); err != nil {
		return err
	}

	if err := a.StartServer(ctx); err != nil {
		return err
	}

	return nil
}

// PostgresInit ...
func (a *App) PostgresInit(ctx context.Context) error {
	conn := postgres.New(a.config, a.logger)

	if err := conn.Connect(ctx); err != nil {
		return err
	}

	if err := conn.Migrate(); err != nil {
		return err
	}

	return nil
}

// StartServer ...
func (a *App) StartServer(ctx context.Context) error {
	ch := make(chan error, 2)

	// gRPC server
	go func(ch chan error) {
		listen, err := net.Listen("tcp", a.config.GRPCAddress())
		if err != nil {
			ch <- err
		}
		a.logger.Info("gRPC server start", slog.String("grpc", a.config.GRPCAddress()))

		grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
		reflection.Register(grpcServer)

		pb.RegisterMicroServiceServer(grpcServer, &MicroserviceServer{})

		if err = grpcServer.Serve(listen); err != nil {
			ch <- err
		}
	}(ch)

	// REST server
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		err := pb.RegisterMicroServiceHandlerFromEndpoint(ctx, mux, a.config.GRPCAddress(), opts)
		if err != nil {
			ch <- err
		}

		handler := allowCORS(mux)
		server := &http.Server{
			Addr:         a.config.RESTAddress(),
			Handler:      handler,
			ReadTimeout:  a.config.Timeout.Read,
			WriteTimeout: a.config.Timeout.Write,
			IdleTimeout:  a.config.Timeout.Idle,
		}
		_ = server

		a.logger.Info("REST server start", slog.String("rest", a.config.RESTAddress()))
		err = server.ListenAndServe()
		if err != nil {
			ch <- err
		}
	}()

	for i := 0; i < 2; i++ {
		if err := <-ch; err != nil {
			return err
		}
	}

	return nil
}

// allowCORS ...
func allowCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
