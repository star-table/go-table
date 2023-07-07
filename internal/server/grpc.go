package server

import (
	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	sentrykratos "github.com/go-kratos/sentry"
	middlewareLog "github.com/star-table/go-common/pkg/middleware/log"
	middlewareMeta "github.com/star-table/go-common/pkg/middleware/meta"
	"github.com/star-table/go-table/internal/conf"
	"github.com/star-table/go-table/internal/service"
	v1 "github.com/star-table/interface/golang/table/v1"
	goGrpc "google.golang.org/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.TableService, rowsService *service.RowsService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
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
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	opts = append(opts, grpc.Options(goGrpc.MaxRecvMsgSize(50000000), goGrpc.MaxSendMsgSize(50000000)))
	srv := grpc.NewServer(opts...)
	v1.RegisterTableServer(srv, greeter)
	v1.RegisterRowsServer(srv, rowsService)
	return srv
}
