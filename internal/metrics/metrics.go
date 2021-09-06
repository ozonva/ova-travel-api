package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	createCounter      prometheus.Counter
	multiCreateCounter prometheus.Counter
	listCounter        prometheus.Counter
	describeCounter    prometheus.Counter
	updateCounter      prometheus.Counter
	removeCounter      prometheus.Counter
	totalCounter       prometheus.Counter
}

func (m Metrics) CreateTravelCounterInc() {
	m.createCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) MultiCreateTravelCounterInc() {
	m.multiCreateCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) ListTravelCounterInc() {
	m.listCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) DescribeTravelCounterInc() {
	m.describeCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) UpdateTravelCounterInc() {
	m.updateCounter.Inc()
	m.totalCounter.Inc()
}

func (m Metrics) RemoveTravelCounterInc() {
	m.removeCounter.Inc()
	m.totalCounter.Inc()
}

func NewMetrics() *Metrics {
	m := &Metrics{
		createCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "travel_api_create_total",
			Help: "The total number create requests in travel api",
		}),
		multiCreateCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "travel_api_multi_create_total",
			Help: "The total number multi create requests in travel api",
		}),
		listCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "travel_api_list_total",
			Help: "The total number list requests in travel api",
		}),
		describeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "travel_api_describe_total",
			Help: "The total number describe requests in travel api",
		}),
		updateCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "travel_api_update_total",
			Help: "The total number update requests in travel api",
		}),
		removeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "travel_api_remove_total",
			Help: "The total number remove requests in travel api",
		}),
		totalCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "travel_api_total",
			Help: "The total number requests in travel api",
		}),
	}

	return m
}
