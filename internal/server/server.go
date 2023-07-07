package server

import (
	"log"

	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/prometheus/client_golang/prometheus"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewRegistrar)

var _metricSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "server",
	Subsystem: "requests",
	Name:      "duration_ms",
	Help:      "server requests duration(ms).",
	Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
}, []string{"kind", "operation"})

var _metricRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace: "client",
	Subsystem: "requests",
	Name:      "code_total",
	Help:      "The total number of processed requests",
}, []string{"kind", "operation", "code", "reason"})

func NewRegistrar(sc []constant.ServerConfig, cc constant.ClientConfig) registry.Registrar {
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  &cc,
		},
	)
	if err != nil {
		log.Panic(err)
	}

	prometheus.MustRegister(_metricSeconds, _metricRequests)

	r := nacos.New(client)
	return r
}
