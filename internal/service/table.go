package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-common/pkg/middleware/meta"
	"github.com/star-table/go-table/internal/biz"
	"github.com/star-table/go-table/internal/data/consts"
	commonPb "github.com/star-table/interface/golang/common/v1"
	pb "github.com/star-table/interface/golang/table/v1"
)

type TableService struct {
	pb.UnimplementedTableServer

	table *biz.TableUseCase
	row   *biz.RowUseCase
	org   *biz.OrgColumnsUseCase
	log   *log.Helper
}

func NewTableService(tuc *biz.TableUseCase, ruc *biz.RowUseCase,
	org *biz.OrgColumnsUseCase, logger log.Logger) *TableService {
	return &TableService{
		table: tuc,
		row:   ruc,
		org:   org,
		log:   log.NewHelper(logger),
	}
}

func (s *TableService) CreateSummeryTable(ctx context.Context, req *pb.CreateSummeryTableRequest) (*pb.CreateSummeryTableReply, error) {
	return s.table.CreateSummeryTable(ctx, req)
}

func (s *TableService) CreateTable(ctx context.Context, req *pb.CreateTableRequest) (*pb.CreateTableReply, error) {
	if req.AppType == consts.LcAppTypeUnknown {
		req.AppType = consts.LcAppTypeForPolaris
	}
	return s.table.CreateTable(ctx, req)
}

func (s *TableService) RenameTable(ctx context.Context, req *pb.RenameTableRequest) (*pb.RenameTableReply, error) {
	return s.table.RenameTable(ctx, req)
}

func (s *TableService) CopyTables(ctx context.Context, req *pb.CopyTablesRequest) (*pb.CopyTablesReply, error) {
	return s.table.CopyTables(ctx, req)
}

func (s *TableService) DeleteTable(ctx context.Context, req *pb.DeleteTableRequest) (*pb.DeleteTableReply, error) {
	return s.table.DeleteTable(ctx, req)
}

func (s *TableService) SetAutoSchedule(ctx context.Context, req *pb.SetAutoScheduleRequest) (*pb.SetAutoScheduleReply, error) {
	return s.table.SetAutoSchedule(ctx, req)
}

func (s *TableService) ReadTables(ctx context.Context, req *pb.ReadTablesRequest) (*pb.ReadTablesReply, error) {
	return s.table.ReadTables(ctx, req)
}

func (s *TableService) ReadTable(ctx context.Context, req *pb.ReadTableRequest) (*pb.ReadTableReply, error) {
	return s.table.ReadTable(ctx, req)
}

func (s *TableService) ReadTablesByApps(ctx context.Context, req *pb.ReadTablesByAppsRequest) (*pb.ReadTablesByAppsReply, error) {
	return s.table.ReadTablesByApps(ctx, req)
}

func (s *TableService) ReadTableSchemas(ctx context.Context, req *pb.ReadTableSchemasRequest) (*pb.ReadTableSchemasReply, error) {
	return s.table.ReadTableSchemas(ctx, req)
}

func (s *TableService) ReadTableSchemasByAppId(ctx context.Context, req *pb.ReadTableSchemasByAppIdRequest) (*pb.ReadTableSchemasByAppIdReply, error) {
	return s.table.ReadTableSchemasByAppId(ctx, req)
}

func (s *TableService) ReadSummeryTableId(ctx context.Context, req *pb.ReadSummeryTableIdRequest) (*pb.ReadSummeryTableIdReply, error) {
	return s.table.ReadSummeryTableId(ctx, req)
}

func (s *TableService) ReadOrgTableSchemas(ctx context.Context, req *pb.ReadOrgTableSchemasRequest) (*pb.ReadOrgTableSchemasReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	if ch.OrgId == 0 {
		return nil, errors.Ignore(commonPb.ErrorParamsNotCorrect("orgId can not be zero"))
	}
	return s.table.ReadOrgTableSchemas(ctx, req)
}

func (s *TableService) InitOrgColumns(ctx context.Context, req *pb.InitOrgColumnsRequest) (*pb.InitOrgColumnsReply, error) {
	return s.org.InitOrgColumns(ctx, req)
}

func (s *TableService) ReadOrgColumns(ctx context.Context, req *pb.ReadOrgColumnsRequest) (*pb.ReadOrgColumnsReply, error) {
	return s.org.ReadOrgColumns(ctx, req)
}

