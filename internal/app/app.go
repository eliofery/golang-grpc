package app

//
//import (
//    "context"
//    "fmt"
//    "github.com/eliofery/golang-fullstack/internal/libs/config"
//    "github.com/eliofery/golang-fullstack/pkg/database/postgres"
//    "log/slog"
//    "net"
//    "net/http"
//    "os"
//
//    "github.com/eliofery/golang-fullstack/pkg/eslog"
//    pb "github.com/eliofery/golang-fullstack/pkg/microservice/v1"
//    "github.com/golang/protobuf/ptypes/wrappers"
//    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
//    "google.golang.org/grpc"
//    "google.golang.org/grpc/credentials/insecure"
//    "google.golang.org/grpc/reflection"
//)
//
//// MicroserviceServer ...
//type MicroserviceServer struct {
//	pb.UnimplementedMicroServiceServer
//}
//
//// GetUser ...
//func (s *MicroserviceServer) GetUser(context.Context, *pb.GetUserRequest) (*pb.GetUserResponse, error) {
//	return &pb.GetUserResponse{
//		Result: &wrappers.StringValue{Value: "success"},
//	}, nil
//}
//
//// App ...
//type App struct {
//	ctx    context.Context
//	config *config.Config
//	logger *eslog.Logger
//	conn   *postgres.Postgres
//}
//
//// New ...
//func New(config *config.Config, conn *postgres.Postgres) *App {
//	return &App{
//		ctx:    context.Background(),
//		config: config,
//		logger: config.NewLogger(),
//        conn: conn,
//	}
//}
//
//// Run ...
//func Run(a *App) error {
//    fmt.Println("RUN")
//
//    return nil
//    return a.initDeps(
//		a.initPostgres,
//		a.StartServer,
//	)
//}
//
//// initDeps ...
//func (a *App) initDeps(deps ...func() error) error {
//	for _, dep := range deps {
//		if err := dep(); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//// initPostgres ...
//func (a *App) initPostgres() error {
//	fmt.Println(a.conn)
//	os.Exit(1)
//
//	if err := a.conn.Connect(a.ctx); err != nil {
//		return err
//	}
//
//	if err := a.conn.Migrate(); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// StartServer ...
//func (a *App) StartServer() error {
//	ch := make(chan error, 2)
//
//	// gRPC server
//	go func(ch chan error) {
//		listen, err := net.Listen("tcp", a.config.GRPCAddress())
//		if err != nil {
//			ch <- err
//		}
//		a.logger.Info("gRPC server start", slog.String("grpc", a.config.GRPCAddress()))
//
//		grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
//		reflection.Register(grpcServer)
//
//		pb.RegisterMicroServiceServer(grpcServer, &MicroserviceServer{})
//
//		if err = grpcServer.Serve(listen); err != nil {
//			ch <- err
//		}
//	}(ch)
//
//	// REST server
//	go func() {
//		mux := runtime.NewServeMux()
//		opts := []grpc.DialOption{
//			grpc.WithTransportCredentials(insecure.NewCredentials()),
//		}
//
//		err := pb.RegisterMicroServiceHandlerFromEndpoint(a.ctx, mux, a.config.GRPCAddress(), opts)
//		if err != nil {
//			ch <- err
//		}
//
//		handler := allowCORS(mux)
//		server := &http.Server{
//			Addr:         a.config.RESTAddress(),
//			Handler:      handler,
//			ReadTimeout:  a.config.Timeout.Read,
//			WriteTimeout: a.config.Timeout.Write,
//			IdleTimeout:  a.config.Timeout.Idle,
//		}
//		_ = server
//
//		a.logger.Info("REST server start", slog.String("rest", a.config.RESTAddress()))
//		err = server.ListenAndServe()
//		if err != nil {
//			ch <- err
//		}
//	}()
//
//	for i := 0; i < 2; i++ {
//		if err := <-ch; err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//
