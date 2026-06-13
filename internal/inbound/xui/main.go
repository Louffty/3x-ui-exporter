// Package inboundxui
package inboundxui

import (
	outboundxui "xui-exporter/internal/outbound/xui"

	"github.com/prometheus/client_golang/prometheus"
)

type XUICollector struct {
	client         *outboundxui.XUIClient
	inboundUp      *prometheus.Desc
	inboundDown    *prometheus.Desc
	inboundTotal   *prometheus.Desc
	inboundEnabled *prometheus.Desc
	clientUp       *prometheus.Desc
	clientDown     *prometheus.Desc
	clientTotal    *prometheus.Desc
	clientEnabled  *prometheus.Desc
	clientExpiry   *prometheus.Desc
	scrapeSuccess  *prometheus.Desc
}

func NewXUICollector(client *outboundxui.XUIClient) *XUICollector {
	return &XUICollector{
		client: client,
		inboundUp: prometheus.NewDesc(
			"xui_inbound_up_bytes",
			"Inbound uploaded traffic in bytes.",
			[]string{"id", "remark", "protocol", "port"},
			nil,
		),
		inboundDown: prometheus.NewDesc(
			"xui_inbound_down_bytes",
			"Inbound downloaded traffic in bytes.",
			[]string{"id", "remark", "protocol", "port"},
			nil,
		),
		inboundTotal: prometheus.NewDesc(
			"xui_inbound_total_bytes",
			"Inbound traffic limit in bytes. Zero usually means unlimited.",
			[]string{"id", "remark", "protocol", "port"},
			nil,
		),
		inboundEnabled: prometheus.NewDesc(
			"xui_inbound_enabled",
			"Whether inbound is enabled: 1 enabled, 0 disabled.",
			[]string{"id", "remark", "protocol", "port"},
			nil,
		),
		clientUp: prometheus.NewDesc(
			"xui_client_up_bytes",
			"Client uploaded traffic in bytes.",
			[]string{"inbound_id", "inbound_remark", "email"},
			nil,
		),
		clientDown: prometheus.NewDesc(
			"xui_client_down_bytes",
			"Client downloaded traffic in bytes.",
			[]string{"inbound_id", "inbound_remark", "email"},
			nil,
		),
		clientTotal: prometheus.NewDesc(
			"xui_client_total_bytes",
			"Client traffic limit in bytes. Zero usually means unlimited.",
			[]string{"inbound_id", "inbound_remark", "email"},
			nil,
		),
		clientEnabled: prometheus.NewDesc(
			"xui_client_enabled",
			"Whether client is enabled: 1 enabled, 0 disabled.",
			[]string{"inbound_id", "inbound_remark", "email"},
			nil,
		),
		clientExpiry: prometheus.NewDesc(
			"xui_client_expiry_timestamp_ms",
			"Client expiry timestamp in milliseconds. Zero usually means no expiry.",
			[]string{"inbound_id", "inbound_remark", "email"},
			nil,
		),
		scrapeSuccess: prometheus.NewDesc(
			"xui_scrape_success",
			"Whether the last x-ui scrape succeeded: 1 success, 0 failure.",
			nil,
			nil,
		),
	}
}
