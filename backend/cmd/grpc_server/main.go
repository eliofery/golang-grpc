package main

import (
	"context"
	"fmt"
	pb "github.com/eliofery/golang-fullstack/pkg/microservice/v1"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

type MicroserviceServer struct {
	pb.UnimplementedMicroServiceServer
}

func (s *MicroserviceServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	return &pb.GetUserByIdResponse{
		Result: &wrappers.StringValue{Value: "success"},
	}, nil
}

func main() {
	ch := make(chan error, 2)

	// gRPC server
	go func(ch chan error) {
		fmt.Println("gRPC server start :50051")
		listen, err := net.Listen("tcp", ":50051")
		if err != nil {
			ch <- err
		}

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

		handler := allowCORS(mux)

		fmt.Println("REST server start :8080")
		if err = http.ListenAndServe(":8080", handler); err != nil {
			ch <- err
		}
	}()

	for i := 0; i < 2; i++ {
		if err := <-ch; err != nil {
			panic(err)
		}
	}
}

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
