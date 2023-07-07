package usercentervo

import "github.com/star-table/go-table/internal/data/facade/vo"

var (
	MemberUser = 1
)

type GetMemberSimpleInfoReq struct {
	// 组织id
	OrgId int64 `json:"orgId"`
	// 类型(1成员2部门3角色)
	Type int `json:"type"`
}

type GetMemberSimpleInfoResp struct {
	vo.Err
	Data MemberInfo `json:"data"`
}

type MemberInfo struct {
	Data []SimpleInfo `json:"data"`
}

type SimpleInfo struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
	ParentId int64  `json:"parentId"` //父部门id
}
