package prometheus

import (
	"github.com/fahmyabdul/golibs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/rintik-io/rintik-auth/app"
)

type Prometheus struct{}

var (
	// AppStatusGauge :
	AppStatusGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace:   "rintikauth",
			Subsystem:   "app",
			Name:        "status",
			Help:        "Rintik-Auth - App Status",
			ConstLabels: prometheus.Labels(map[string]string{"version": app.CurrentVersion}),
		},
	)

	// KafkaStatusGauge :
	KafkaStatusGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "rintikauth",
			Subsystem: "kafka",
			Name:      "status",
			Help:      "Rintik-Auth - Kafka Status",
		},
	)

	// EndpointHitsCounter :
	EndpointHitsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "rintikauth",
			Subsystem: "endpoint",
			Name:      "hits",
			Help:      "The total number of endpoints hitted",
		},
		[]string{
			"url",
		},
	)
)

// Start : Starting Prometheus
func (p Prometheus) Start() error {
	golibs.Log.Println("| Prometheus | Initialize")

	err := prometheus.Register(collectors.NewBuildInfoCollector())
	if err != nil {
		return err
	}

	err = prometheus.Register(AppStatusGauge)
	if err != nil {
		return err
	}

	err = prometheus.Register(KafkaStatusGauge)
	if err != nil {
		return err
	}

	err = prometheus.Register(EndpointHitsCounter)
	if err != nil {
		return err
	}

	golibs.Log.Println("| Prometheus | Initialize Done")

	return nil
}
