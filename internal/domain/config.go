package domain

import "time"

type Config struct {
	PanelURL       string
	Username       string
	Password       string
	ListenAddr     string
	InsecureTLS    bool
	RequestTimeout time.Duration
}
