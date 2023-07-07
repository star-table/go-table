package projectvo

type TableEvent struct {
	OrgId     int64       `json:"orgId"`
	AppId     int64       `json:"appId,string"`
	ProjectId int64       `json:"projectId"`
	TableId   int64       `json:"tableId,string"`
	UserId    int64       `json:"userId"`
	New       interface{} `json:"new,omitempty"`
	Old       interface{} `json:"old,omitempty"`
}
