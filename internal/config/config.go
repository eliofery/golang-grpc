package config

// ServerConfig ...
type ServerConfig interface {
	Address() string
}

// DatabaseConfig ...
type DatabaseConfig interface {
	DSN() string
}
