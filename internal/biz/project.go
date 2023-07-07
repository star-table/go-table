package biz

import (
	"context"

	"github.com/star-table/go-table/internal/data/facade/vo"
	"github.com/star-table/go-table/internal/data/facade/vo/projectvo"
	msgPb "github.com/star-table/interface/golang/msg/v1"
)

type ProjectRepo interface {
	ReportTableEvent(ctx context.Context, eventType msgPb.EventType, traceId string, event *projectvo.TableEvent) (*vo.CommonRespVo, error)
}
