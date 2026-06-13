package config

import (
	"os"
)

func getEnv(key, fallback string) string {
	env := os.Getenv(key)
	if env == "" {
		env = fallback
	}

	return env
}