func (s *TableService) CreateOrgColumn(ctx context.Context, req *pb.CreateOrgColumnRequest) (*pb.CreateOrgColumnReply, error) {
	if req.Column.Name == "" || req.Column.Label == "" {
		return nil, errors.Ignore(commonPb.ErrorParamsNotCorrect("name or label can not be null"))
	}

	return s.org.CreateOrgColumn(ctx, req)
}

func (s *TableService) DeleteOrgColumn(ctx context.Context, req *pb.DeleteOrgColumnRequest) (*pb.DeleteOrgColumnReply, error) {
	return s.org.DeleteOrgColumn(ctx, req)
}

func (s *TableService) CreateColumn(ctx context.Context, req *pb.CreateColumnRequest) (*pb.CreateColumnReply, error) {
	if req.Column.Name == "" || req.Column.Label == "" {
		return nil, errors.Ignore(commonPb.ErrorParamsNotCorrect("name or label can not be null"))
	}

	if req.Column.IsOrg && req.SourceOrgColumnId == "" {
		return nil, errors.Ignore(commonPb.ErrorParamsNotCorrect("org column must has sourceOrgColumnId"))
	}

	err := s.checkSpecialColumnParams(req.Column)
	if err != nil {
		return nil, err
	}

	return s.table.CreateColumn(ctx, req)
}

func (s *TableService) checkSpecialColumnParams(column *pb.Column) error {
	// 如果是引用
	if column.Field.Type.String() == pb.ColumnType_reference.String() {
		reference := column.Field.Props.GetFields()[column.Field.Type.String()]
		if reference == nil || reference.GetStructValue() == nil {
			return errors.Ignore(commonPb.ErrorParamsNotCorrect("reference type column must need reference struct"))
		}
		relateColumnId := reference.GetStructValue().Fields[consts.ReferenceColumnId]
		columnId := reference.GetStructValue().Fields[consts.RelateColumnId]
		if relateColumnId == nil || relateColumnId.GetStringValue() == "" || columnId == nil || columnId.GetStringValue() == "" {
			return errors.Ignore(commonPb.ErrorParamsNotCorrect("reference type column must need reference struct"))
		}
	}

	return nil
}

func (s *TableService) CopyColumn(ctx context.Context, req *pb.CopyColumnRequest) (*pb.CopyColumnReply, error) {
	return s.table.CopyColumn(ctx, req)
}

func (s *TableService) UpdateColumn(ctx context.Context, req *pb.UpdateColumnRequest) (*pb.UpdateColumnReply, error) {
	if req.Column.Name == "" || req.Column.Label == "" {
		return nil, errors.Ignore(commonPb.ErrorParamsNotCorrect("name or label can not be null"))
	}

	err := s.checkSpecialColumnParams(req.Column)
	if err != nil {
		return nil, err
	}

	return s.table.UpdateColumn(ctx, req)
}

func (s *TableService) UpdateColumnDescription(ctx context.Context, req *pb.UpdateColumnDescriptionRequest) (*pb.UpdateColumnDescriptionReply, error) {
	return s.table.UpdateColumnDescription(ctx, req)
}

func (s *TableService) DeleteColumn(ctx context.Context, req *pb.DeleteColumnRequest) (*pb.DeleteColumnReply, error) {
	return s.table.DeleteColumn(ctx, req)
}

func (s *TableService) CreateRows(ctx context.Context, req *pb.CreateRowsRequest) (*pb.CreateRowsReply, error) {
	return s.row.CreateRows(ctx, req)
}

func (s *TableService) MoveRow(ctx context.Context, req *pb.MoveRowRequest) (*pb.MoveRowReply, error) {
	return s.row.MoveRow(ctx, req)
}

func (s *TableService) CopyRow(ctx context.Context, req *pb.CopyRowRequest) (*pb.CopyRowReply, error) {
	return s.row.CopyRow(ctx, req)
}

func (s *TableService) DeleteRow(ctx context.Context, req *pb.DeleteRowRequest) (*pb.DeleteRowReply, error) {
	return s.row.DeleteRow(ctx, req)
}

func (s *TableService) ReadOrgTables(ctx context.Context, req *pb.ReadOrgTablesRequest) (*pb.ReadOrgTablesReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	if ch.OrgId == 0 {
		return nil, errors.Ignore(commonPb.ErrorParamsNotCorrect("orgId can not be zero"))
	}
	return s.table.GetOrgTables(ctx, req)
}
