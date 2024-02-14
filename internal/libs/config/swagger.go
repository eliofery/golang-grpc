package config

// Editor ...
type Editor struct {
	Port int `yaml:"port" env-default:"8085" env:"SWAGGER_EDITOR_PORT"`
}

// UI ...
type UI struct {
	Port int `yaml:"port" env-default:"8086" env:"SWAGGER_UI_PORT"`
}

// Swagger ...
type Swagger struct {
	Editor `yaml:"editor"`
	UI     `yaml:"ui"`
}
