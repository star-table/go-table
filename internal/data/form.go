package data

import (
	"context"

	"github.com/star-table/go-common/pkg/errors"
	comomPb "github.com/star-table/interface/golang/common/v1"

	"github.com/star-table/go-table/internal/data/facade/vo/form"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/star-table/go-table/internal/conf"
)

type formRepo struct {
	dc  *http.Client
	log *log.Helper
}

func NewFormRepo(conf *conf.Data, r registry.Discovery, logger log.Logger) *formRepo {
	conn, err := getHttpConn(conf.Facade.FormServer, r, logger)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &formRepo{
		dc:  conn,
		log: log.NewHelper(logger),
	}
}

func (d *formRepo) QuerySql(ctx context.Context, req *form.QuerySqlReq) (*form.QuerySqlResp, error) {
	respVo := form.QuerySqlResp{}
	path := "/form/inner/api/v1/querySql"
	err := d.dc.Invoke(ctx, "POST", path, req, &respVo)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if respVo.Failure() {
		return nil, errors.Wrapf(comomPb.ErrorServerInternal("QuerySql error"), "[QuerySql] req:%v, error:%v", req, respVo.Error())
	}

	return &respVo, nil
}
