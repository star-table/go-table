package appvo

import (
	"github.com/star-table/go-table/internal/data/facade/vo"
)

type AppInfoRespVo struct {
	vo.Err
	AppInfo *vo.AppInfo `json:"data"`
}

type AppInfoReqVo struct {
	AppCode string `json:"appCode"`
}

type CreateAppInfoReqVo struct {
	CreateAppInfo vo.CreateAppInfoReq `json:"input"`
	UserId        int64               `json:"userId"`
	OrgId         int64               `json:"orgId"`
}

type UpdateAppInfoReqVo struct {
	Input  vo.UpdateAppInfoReq `json:"input"`
	UserId int64               `json:"userId"`
	OrgId  int64               `json:"orgId"`
}

type DeleteAppInfoReqVo struct {
	Input  vo.DeleteAppInfoReq `json:"input"`
	UserId int64               `json:"userId"`
	OrgId  int64               `json:"orgId"`
}

type AppListReqVo struct {
	OrgId    int64 `json:"orgId"`
	Type     int64 `json:"type"`
	ParentId int64 `json:"parentId"`
}

type AppListRespVo struct {
	vo.Err
	Data []*AppListData `json:"data"`
}

type AppListData struct {
	Id        int64  `json:"id"`
	OrgId     int64  `json:"orgId"`
	ExtendsId int64  `json:"extendsId"`
	Name      string `json:"name"`
	Type      int64  `json:"type"`
	Icon      string `json:"icon"`
	Status    int64  `json:"status"`
	Creator   int64  `json:"creator"`
	//CreateTime   time.Time `json:"createTime"`
	Updator int64 `json:"updator"`
	//UpdateTime   time.Time `json:"updateTime"`
	ParentId     int64 `json:"parentId"`
	FormId       int64 `json:"formId"`
	DashboardId  int64 `json:"dashboardId"`
	WorkflowFlag int64 `json:"workflowFlag"`
	TemplateFlag int64 `json:"templateFlag"`
	MirrorViewId int64 `json:"mirrorViewId"`
	MirrorAppId  int64 `json:"mirrorAppId"`
}
