package inboundxui

import (
	"log"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

func (c *XUICollector) Collect(ch chan<- prometheus.Metric) {
	inbounds, err := c.client.GetInbounds()
	if err != nil {
		log.Printf("scrape failed: %v", err)
		ch <- prometheus.MustNewConstMetric(
			c.scrapeSuccess,
			prometheus.GaugeValue,
			0,
		)
		return
	}
	ch <- prometheus.MustNewConstMetric(
		c.scrapeSuccess,
		prometheus.GaugeValue,
		1,
	)
	for _, inbound := range inbounds {
		inboundID := strconv.Itoa(inbound.ID)
		port := strconv.Itoa(inbound.Port)
		labels := []string{
			inboundID,
			inbound.Remark,
			inbound.Protocol,
			port,
		}
		ch <- prometheus.MustNewConstMetric(
			c.inboundUp,
			prometheus.GaugeValue,
			float64(inbound.Up),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.inboundDown,
			prometheus.GaugeValue,
			float64(inbound.Down),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.inboundTotal,
			prometheus.GaugeValue,
			float64(inbound.Total),
			labels...,
		)
		ch <- prometheus.MustNewConstMetric(
			c.inboundEnabled,
			prometheus.GaugeValue,
			boolToFloat(inbound.Enable),
			labels...,
		)
		for _, stat := range inbound.ClientStats {
			clientLabels := []string{
				inboundID,
				inbound.Remark,
				stat.Email,
			}
			ch <- prometheus.MustNewConstMetric(
				c.clientUp,
				prometheus.GaugeValue,
				float64(stat.Up),
				clientLabels...,
			)
			ch <- prometheus.MustNewConstMetric(
				c.clientDown,
				prometheus.GaugeValue,
				float64(stat.Down),
				clientLabels...,
			)
			ch <- prometheus.MustNewConstMetric(
				c.clientTotal,
				prometheus.GaugeValue,
				float64(stat.Total),
				clientLabels...,
			)
			ch <- prometheus.MustNewConstMetric(
				c.clientEnabled,
				prometheus.GaugeValue,
				boolToFloat(stat.Enable),
				clientLabels...,
			)
			ch <- prometheus.MustNewConstMetric(
				c.clientExpiry,
				prometheus.GaugeValue,
				float64(stat.ExpiryTime),
				clientLabels...,
			)
		}
	}
}
