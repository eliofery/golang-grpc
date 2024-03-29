package core

import (
	"flag"
	"fmt"
)

// ConfigPath ...
type ConfigPath struct {
	EnvPath  string
	YamlPath string
}

// Migration ...
type Migration struct {
	IsMigration bool
}

// Options cli options struct
type Options struct {
	ConfigPath
	Migration
}

// Option cli options
var Option Options

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprint(flag.CommandLine.Output(), "Usage: bin/grpc-server [options]\n\n")
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "%s\n", Option.descPostgres())
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "%s\n", Option.descRedis())
		_, _ = fmt.Fprint(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&Option.ConfigPath.EnvPath, "config-env-path", ".env", "path to env config file")
	flag.StringVar(&Option.ConfigPath.YamlPath, "config-yaml-path", "config/local.yaml", "path to yaml config file")
	flag.BoolVar(&Option.Migration.IsMigration, "migration", false, "run database migration on project start")

	flag.Parse()
}

// NewCli ...
func NewCli() *Options {
	return &Option
}

// descPostgres ...
func (o *Options) descPostgres() string {
	return `Postgres server:
To start the postgres server you need to run docker compose and specify parameters in env:

  POSTGRES_USER			postgres login database
  POSTGRES_PASSWORD		postgres password database
  POSTGRES_DATABASE		postgres name database

  Example:
  POSTGRES_USER=root POSTGRES_PASSWORD=123456 POSTGRES_DATABASE=test bin/grpc-server
`
}

// descRedis ...
func (o *Options) descRedis() string {
	return `Redis server:
To start the redis server you need to run docker compose and specify parameters in env:

  REDIS_USER			redis login database
  REDIS_PASSWORD		redis password database
  REDIS_DATABASE		redis name database

  Example:
  REDIS_USER=root REDIS_PASSWORD=123456 REDIS_DATABASE=0 bin/grpc-server
`
}
