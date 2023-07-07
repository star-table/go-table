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

var OrgColumnCovert = &orgColumnCovert{}

type orgColumnCovert struct {
}

func (c *orgColumnCovert) ToColumn(orgColumn *po.OrgColumn) (*bo.OrgColumn, error) {
	boORgColumn := &bo.OrgColumn{}
	boORgColumn.OrgId = orgColumn.OrgId
	boORgColumn.ColumnType = orgColumn.ColumnType
	boORgColumn.ColumnId = orgColumn.ColumnId
	boORgColumn.Schema = &pb.Column{}
	err := encoding.GetJsonCodec().Unmarshal([]byte(orgColumn.Schema), boORgColumn.Schema)

	return boORgColumn, errors.WithStack(err)
}

func (c *orgColumnCovert) ToColumns(orgColumns []*po.OrgColumn) ([]*bo.OrgColumn, error) {
	boOrgColumns := make([]*bo.OrgColumn, 0, len(orgColumns))
	for _, orgColumn := range orgColumns {
		temp, err := c.ToColumn(orgColumn)
		if err != nil {
			return nil, err
		}
		boOrgColumns = append(boOrgColumns, temp)
	}

	return boOrgColumns, nil
}

func (c *orgColumnCovert) ToPbColumns(boOrgColumns []*bo.OrgColumn) []*pb.Column {
	pbColumns := make([]*pb.Column, 0, len(boOrgColumns))
	for _, boOrgColumn := range boOrgColumns {
		pbColumns = append(pbColumns, boOrgColumn.Schema)
	}

	return pbColumns
}

func (c *orgColumnCovert) ToPoColumn(orgColumn *pb.Column, orgId int64) (*po.OrgColumn, error) {
	orgColumn.IsOrg = true
	schema, err := protojson.Marshal(orgColumn)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &po.OrgColumn{
		OrgId:      orgId,
		ColumnId:   orgColumn.Name,
		ColumnType: orgColumn.Field.Type.String(),
		Schema:     unsafe.BytesString(schema),
	}, nil
}

func (c *orgColumnCovert) ToPoColumns(columns []*pb.Column, orgId int64) ([]*po.OrgColumn, error) {
	poOrgColumns := make([]*po.OrgColumn, 0, len(columns))
	for _, column := range columns {
		if column.Name == "" || column.Label == "" {
			continue
		}
		temp, err := c.ToPoColumn(column, orgId)
		if err != nil {
			return nil, err
		}
		poOrgColumns = append(poOrgColumns, temp)
	}

	return poOrgColumns, nil
}
