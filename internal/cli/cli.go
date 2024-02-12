// Package cli provides cli options
package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	pathDescription = "internal/cli/description.txt"
)

// Config ...
type Config struct {
	EnvPath  string
	YamlPath string
}

// Migration ...
type Migration struct {
	IsMigration bool
}

// Options cli options struct
type Options struct {
	Config
	Migration
}

// Option cli options
var Option Options

func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "%s\n", Option.descriptionText())
		_, _ = fmt.Fprint(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&Option.Config.EnvPath, "config-env-path", ".env", "path to env config file")
	flag.StringVar(&Option.Config.YamlPath, "config-yaml-path", "config/local.yaml", "path to yaml config file")
	flag.BoolVar(&Option.Migration.IsMigration, "migration", false, "run database migration on project start")

	flag.Parse()
}

func (o *Options) descriptionText() string {
	file, _ := os.OpenFile(pathDescription, os.O_RDONLY, 0644)
	defer file.Close()

	b, _ := io.ReadAll(file)

	return string(b)
}
