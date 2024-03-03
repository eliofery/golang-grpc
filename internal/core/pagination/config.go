package pagination

import (
	"fmt"

	"go.uber.org/config"
)

const (
	configKeyName = "pagination"

	defaultLimit = 10
)

// Config ...
type Config struct {
	Limit uint64 `yaml:"limit"`
}

// NewConfig ...
func NewConfig(provider config.Provider) (*Config, error) {
	var conf Config
	conf.Limit = defaultLimit
	if err := provider.Get(configKeyName).Populate(&conf); err != nil {
		return nil, fmt.Errorf("error to populate pagination config: %w", err)
	}

	return &conf, nil
}
