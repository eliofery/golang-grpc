package config

// Adminer ...
type Adminer struct {
	Port int `yaml:"port" env-default:"3333" env:"ADMINER_PORT"`
}
