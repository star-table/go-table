package biz

import (
	"context"

	"github.com/star-table/go-table/internal/data/facade/vo/permissionvo"
)

type PermissionRepo interface {
	InitAppPermissionFieldAuthCreateTable(ctx context.Context, req *permissionvo.InitAppPermissionFieldAuthCreateTableReq) error
	InitAppPermissionFieldAuthDeleteTable(ctx context.Context, req *permissionvo.InitAppPermissionFieldAuthDeleteTableReq) error
}
