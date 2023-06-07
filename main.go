package main

import (
	"flag"
	"fmt"
	"github.com/luizfnunesmarques/any-metric/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return ""
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var metricFlags arrayFlags

	flag.Var(&metricFlags, "metric", "Single metric to be generated.")
	flag.Parse()

	metricsConfig, err := parseMetricFlags(metricFlags)
	if err != nil {
		log.Fatalf("Error parsing metric flags: %v", err)
	}

	collectors := metrics.StartMetricUpdates(metricsConfig)

	registry := prometheus.NewRegistry()

	for _, metric := range collectors {
		registry.Register(metric)
	}

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})

	http.Handle("/metrics", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseMetricFlags(metricFlags arrayFlags) ([]*metrics.MetricConfig, error) {
	metricsConfig := []*metrics.MetricConfig{}

	for _, metricSpec := range metricFlags {
		parts := strings.Split(metricSpec, ":")
		if len(parts) != 4 {
			return nil, fmt.Errorf("invalid metric specification: %s", metricSpec)
		}

		name := parts[0]
		frequency, err := time.ParseDuration(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid metric frequency: %s", parts[1])
		}

		increment, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid metric increment: %s", parts[2])
		}

		metricType := parts[3]

		metric := &metrics.MetricConfig{
			Name:      name,
			Frequency: frequency,
			Increment: increment,
			Type:      metricType,
		}
		metricsConfig = append(metricsConfig, metric)
	}

	return metricsConfig, nil
}
