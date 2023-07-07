package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-table/internal/biz"
	"github.com/star-table/go-table/internal/conf"
	"github.com/star-table/go-table/internal/data/facade/vo/permissionvo"
	comomPb "github.com/star-table/interface/golang/common/v1"
)

type permissionRepo struct {
	permissionClient *http.Client
	log              *log.Helper
}

func NewPermissionRepo(conf *conf.Data, r registry.Discovery, logger log.Logger) biz.PermissionRepo {
	conn, err := getHttpConn(conf.Facade.PermissionServer, r, logger)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &permissionRepo{
		permissionClient: conn,
		log:              log.NewHelper(logger),
	}
}

func (p *permissionRepo) GetAppRoleList(ctx context.Context, orgId, appId int64) (*permissionvo.GetAppRoleListResp, error) {
	respVo := &permissionvo.GetAppRoleListResp{}

	path := "/permission/inner/api/v1/app-permission/getAppPermissionGroupList"
	queryParams := map[string]interface{}{}
	queryParams["orgId"] = orgId
	queryParams["appId"] = appId
	err := p.permissionClient.Invoke(ctx, "POST", path, queryParams, &respVo)
	if err != nil {
		return nil, err
	}
	if respVo.Failure() {
		return nil, errors.Wrapf(comomPb.ErrorResourceNotExist("execute error"), "[GetAppRoleList] get app failed, req:%v, error:%v", queryParams, respVo.Error())
	}

	return respVo, nil
}

func (p *permissionRepo) InitAppPermissionFieldAuthCreateTable(ctx context.Context, req *permissionvo.InitAppPermissionFieldAuthCreateTableReq) error {
	respVo := &permissionvo.InitAppPermissionFieldAuthCreateTableResp{}
	path := "/permission/inner/api/v1/app-permission/initAppPermissionFieldAuthCreateTable"
	err := p.permissionClient.Invoke(ctx, "POST", path, req, &respVo)
	if err != nil {
		return err
	}
	if respVo.Failure() {
		return errors.Wrapf(comomPb.ErrorServerInternal("execute error"), "[InitAppPermissionFieldAuthCreateTable] failed, req:%v, error:%v", req, respVo.Error())
	}

	return nil
}

func (p *permissionRepo) InitAppPermissionFieldAuthDeleteTable(ctx context.Context, req *permissionvo.InitAppPermissionFieldAuthDeleteTableReq) error {
	respVo := &permissionvo.InitAppPermissionFieldAuthDeleteTableResp{}
	path := "/permission/inner/api/v1/app-permission/initAppPermissionFieldAuthDeleteTable"
	err := p.permissionClient.Invoke(ctx, "POST", path, req, &respVo)
	if err != nil {
		return err
	}
	if respVo.Failure() {
		return errors.Wrapf(comomPb.ErrorServerInternal("execute error"), "[InitAppPermissionFieldAuthDeleteTable] failed, req:%v, error:%v", req, respVo.Error())
	}

	return nil
}
