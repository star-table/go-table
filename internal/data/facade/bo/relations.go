package bo

type RowRelationList struct {
	Total int32              `json:"total"`
	List  []*RelationRowInfo `json:"list"`
}

type RelationRowInfo struct {
	RowId      int64             `json:"rowId"`
	RowName    string            `json:"rowName"`
	Owner      map[string]string `json:"owner"`
	TableId    int64             `json:"tableId"`
	TableName  string            `json:"tableName"`
	StatusId   int64             `json:"statusId"`
	StatusName string            `json:"statusName"`
	Title      string            `json:"title"`
	Type       int32             `json:"type"`
}
