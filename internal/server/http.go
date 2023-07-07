package server

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	sentrykratos "github.com/go-kratos/sentry"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	middlewareLog "github.com/star-table/go-common/pkg/middleware/log"
	middlewareMeta "github.com/star-table/go-common/pkg/middleware/meta"
	"github.com/star-table/go-table/internal/conf"
	"github.com/star-table/go-table/internal/service"
	"github.com/star-table/interface/golang/ping"
	v1 "github.com/star-table/interface/golang/table/v1"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, tableService *service.TableService, rowsService *service.RowsService, pingService *service.PingService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(recovery.WithLogger(logger)),
			validate.Validator(),
			sentrykratos.Server(),
			tracing.Server(),
			middlewareMeta.Server(),
			middlewareLog.LogServerMiddleware(logger, true),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
				metrics.WithRequests(prom.NewCounter(_metricRequests)),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/prometheus", promhttp.Handler())
	v1.RegisterTableHTTPServer(srv, tableService)
	v1.RegisterRowsHTTPServer(srv, rowsService)
	ping.RegisterPingHTTPServer(srv, pingService)
	return srv
}
