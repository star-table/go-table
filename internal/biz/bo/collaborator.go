package bo

import pb "github.com/star-table/interface/golang/table/v1"

type CollaboratorColumn struct {
	Id       string `gorm:"column:id"`
	TableId  int64  `gorm:"column:tableId"`
	ColumnId string `gorm:"column:columnId"`
}

type ColumnWithCollaboratorRoles struct {
	Column   *pb.Column
	GotRoles bool
	Roles    []string
}
