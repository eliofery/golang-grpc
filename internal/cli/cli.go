// Package cli provides cli options
package cli

import (
	"flag"
	"fmt"
	"os"
)

// Config ...
type Config struct {
	EnvPath  string
	YamlPath string
}

// Options cli options struct
type Options struct {
	Config
}

// Option cli options
var Option Options

func init() {
	config := flag.NewFlagSet("config options", flag.ContinueOnError)
	config.StringVar(&Option.Config.EnvPath, "config-env-path", ".env", "path to env config file")
	config.StringVar(&Option.Config.YamlPath, "config-yaml-path", "config/local.yaml", "path to yaml config file")
	configUsage := config.Usage

	//example := flag.NewFlagSet("example options", flag.ContinueOnError)
	//example.StringVar(&Option.Config.EnvPath, "example", "", "example one")
	//example.StringVar(&Option.Config.EnvPath, "example-2", "", "example two")
	//exampleUsage := example.Usage

	config.Usage = func() {
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "GRPC/REST server:")
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "  To start the server you need to specify parameters in env:")
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "  POSTGRES_USER\t\t\tpostgres user name")
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "  POSTGRES_PASSWORD\t\tpostgres user password")
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "  POSTGRES_DATABASE\t\tpostgres name database")
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "  Example:")
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "  POSTGRES_USER=root POSTGRES_PASSWORD=123456 POSTGRES_DATABASE=test bin/grpc-server")

		optionsDisplay(configUsage)
		//optionsDisplay(exampleUsage)
	}

	_ = config.Parse(os.Args[1:])

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		os.Exit(0)
	}
}

// optionsDisplay ...
func optionsDisplay(usage func()) {
	_, _ = fmt.Fprintln(flag.CommandLine.Output(), "")
	usage()
}
