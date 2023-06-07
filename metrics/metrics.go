package metrics

import (
	"fmt"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// MetricConfig holds the information about a specific metric.
type MetricConfig struct {
	Name      string
	Frequency time.Duration
	Increment float64
	Type      string
}

func StartMetricUpdates(metricsConfig []*MetricConfig) []prometheus.Collector {
	log.Printf("Setting any metric engine.")

	var collectors []prometheus.Collector

	for _, metricConfig := range metricsConfig {
		switch metricConfig.Type {
		case "gauge":
			gauge := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: metricConfig.Name,
				Help: fmt.Sprintf("Metric name: %s", metricConfig.Name),
			})

			increment := func() { gauge.Add(metricConfig.Increment) }
			go runIncrement(increment, *metricConfig)

			collectors = append(collectors, gauge)
		case "counter":
			counter := prometheus.NewCounter(prometheus.CounterOpts{
				Name: metricConfig.Name,
				Help: fmt.Sprintf("Metric name: %s", metricConfig.Name),
			})

			increment := func() { counter.Add(metricConfig.Increment) }
			go runIncrement(increment, *metricConfig)

			collectors = append(collectors, counter)
		default:
			log.Printf("Metric %s not initialized: %s type not supported.", metricConfig.Name, metricConfig.Type)
		}
	}

	return collectors
}

func runIncrement(fn func(), config MetricConfig) {
	ticker := time.NewTicker(config.Frequency)

	defer ticker.Stop()

	for range ticker.C {
		log.Printf("%s updated", config.Name)
		fn()
	}
}
