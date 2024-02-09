package main

import (
	"context"
	"flag"
	"github.com/eliofery/golang-fullstack/pkg/database"
	"github.com/eliofery/golang-fullstack/pkg/database/postgres"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/eliofery/golang-fullstack/internal/cli"
	"github.com/eliofery/golang-fullstack/internal/config/env"
	"github.com/eliofery/golang-fullstack/pkg/config"
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
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig := env.NewServerConfig(env.GrpcEnv)
	restConfig := env.NewServerConfig(env.RestEnv)
	pgConfig, err := env.NewPostgresConfig()
	if err != nil {
		log.Printf("failed to get postgres config: %v", err)
	}

	if err = database.Connect(ctx, postgres.New(pgConfig)); err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	ch := make(chan error, 2)

	// gRPC server
	go func(ch chan error) {
		listen, err := net.Listen("tcp", grpcConfig.Address())
		if err != nil {
			ch <- err
		}
		log.Printf("gRPC server start %s", grpcConfig.Address())

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
		server := &http.Server{
			Addr:         restConfig.Address(),
			Handler:      handler,
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
		}

		log.Printf("REST server start %s", restConfig.Address())
		if err = server.ListenAndServe(); err != nil {
			ch <- err
		}
	}()

	for i := 0; i < 2; i++ {
		if err := <-ch; err != nil {
			panic(err)
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
