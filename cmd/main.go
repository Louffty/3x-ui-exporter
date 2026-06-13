package main

import (
	"log"
	"net/http"

	"xui-exporter/internal/config"
	inboundxui "xui-exporter/internal/inbound/xui"
	outboundxui "xui-exporter/internal/outbound/xui"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.ReadConfig()

	xuiClient := outboundxui.NewXUIClient(cfg)

	if err := xuiClient.Login(); err != nil {
		log.Fatalf("Cannot login: %v", err)
	}

	collector := inboundxui.NewXUICollector(xuiClient)
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())

	log.Printf("x-ui exporter listening on %s", cfg.ListenAddr)
	if err := http.ListenAndServe(cfg.ListenAddr, nil); err != nil {
		log.Fatal(err)
	}
}
