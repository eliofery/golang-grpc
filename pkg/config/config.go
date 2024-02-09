package config

import "github.com/joho/godotenv"

// Load ...
func Load(env string) error {
	if err := godotenv.Load(env); err != nil {
		return err
	}

	return nil
}
