package biz

import (
	"context"

	commonPb "github.com/star-table/interface/golang/common/v1"

	"github.com/star-table/go-common/pkg/errors"

	"github.com/star-table/go-table/internal/biz/bo/covert"

	"github.com/star-table/go-table/internal/biz/bo"

	"github.com/star-table/go-common/pkg/middleware/meta"

	pb "github.com/star-table/interface/golang/table/v1"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/star-table/go-table/internal/data/po"
)

type OrgColumnsRepo interface {
	GetColumns(ctx context.Context, orgId int64, columnIds []string) ([]*bo.OrgColumn, error)
	CreateColumns(ctx context.Context, orgId int64, columns []*po.OrgColumn) error
	DeleteColumn(ctx context.Context, orgId int64, columnId string) error
}

type OrgColumnsUseCase struct {
	orgRepo   OrgColumnsRepo
	tableRepo TableRepo
	log       *log.Helper
}

func NewOrgColumnsUseCase(repo OrgColumnsRepo, tableRepo TableRepo, logger log.Logger) *OrgColumnsUseCase {
	return &OrgColumnsUseCase{
		orgRepo:   repo,
		tableRepo: tableRepo,
		log:       log.NewHelper(logger),
	}
}

func (o *OrgColumnsUseCase) InitOrgColumns(ctx context.Context, req *pb.InitOrgColumnsRequest) (
	*pb.InitOrgColumnsReply, error) {

	ch := meta.GetCommonHeaderFromCtx(ctx)
	poOrgColumns, err := covert.OrgColumnCovert.ToPoColumns(req.Columns, ch.OrgId)
	if err != nil {
		return nil, err
	}

	err = o.orgRepo.CreateColumns(ctx, ch.OrgId, poOrgColumns)

	return nil, err
}

func (o *OrgColumnsUseCase) ReadOrgColumns(ctx context.Context, req *pb.ReadOrgColumnsRequest) (
	*pb.ReadOrgColumnsReply, error) {

	ch := meta.GetCommonHeaderFromCtx(ctx)
	boOrgColumns, err := o.orgRepo.GetColumns(ctx, ch.OrgId, nil)
	if err != nil {
		return nil, err
	}

	return &pb.ReadOrgColumnsReply{Columns: covert.OrgColumnCovert.ToPbColumns(boOrgColumns)}, nil
}

func (o *OrgColumnsUseCase) CreateOrgColumn(ctx context.Context, req *pb.CreateOrgColumnRequest) (
	*pb.CreateOrgColumnReply, error) {

	ch := meta.GetCommonHeaderFromCtx(ctx)
	poOrgColumn, err := covert.OrgColumnCovert.ToPoColumn(req.Column, ch.OrgId)
	if err != nil {
		return nil, err
	}
	err = o.orgRepo.CreateColumns(ctx, ch.OrgId, []*po.OrgColumn{poOrgColumn})

	return &pb.CreateOrgColumnReply{}, err
}

func (o *OrgColumnsUseCase) DeleteOrgColumn(ctx context.Context, req *pb.DeleteOrgColumnRequest) (
	*pb.DeleteOrgColumnReply, error) {

	ch := meta.GetCommonHeaderFromCtx(ctx)
	hadUse, err := o.tableRepo.CheckOrgColumnIdHadUseInOrg(ctx, ch.OrgId, req.ColumnId)
	if err != nil {
		return nil, err
	}
	if hadUse {
		return nil, errors.Ignore(commonPb.ErrorCanNotDeleteByUse("can not delete by use"))
	}

	err = o.orgRepo.DeleteColumn(ctx, ch.OrgId, req.ColumnId)

	return &pb.DeleteOrgColumnReply{}, err
}
