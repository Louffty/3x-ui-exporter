// Package outboundxui
package outboundxui

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/cookiejar"

	"xui-exporter/internal/domain"
)

type XUIClient struct {
	config *domain.Config
	client *http.Client
}

func NewXUIClient(cfg *domain.Config) *XUIClient {
	if cfg == nil {
		log.Fatal("Please provide an config")
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("cannot create cookie jar: %v", err)
	}

	transport := &http.Transport{}

	if cfg.InsecureTLS {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return &XUIClient{
		config: cfg,
		client: &http.Client{
			Transport: transport,
			Timeout:   cfg.RequestTimeout,
			Jar:       jar,
		},
	}
}
