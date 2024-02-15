package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/eliofery/golang-fullstack/internal/libs/config"
	"github.com/eliofery/golang-fullstack/pkg/eslog"
	desc "github.com/eliofery/golang-fullstack/pkg/microservice/user/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// REST ...
type REST struct {
	server *runtime.ServeMux
	config *config.Config
	logger *eslog.Logger
}

// NewREST ...
func NewREST(config *config.Config, logger *eslog.Logger) *REST {
	return &REST{
		config: config,
		logger: logger,
	}
}

// Init ...
func (r *REST) Init() error {
	r.server = runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := desc.RegisterUserV1ServiceHandlerFromEndpoint(context.Background(), r.server, r.config.Server.RESTAddress(), opts)
	if err != nil {
		return err
	}

	return nil
}

// Run ...
func (r *REST) Run() error {
	server := &http.Server{
		Addr:         r.config.Server.RESTAddress(),
		Handler:      r.server, //handler := allowCORS(mux)
		ReadTimeout:  r.config.Server.Read,
		WriteTimeout: r.config.Server.Write,
		IdleTimeout:  r.config.Server.Idle,
	}
	_ = server

	r.logger.Info("REST server start", slog.String("rest", r.config.Server.RESTAddress()))
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// allowCORS ...
//func allowCORS(handler http.Handler) http.Handler {
//    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//        w.Header().Set("Access-Control-Allow-Origin", "*")
//        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//        w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
//
//        if r.Method == "OPTIONS" {
//            w.WriteHeader(http.StatusNoContent)
//            return
//        }
//
//        handler.ServeHTTP(w, r)
//    })
//}
