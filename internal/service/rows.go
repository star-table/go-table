package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/star-table/go-table/internal/biz"

	pb "github.com/star-table/interface/golang/table/v1"
)

type RowsService struct {
	pb.UnimplementedRowsServer

	row *biz.RowUseCase
	log *log.Helper
}

func NewRowsService(row *biz.RowUseCase, logger log.Logger) *RowsService {
	return &RowsService{
		row: row,
		log: log.NewHelper(logger),
	}
}

func (s *RowsService) RecycleAttachment(ctx context.Context, req *pb.RecycleAttachmentRequest) (*pb.RecycleAttachmentReply, error) {
	return s.row.RecycleAttachment(ctx, req)
}

func (s *RowsService) RecoverAttachment(ctx context.Context, req *pb.RecoverAttachmentRequest) (*pb.RecoverAttachmentReply, error) {
	return s.row.RecoverAttachment(ctx, req)
}

func (s *RowsService) DeleteValues(ctx context.Context, req *pb.DeleteValuesRequest) (*pb.DeleteValuesReply, error) {
	return &pb.DeleteValuesReply{}, nil
}

func (s *RowsService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	return s.row.List(ctx, req)
}

func (s *RowsService) ListRaw(ctx context.Context, req *pb.ListRawRequest) (*pb.ListRawReply, error) {
	if req.Size > 50000 || req.Size <= 0 {
		req.Size = 50000
	}
	return s.row.ListRaw(ctx, req)
}

func (s *RowsService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return s.row.Delete(ctx, req)
}

func (s *RowsService) CheckIsAppCollaborator(ctx context.Context, req *pb.CheckIsAppCollaboratorRequest) (*pb.CheckIsAppCollaboratorReply, error) {
	return s.row.CheckIsAppCollaborator(ctx, req)
}

func (s *RowsService) GetUserAppCollaboratorRoles(ctx context.Context, req *pb.GetUserAppCollaboratorRolesRequest) (*pb.GetUserAppCollaboratorRolesReply, error) {
	return s.row.GetUserAppCollaboratorRoles(ctx, req)
}

func (s *RowsService) GetAppCollaboratorRoles(ctx context.Context, req *pb.GetAppCollaboratorRolesRequest) (*pb.GetAppCollaboratorRolesReply, error) {
	return s.row.GetAppCollaboratorRoles(ctx, req)
}

func (s *RowsService) GetDataCollaborators(ctx context.Context, req *pb.GetDataCollaboratorsRequest) (*pb.GetDataCollaboratorsReply, error) {
	return s.row.GetDataCollaborators(ctx, req)
}

func (s *RowsService) ExchangeSummaryCondition(ctx context.Context, req *pb.ExchangeSummaryConditionRequest) (*pb.ExchangeSummaryConditionReply, error) {
	return &pb.ExchangeSummaryConditionReply{}, nil
}
