package inboundxui

import "github.com/prometheus/client_golang/prometheus"

func (c *XUICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.inboundUp
	ch <- c.inboundDown
	ch <- c.inboundTotal
	ch <- c.inboundEnabled
	ch <- c.clientUp
	ch <- c.clientDown
	ch <- c.clientTotal
	ch <- c.clientEnabled
	ch <- c.clientExpiry
	ch <- c.scrapeSuccess
}
