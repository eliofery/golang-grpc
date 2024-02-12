package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/eliofery/golang-fullstack/internal/cli"
	"github.com/eliofery/golang-fullstack/internal/config"
	"github.com/eliofery/golang-fullstack/pkg/database/postgres"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	"github.com/eliofery/golang-fullstack/pkg/eslog/pretty"
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

func main() {
	cfg := config.MustLoad(cli.Option)

	logger := eslog.New(pretty.NewHandler(os.Stdout, &pretty.HandlerOptions{
		SlogOptions: &slog.HandlerOptions{
			Level:     cfg.LevelVar(),
			AddSource: true,
		},
		JSON: false,
	}), cfg.LevelVar())

	logger.Debug("Debug message", slog.Int("price", 12345), slog.Any("name", "John"))
	logger.Info("Info message")
	logger.Warn("Warn message")
	logger.Error("Error message")
	//logger.Fatal("Fatal message")

	db := postgres.New(cfg, logger)
	if err := db.Connect(context.Background()); err != nil {
		logger.Fatal("failed to connect database", slog.Any("err", err))
	}

	if err := db.Migrate(); err != nil {
		logger.Fatal("failed to migrate", slog.Any("err", err))
	}

	ch := make(chan error, 2)

	// gRPC server
	go func(ch chan error) {
		listen, err := net.Listen("tcp", cfg.GRPCAddress())
		if err != nil {
			ch <- err
		}
		logger.Info("gRPC server start", slog.String("grpc", cfg.GRPCAddress()))

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

		err := pb.RegisterMicroServiceHandlerFromEndpoint(context.Background(), mux, cfg.GRPCAddress(), opts)
		if err != nil {
			ch <- err
		}

		handler := allowCORS(mux)
		server := &http.Server{
			Addr:         cfg.RESTAddress(),
			Handler:      handler,
			ReadTimeout:  cfg.Timeout.Read,
			WriteTimeout: cfg.Timeout.Write,
			IdleTimeout:  cfg.Timeout.Idle,
		}
		_ = server

		logger.Info("REST server start", slog.String("rest", cfg.RESTAddress()))
		err = server.ListenAndServe()
		if err != nil {
			ch <- err
		}
	}()

	for i := 0; i < 2; i++ {
		if err := <-ch; err != nil {
			logger.Fatal("failed to start server", slog.Any("err", err))
		}
	}
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
