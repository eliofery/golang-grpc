package env

import (
	"net"
	"os"

	"github.com/eliofery/golang-fullstack/internal/config"
)

const (
	// GrpcEnv ...
	GrpcEnv = "GRPC"
	// RestEnv ...
	RestEnv = "REST"

	grpcHostEnv = "GRPC_HOST"
	grpcPortEnv = "GRPC_PORT"

	restHostEnv = "REST_HOST"
	restPortEnv = "REST_PORT"

	hostDefault     = "localhost"
	grpcPortDefault = "50051"
	restPortDefault = "8080"
)

type serverConfig struct {
	hostEnv     string
	portEnv     string
	portDefault string
}

// NewServerConfig ...
func NewServerConfig(confEnv string) config.ServerConfig {
	var conf serverConfig

	switch confEnv {
	case GrpcEnv:
		conf = serverConfig{
			hostEnv:     os.Getenv(grpcHostEnv),
			portEnv:     os.Getenv(grpcPortEnv),
			portDefault: grpcPortDefault,
		}
	case RestEnv:
		conf = serverConfig{
			hostEnv:     os.Getenv(restHostEnv),
			portEnv:     os.Getenv(restPortEnv),
			portDefault: restPortDefault,
		}
	}

	if len(conf.hostEnv) == 0 {
		conf.hostEnv = hostDefault
	}

	if len(conf.portEnv) == 0 {
		conf.portEnv = conf.portDefault
	}

	return &conf
}

func (cfg *serverConfig) Address() string {
	return net.JoinHostPort(cfg.hostEnv, cfg.portEnv)
}
