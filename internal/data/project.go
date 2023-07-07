package data

import (
	"context"

	msgPb "github.com/star-table/interface/golang/msg/v1"

	"github.com/star-table/go-table/internal/data/facade/vo"
	"github.com/star-table/go-table/internal/data/facade/vo/projectvo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/star-table/go-table/internal/biz"
	"github.com/star-table/go-table/internal/conf"
)

type projectRepo struct {
	uc  *http.Client
	log *log.Helper
}

func NewProjectRepo(conf *conf.Data, r registry.Discovery, logger log.Logger) biz.ProjectRepo {
	conn, err := getHttpConn(conf.Facade.ProjectServer, r, logger)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &projectRepo{
		uc:  conn,
		log: log.NewHelper(logger),
	}
}

func (u *projectRepo) ReportTableEvent(ctx context.Context, eventType msgPb.EventType, traceId string, req *projectvo.TableEvent) (*vo.CommonRespVo, error) {
	respVo := &vo.CommonRespVo{}
	query := map[string]interface{}{
		"eventType": int32(eventType),
	}
	if len(traceId) > 0 {
		query["traceId"] = traceId
	}
	path := "/api/projectsvc/reportTableEvent" + convertToQueryParams(query)
	err := u.uc.Invoke(ctx, "POST", path, req, respVo)
	if err != nil {
		return nil, err
	}
	return respVo, nil
}
