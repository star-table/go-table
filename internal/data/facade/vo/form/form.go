package form

import "github.com/star-table/go-table/internal/data/facade/vo"

type QuerySqlReq struct {
	OrgId          int64  `json:"orgId"`
	UserId         int64  `json:"userId"`
	Query          string `json:"query"`
	FieldParams    string `json:"fieldParams"`
	SummaryTableId int64  `json:"summaryTableId"`
	TableId        int64  `json:"-"`
}

type QuerySqlResp struct {
	vo.Err
	Data *SqlData `json:"data"`
}

type SqlData struct {
	Sql  string `json:"sql"`
	Args string `json:"args"`
}

type LessIssueListReq struct {
	Condition      *vo.LessCondsData `json:"condition"`
	Orders         []*vo.LessOrder   `json:"orders"`
	RedirectIds    []int64           `json:"redirectIds"`
	AppId          int64             `json:"appId"`
	OrgId          int64             `json:"orgId"`
	UserId         int64             `json:"userId"`
	Page           int64             `json:"page"`
	Size           int64             `json:"size"`
	Columns        []string          `json:"columns"`
	FilterColumns  []string          `json:"filterColumns"`
	Groups         []string          `json:"groups"`
	TableId        int64             `json:"tableId"`
	Export         bool              `json:"export"`              // 是否是导出，导出不限制size
	NeedTotal      bool              `json:"needTotal,omitempty"` // 是否需要总数，大部分情况都不需要总数
	NeedRefColumn  bool              `json:"needRefColumn"`       // 是否需要引用列
	AggNoLimit     bool              `json:"aggNoLimit"`          // 引用列是否限制返回合并个数
	NeedDeleteData bool              `json:"needDeleteData"`      // 需要删除的数据
	NeedChangeId   bool              `json:"-"`
}
