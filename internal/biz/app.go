package biz

import (
	"context"

	"github.com/star-table/go-table/internal/data/facade/vo/appvo"
)

type AppRepo interface {
	GetAppInfoByAppId(ctx context.Context, req *appvo.GetAppInfoByAppIdReq) (*appvo.GetAppInfoByAppIdResp, error)
	// GetRealAppInfoByAppId 如果是镜像id，则会请求真正的id
	GetRealAppInfoByAppId(ctx context.Context, req *appvo.GetAppInfoByAppIdReq) (*appvo.GetAppInfoByAppIdResp, error)
	// GetAppList 通过orgId获取app应用列表，type为5的是汇总表
	GetAppList(ctx context.Context, req *appvo.AppListReqVo) (*appvo.AppListRespVo, error)
	// GetSummeryAppId 获取汇总表appId
	GetSummeryAppId(ctx context.Context, orgId int64) (int64, error)
}
