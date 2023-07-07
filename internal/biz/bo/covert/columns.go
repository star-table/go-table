package covert

import (
	encoding "github.com/star-table/go-common/pkg/encoding"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-common/utils/unsafe"
	"github.com/star-table/go-table/internal/biz/bo"
	"github.com/star-table/go-table/internal/data/po"
	pb "github.com/star-table/interface/golang/table/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

var ColumnCovert = &columnCovert{}

type columnCovert struct {
}

func (c *columnCovert) ToColumn(pc *po.TableColumn) (*bo.Column, error) {
	boColumn := &bo.Column{}
	boColumn.TableId = pc.TableId
	boColumn.ColumnType = pc.ColumnType
	boColumn.ColumnId = pc.ColumnId
	boColumn.Schema = &pb.Column{}
	err := encoding.GetJsonCodec().Unmarshal([]byte(pc.Schema), boColumn.Schema)
	boColumn.Schema.Description = pc.Description

	return boColumn, errors.WithStack(err)
}

func (c *columnCovert) ToColumns(poColumns []*po.TableColumn) ([]*bo.Column, error) {
	boColumns := make([]*bo.Column, 0, len(poColumns))
	for _, pc := range poColumns {
		temp, err := c.ToColumn(pc)
		if err != nil {
			return nil, err
		}
		boColumns = append(boColumns, temp)
	}

	return boColumns, nil
}

func (c *columnCovert) ToPbColumns(poColumns []*po.TableColumn) ([]*pb.Column, error) {
	pbColumns := make([]*pb.Column, 0, len(poColumns))
	for _, poColumn := range poColumns {
		pbColumn, err := c.ToPbColumn(poColumn)
		if err != nil {
			return nil, err
		}
		pbColumns = append(pbColumns, pbColumn)
	}

	return pbColumns, nil
}

func (c *columnCovert) ToPbColumn(poColumn *po.TableColumn) (*pb.Column, error) {
	pbColumn := &pb.Column{}
	err := encoding.GetJsonCodec().Unmarshal([]byte(poColumn.Schema), pbColumn)
	if err != nil {
		return nil, errors.Wrapf(err, "[ToPbColumns] Unmarshal error,string:%v", poColumn.Schema)
	}

	return pbColumn, nil
}

// ToPoColumn pb的类型转到po
func (c *columnCovert) ToPoColumn(column *pb.Column, orgId, tableId int64) (*po.TableColumn, error) {
	description := column.Description
	column.Description = ""
	bts, err := protojson.Marshal(column)
	if err != nil {
		return nil, errors.Wrapf(err, "[ToPoColumn] protojson.Marshal error, column:%v", column)
	}

	poColumn := &po.TableColumn{
		OrgId:       orgId,
		TableId:     tableId,
		ColumnId:    column.Name,
		ColumnType:  column.Field.Type.String(),
		Schema:      unsafe.BytesString(bts),
		Description: description,
	}

	return poColumn, nil
}

func (c *columnCovert) ToPoColumns(columns []*pb.Column, orgId, tableId int64) ([]*po.TableColumn, error) {
	poColumns := make([]*po.TableColumn, 0, len(columns))
	for _, column := range columns {
		if column.Name == "" || column.Label == "" {
			continue
		}
		temp, err := c.ToPoColumn(column, orgId, tableId)
		if err != nil {
			return nil, err
		}

		poColumns = append(poColumns, temp)
	}

	return poColumns, nil
}
