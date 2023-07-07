package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/star-table/go-table/internal/conf"
	pushPb "github.com/star-table/interface/golang/push/v1"
)

type goPushRepo struct {
	pushPb.PushHTTPClient
	log *log.Helper
}

func NewGoPushRepo(conf *conf.Data, r registry.Discovery, logger log.Logger) pushPb.PushHTTPClient {
	conn, err := getHttpConn(conf.Facade.GoPushServer, r, logger)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &goPushRepo{
		PushHTTPClient: pushPb.NewPushHTTPClient(conn),
		log:            log.NewHelper(logger),
	}
}
