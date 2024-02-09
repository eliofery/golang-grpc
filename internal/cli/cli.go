// Package cli provides cli options
package cli

import "flag"

// Options cli options struct
type options struct {
	ConfigPath string // Path to config file
}

// Option cli options
var Option options

func init() {
	flag.StringVar(&Option.ConfigPath, "config-path", ".env", "path to config file")
}
