// Package cli provides cli options
package cli

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	cliDirPath      = "docs/cli"
	fileDescription = "description.txt"
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
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "%s\n", Option.getText(fileDescription))
		_, _ = fmt.Fprint(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&Option.Config.EnvPath, "config-env-path", ".env", "path to env config file")
	flag.StringVar(&Option.Config.YamlPath, "config-yaml-path", "config/local.yaml", "path to yaml config file")
	flag.BoolVar(&Option.Migration.IsMigration, "migration", false, "run database migration on project start")

	flag.Parse()
}

// getText ...
func (o *Options) getText(fileName string) string {
	fullPath := filepath.Clean(filepath.Join(cliDirPath, fileName))
	if !strings.HasPrefix(fullPath, cliDirPath) {
		log.Printf("file path is not allowed: %s", fullPath)
		return ""
	}

	file, err := os.Open(fullPath)
	if err != nil {
		log.Printf("failed to open file: %v", err)
		return ""
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}(file)

	var buf bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buf.Write(scanner.Bytes())
		buf.WriteByte('\n')
	}
	if err = scanner.Err(); err != nil {
		log.Printf("failed to scan file: %v", err)
		return ""
	}

	return buf.String()
}
