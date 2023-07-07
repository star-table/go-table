package bo

import pb "github.com/star-table/interface/golang/table/v1"

type OrgColumn struct {
	OrgId      int64
	ColumnId   string
	ColumnType string
	Schema     *pb.Column
}
