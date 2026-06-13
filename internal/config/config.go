// Package config need for parsing config
package config

import (
	"os"
	"strconv"
	"time"

	"xui-exporter/internal/domain"
)

func ReadConfig() *domain.Config {
	cfg := &domain.Config{}

	cfg.PanelURL = os.Getenv("PANEL_URL")
	cfg.Username = os.Getenv("USERNAME")
	cfg.Password = os.Getenv("PASSWORD")
	cfg.ListenAddr = getEnv("LISTEN_ADDR", "0.0.0.0:9100")

	insecureTLS := os.Getenv("INSECURE_TLS")
	insecureTLSBool, err := strconv.ParseBool(insecureTLS)
	if err != nil {
		panic(err)
	}

	cfg.InsecureTLS = insecureTLSBool

	cfg.RequestTimeout = 10 * time.Second

	return cfg
}
