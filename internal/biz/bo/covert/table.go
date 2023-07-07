package covert

import (
	"github.com/star-table/go-table/internal/data/po"
	pb "github.com/star-table/interface/golang/table/v1"
)

var TableCovert = &tableCovert{}

type tableCovert struct {
}

func (t *tableCovert) ToPbTables(tables []*po.Table) []*pb.TableMeta {
	pbMetas := make([]*pb.TableMeta, 0, len(tables))
	for _, table := range tables {
		pbMetas = append(pbMetas, t.ToPbTable(table))
	}

	return pbMetas
}

func (t *tableCovert) ToPbTable(table *po.Table) *pb.TableMeta {
	return &pb.TableMeta{
		AppId:            table.AppId,
		TableId:          table.ID,
		Name:             table.Name,
		AutoScheduleFlag: table.AutoScheduleFlag,
		SummaryFlag:      table.SummeryFlag,
	}
}
