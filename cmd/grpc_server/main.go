package main

import (
	"context"
	"flag"
	l "log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/eliofery/golang-fullstack/internal/cli"
	"github.com/eliofery/golang-fullstack/internal/config/env"
	"github.com/eliofery/golang-fullstack/pkg/config"
	"github.com/eliofery/golang-fullstack/pkg/database/postgres"
	"github.com/eliofery/golang-fullstack/pkg/log"
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
	flag.Parse()

	ctx := context.Background()

	if err := config.Load(cli.Option.ConfigPath); err != nil {
		l.Fatalf("failed to load config: %v", err)
	}

	loggerConfig := env.NewLoggerConfig()
	logger := log.New(loggerConfig)

	pgConfig, err := env.NewPostgresConfig()
	if err != nil {
		logger.Fatal("failed to get postgres config", slog.String("err", err.Error()))
	}

	db := postgres.New()
	if err = db.Connect(ctx, pgConfig.DSN()); err != nil {
		logger.Fatal("failed to connect database", slog.String("err", err.Error()))
	}

	ch := make(chan error, 2)

	// gRPC server
	go func(ch chan error) {
		grpcConfig := env.NewServerConfig(env.GrpcEnv)
		listen, err := net.Listen("tcp", grpcConfig.Address())
		if err != nil {
			ch <- err
		}
		logger.Info("gRPC server start", slog.String("grpc", grpcConfig.Address()))

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

		err := pb.RegisterMicroServiceHandlerFromEndpoint(context.Background(), mux, ":50051", opts)
		if err != nil {
			ch <- err
		}

		readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
		if err != nil {
			ch <- err
		}

		writeTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
		if err != nil {
			ch <- err
		}

		handler := allowCORS(mux)
		restConfig := env.NewServerConfig(env.RestEnv)
		server := &http.Server{
			Addr:         restConfig.Address(),
			Handler:      handler,
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
		}
		_ = server

		logger.Info("REST server start", slog.String("rest", restConfig.Address()))
		err = server.ListenAndServe()
		if err != nil {
			ch <- err
		}
	}()

	for i := 0; i < 2; i++ {
		if err = <-ch; err != nil {
			logger.Fatal("failed to start server", slog.String("err", err.Error()))
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
