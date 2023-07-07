package data

import (
	"context"
	"fmt"

	"github.com/spf13/cast"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/star-table/go-table/internal/conf"

	"github.com/star-table/go-table/internal/data/consts"

	comomPb "github.com/star-table/interface/golang/common/v1"

	"github.com/star-table/go-common/pkg/errors"

	"github.com/star-table/go-table/internal/biz"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/star-table/go-table/internal/data/facade/vo/appvo"
)

type appRepo struct {
	cc  *http.Client
	log *log.Helper
}

func NewAppRepo(conf *conf.Data, r registry.Discovery, logger log.Logger) biz.AppRepo {
	conn, err := getHttpConn(conf.Facade.AppServer, r, logger)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &appRepo{
		cc:  conn,
		log: log.NewHelper(logger),
	}
}

func (a *appRepo) GetAppInfoByAppId(ctx context.Context, req *appvo.GetAppInfoByAppIdReq) (*appvo.GetAppInfoByAppIdResp, error) {
	respVo := &appvo.GetAppInfoByAppIdResp{}
	path := fmt.Sprintf("/app/inner/api/v1/apps/get-app-info?orgId=%v&appId=%v", req.OrgId, req.AppId)
	err := a.cc.Invoke(ctx, "GET", path, req, respVo)
	if err != nil {
		return nil, err
	}

	if respVo.Failure() {
		return nil, errors.Wrapf(comomPb.ErrorResourceNotExist("app not exist"), "[GetAppInfoByAppId] get app failed, req:%v, error:%v", req, respVo.Error())
	}

	return respVo, nil
}

// GetRealAppInfoByAppId 如果发现是镜像app，则继续查找最终的app
func (a *appRepo) GetRealAppInfoByAppId(ctx context.Context, req *appvo.GetAppInfoByAppIdReq) (*appvo.GetAppInfoByAppIdResp, error) {
	resp, err := a.GetAppInfoByAppId(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Data.Type == consts.MirrorApp {
		resp, err = a.GetAppInfoByAppId(ctx, &appvo.GetAppInfoByAppIdReq{AppId: resp.Data.MirrorAppId, OrgId: req.OrgId})
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (a *appRepo) GetSummeryAppId(ctx context.Context, orgId int64) (int64, error) {
	appList, err := a.GetAppList(ctx, &appvo.AppListReqVo{
		OrgId: orgId,
		Type:  int64(consts.LcAppTypeForSummaryTable),
	})
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return cast.ToInt64(appList.Data[0].Id), nil
}

// 获取应用列表
func (a *appRepo) GetAppList(ctx context.Context, req *appvo.AppListReqVo) (*appvo.AppListRespVo, error) {
	respVo := &appvo.AppListRespVo{}
	path := fmt.Sprintf("/app/inner/api/v1/apps/get-app-list?orgId=%v&type=%v&parentId=%v", req.OrgId, req.Type, req.ParentId)
	err := a.cc.Invoke(ctx, "GET", path, req, respVo)
	if err != nil {
		return nil, err
	}
	if respVo.Failure() {
		return nil, errors.Wrapf(comomPb.ErrorResourceNotExist("app not exist"), "get app failed, req:%v, error:%v", req, respVo.Error())
	}

	return respVo, nil
}
