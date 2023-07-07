package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/star-table/go-table/internal/data/facade/vo/projectvo"

	pushPb "github.com/star-table/interface/golang/push/v1"

	"github.com/star-table/go-common/utils/goroutine"

	"github.com/star-table/go-table/internal/data/facade/vo/permissionvo"

	"github.com/star-table/go-table/internal/data/facade/vo/appvo"

	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-common/pkg/middleware/meta"
	"github.com/star-table/go-table/internal/biz/bo"
	"github.com/star-table/go-table/internal/biz/bo/covert"
	"github.com/star-table/go-table/internal/data/consts"
	"github.com/star-table/go-table/internal/data/po"
	commonPb "github.com/star-table/interface/golang/common/v1"
	msgPb "github.com/star-table/interface/golang/msg/v1"
	pb "github.com/star-table/interface/golang/table/v1"
	tablePb "github.com/star-table/interface/golang/table/v1"
	"gorm.io/gorm"
)

type LockRepo interface {
	// TryGetDistributedLock 锁住，防止并发
	TryGetDistributedLock(ctx context.Context, lockKey string, obtainDuration ...time.Duration) (func(), error)
}

// TableRepo 表头相关
type TableRepo interface {
	// StartTransactionLcGo 开启一个事务
	StartTransactionLcGo(ctx context.Context, fc func(tx *gorm.DB) error) error
	// CreateTable 创建一张表
	CreateTable(ctx context.Context, table *po.Table, columns []*po.TableColumn, tx *gorm.DB) error
	// GetTable 获取一张表
	GetTable(ctx context.Context, tableId int64) (*po.Table, error)
	// GetTables 获取一个app下的所有表
	GetTables(ctx context.Context, appId int64, tx ...*gorm.DB) ([]*po.Table, error)
	// GetTablesByApps 获取n个app下的所有表
	GetTablesByApps(ctx context.Context, appIds []int64) (map[int64][]*po.Table, error)
	// GetTablesByOrgId 获取一个组织下的所有表
	GetTablesByOrgId(ctx context.Context, orgId int64) ([]*po.Table, error)
	// UpdateTableName 更新表名字
	UpdateTableName(ctx context.Context, appId, tableId int64, name string, userId int64) error
	// DeleteTable 删除一张表
	DeleteTable(ctx context.Context, appId, tableId int64, tx *gorm.DB) error
	// CopyTables 拷贝表，如果没有传tableIds，则会拷贝整个appId下的表
	CopyTables(ctx context.Context, srcAppId, destAppId int64,
		tableIds []int64, oldToNewPermission map[int64]int64) (map[int64]int64, error)
	// SetAutoSchedule 设置自动排期开关
	SetAutoSchedule(ctx context.Context, appId, tableId int64, autoScheduleFlag int32, userId int64) error
	// GetSummeryTableId 获取汇总表的tableId
	GetSummeryTableId(ctx context.Context, orgId int64) (int64, error)
	// GetColumn 获取一列表头
	GetColumn(ctx context.Context, tableId int64, columnId string, isNeedDescription bool) (*bo.Column, error)
	// GetColumns 获取一个table下的所有表头
	GetColumns(ctx context.Context, tableId int64) ([]*bo.Column, error)
	// GetColumnsByTables 获取n个table下的表头
	GetColumnsByTables(ctx context.Context, tableIds []int64, columnIds []string, isNeedDescription bool) ([]*bo.Columns, error)
	// GetAppColumnIdsByType 获取某种类型的列名
	GetAppColumnIdsByType(ctx context.Context, appId int64, columnType string) ([]string, error)
	// GetColumnsMap 获取表头的map
	GetColumnsMap(ctx context.Context, tableId int64) (map[string]*bo.Column, error)
	// GetRefColumns 获取引用列
	GetRefColumns(ctx context.Context, columns []*tablePb.Column) (map[string]*bo.Column, error)
	// GetAppTableColumns 获取appId下的所有表
	GetAppTableColumns(ctx context.Context, appId int64, columnIds []string) ([]*bo.Columns, error)
	// CreateColumn 创建一列
	CreateColumn(ctx context.Context, poColumn *po.TableColumn, tx ...*gorm.DB) error
	// UpdateColumn 更新一列数据
	UpdateColumn(ctx context.Context, poColumn *po.TableColumn, tx ...*gorm.DB) error
	// UpdateColumnWithOldColumnId 更新column，而且更新列名
	UpdateColumnWithOldColumnId(ctx context.Context, tc *po.TableColumn, oldColumnId string, tx ...*gorm.DB) error
	// UpdateColumnAndResetOrgColumnId 更新列数据以及重置关联组织字段
	UpdateColumnAndResetOrgColumnId(ctx context.Context, tc *po.TableColumn, oldColumnId string, tx ...*gorm.DB) error
	// ChangeColumnId 更改列名
	ChangeColumnId(ctx context.Context, oldColumnId string, poColumn *po.TableColumn, tx ...*gorm.DB) error
	// UpdateColumnDescription 更新列的描述
	UpdateColumnDescription(ctx context.Context, tableId int64, columnId, description string) error
	// DeleteColumn 删除一列
	DeleteColumn(ctx context.Context, tableId int64, columnId string, tx *gorm.DB) error
	// CheckOrgColumnIdHadUseInOrg 查询一个组织字段是否在组织内使用
	CheckOrgColumnIdHadUseInOrg(ctx context.Context, orgId int64, columnId string) (bool, error)
	// CheckOrgColumnIdHadUseInTable 查询一个组织字段是否在一个表内使用
	CheckOrgColumnIdHadUseInTable(ctx context.Context, tableId int64, columnId string) (bool, error)
	// GetNewColumnId 生成一个新的列名
	GetNewColumnId(category int) string

	GetColumnCollaboratorRoleIds(column *tablePb.Column) []string
	CheckColumnCollaboratorSwitchOn(column *tablePb.Column) bool
	CheckIsCollaboratorColumn(columnType string) bool

	GetColumnRefTableInfo(column *pb.Column, columns []*pb.Column) (int64, string)
	GetColumnPropsStringValue(column *pb.Column, key string) string
}

type TableUseCase struct {
	tableRepo      TableRepo
	lockRepo       LockRepo
	appRepo        AppRepo
	datacenterRepo DatacenterRepo
	orgRepo        OrgColumnsRepo
	permissionRepo PermissionRepo
	goPushRepo     pushPb.PushHTTPClient
	projectRepo    ProjectRepo
	rows           *RowUseCase
	snowFlag       *snowflake.Node

	log *log.Helper
}

func NewTableUseCase(repo TableRepo, lockRepo LockRepo, appRepo AppRepo, datacenterRepo DatacenterRepo,
	Rows *RowUseCase, orgRepo OrgColumnsRepo, permissionRepo PermissionRepo,
	goPushRepo pushPb.PushHTTPClient, projectRepo ProjectRepo, logger log.Logger) *TableUseCase {

	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Errorf("snowflake new node failed, the error is '%v'", err)
		return nil
	}

	return &TableUseCase{
		tableRepo:      repo,
		lockRepo:       lockRepo,
		appRepo:        appRepo,
		datacenterRepo: datacenterRepo,
		orgRepo:        orgRepo,
		permissionRepo: permissionRepo,
		goPushRepo:     goPushRepo,
		projectRepo:    projectRepo,
		rows:           Rows,
		snowFlag:       node,
		log:            log.NewHelper(logger),
	}
}

// CreateSummeryTable 创建汇总表
func (t *TableUseCase) CreateSummeryTable(ctx context.Context, req *tablePb.CreateSummeryTableRequest) (*tablePb.CreateSummeryTableReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)

	summeryTableId, err := t.tableRepo.GetSummeryTableId(ctx, ch.OrgId)
	if err != nil && !commonPb.IsResourceNotExist(errors.Cause(err)) {
		return nil, err
	}

	if summeryTableId > 0 {
		return nil, errors.Wrapf(commonPb.ErrorDuplicateOperation("can not create again"), "[CreateSummeryTable] appId:%v", req.AppId)
	}

	tableId, err := t.createTable(ctx, ch, req.AppId, "汇总表", req.Columns, consts.SummaryFlagAll, true, consts.LcAppTypeForSummaryTable)
	if err != nil {
		return nil, err
	}

	return &tablePb.CreateSummeryTableReply{
		AppId:   req.AppId,
		TableId: tableId,
	}, nil
}

// CreateTable 创建普通的表
func (t *TableUseCase) CreateTable(ctx context.Context, req *tablePb.CreateTableRequest) (*tablePb.CreateTableReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	var (
		columns []*tablePb.Column
		err     error
	)
	// 需要表头数据
	if req.IsNeedColumn {
		columns, err = t.getMergeColumns(ctx, ch.OrgId, req.BasicColumns, req.NotNeedSummeryColumnIds)
		if err != nil {
			return nil, err
		}
		// 用传进来的替换掉，没有的加在后面
		for i, column := range req.Columns {
			isIn := false
			for j, mc := range columns {
				if mc.Name == column.Name {
					isIn = true
					columns[j] = req.Columns[i]
					break
				}
			}
			if !isIn {
				columns = append(columns, req.Columns[i])
			}
		}
	}

	// 检查下，只能创建一个项目汇总表
	summaryFlag := int32(consts.SummaryFlagNormal)
	if req.SummaryFlag != 0 && req.SummaryFlag != consts.SummaryFlagNormal {
		summaryFlag = req.SummaryFlag
		tables, err := t.tableRepo.GetTables(ctx, req.AppId)
		if err != nil {
			return nil, err
		}
		for _, table := range tables {
			if table.SummeryFlag == summaryFlag {
				return nil, errors.Ignore(commonPb.ErrorDuplicateOperation(fmt.Sprintf("can not duplicate create summary table")))
			}
		}
	}

	tableId, err := t.createTable(ctx, ch, req.AppId, req.Name, columns, summaryFlag, req.IsNeedColumn, req.AppType)
	if err != nil {
		return nil, err
	}

	return &tablePb.CreateTableReply{AppId: req.AppId, Table: &tablePb.TableSchema{
		TableId: tableId,
		Name:    req.Name,
		Columns: columns,
	}}, nil
}

// createTable 新建表，插入表头
func (t *TableUseCase) createTable(ctx context.Context, ch *meta.CommonHeader, appId int64, name string,
	columns []*tablePb.Column, summeryFlag int32, isNeedColumn bool, appType int32) (int64, error) {
	// 新建表
	table := &po.Table{
		Name:        name,
		OrgId:       ch.OrgId,
		AppId:       appId,
		ColumnFlag:  consts.FlagNo,
		Config:      "{}",
		Creator:     ch.UserId,
		SummeryFlag: summeryFlag,
	}
	if isNeedColumn {
		table.ColumnFlag = consts.FlagYes
	}

	// 表头数据
	poColumns, err := covert.ColumnCovert.ToPoColumns(columns, ch.OrgId, 0)
	if err != nil {
		return 0, err
	}

	err = t.tableRepo.StartTransactionLcGo(ctx, func(tx *gorm.DB) error {
		err := t.tableRepo.CreateTable(ctx, table, poColumns, tx)
		if err != nil {
			return err
		}

		// 初始化表权限
		if appType == consts.LcAppTypeForPolaris || appType == consts.LcAppTypeForForm {
			permissionReq := &permissionvo.InitAppPermissionFieldAuthCreateTableReq{
				OrgId:                      ch.OrgId,
				AppId:                      appId,
				TableId:                    table.ID,
				UserId:                     ch.UserId,
				DefaultPermissionGroupType: consts.DefaultAppPermissionGroupTypeProject,
			}
			if appType == consts.LcAppTypeForForm {
				permissionReq.DefaultPermissionGroupType = consts.DefaultAppPermissionGroupTypeForm
			}
			err = t.permissionRepo.InitAppPermissionFieldAuthCreateTable(ctx, permissionReq)
			if err != nil {
				return err
			}
		}

		return nil
	})

	t.reportEvent(ctx, ch.OrgId, appId, table.ID, ch.UserId, msgPb.EventType_TableRefresh, table, nil)

	return table.ID, errors.WithStack(err)
}

// RenameTable 重命名表
func (t *TableUseCase) RenameTable(ctx context.Context, req *tablePb.RenameTableRequest) (*tablePb.RenameTableReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	err := t.tableRepo.UpdateTableName(ctx, req.AppId, req.TableId, req.Name, ch.UserId)

	table, _ := t.tableRepo.GetTable(ctx, req.TableId)
	if table != nil {
		t.reportEvent(ctx, ch.OrgId, req.AppId, req.TableId, ch.UserId, msgPb.EventType_TableRefresh, table, nil)
	}

	return &tablePb.RenameTableReply{TableId: req.TableId, Name: req.Name}, err
}

// CopyTables 拷贝表以及列数据，目前用于应用模板
func (t *TableUseCase) CopyTables(ctx context.Context, req *tablePb.CopyTablesRequest) (*tablePb.CopyTablesReply, error) {
	oldToNewId, err := t.tableRepo.CopyTables(ctx, req.SrcAppId, req.DstAppId, req.SrcTableIds, req.OldToNewPermission)
	return &tablePb.CopyTablesReply{AppId: req.DstAppId, OldToNewTableId: oldToNewId}, err
}

// DeleteTable 删除一张表
func (t *TableUseCase) DeleteTable(ctx context.Context, req *tablePb.DeleteTableRequest) (*tablePb.DeleteTableReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	columns, err := t.tableRepo.GetColumns(ctx, req.TableId)
	if err != nil {
		return nil, err
	}

	err = t.tableRepo.StartTransactionLcGo(ctx, func(tx *gorm.DB) error {
		err = t.tableRepo.DeleteTable(ctx, req.AppId, req.TableId, tx)
		if err != nil {
			return err
		}

		for _, column := range columns {
			// 如果是关联字段，则删除对应的表的关联字段，关联字段是成对出现的
			if column.ColumnType == tablePb.ColumnType_relating.String() {
				err = t.deleteRelateColumn(ctx, ch.OrgId, column, tx)
				if err != nil {
					return err
				}
			}
		}

		permissionReq := &permissionvo.InitAppPermissionFieldAuthDeleteTableReq{
			OrgId:                      ch.OrgId,
			AppId:                      req.AppId,
			TableId:                    req.TableId,
			UserId:                     ch.UserId,
			DefaultPermissionGroupType: consts.DefaultAppPermissionGroupTypeProject,
		}

		return t.permissionRepo.InitAppPermissionFieldAuthDeleteTable(ctx, permissionReq)
	})

	t.reportEvent(ctx, ch.OrgId, req.AppId, req.TableId, ch.UserId, msgPb.EventType_TableDeleted, nil, nil)

	return &tablePb.DeleteTableReply{TableId: req.TableId}, err
}

// SetAutoSchedule 设置自动排期开关
func (t *TableUseCase) SetAutoSchedule(ctx context.Context, req *tablePb.SetAutoScheduleRequest) (*tablePb.SetAutoScheduleReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	err := t.tableRepo.StartTransactionLcGo(ctx, func(tx *gorm.DB) error {
		err := t.tableRepo.SetAutoSchedule(ctx, req.AppId, req.TableId, req.AutoScheduleFlag, ch.UserId)
		if err != nil {
			return err
		}
		return nil
	})
	return &tablePb.SetAutoScheduleReply{
		TableId:          req.TableId,
		AutoScheduleFlag: req.AutoScheduleFlag,
	}, err
}

// ReadTables 读取一个app下的所有表
func (t *TableUseCase) ReadTables(ctx context.Context, req *tablePb.ReadTablesRequest) (*tablePb.ReadTablesReply, error) {
	tables, err := t.tableRepo.GetTables(ctx, req.AppId)
	if err != nil {
		return nil, err
	}

	reply := &tablePb.ReadTablesReply{AppId: req.AppId, Tables: covert.TableCovert.ToPbTables(tables)}
	return reply, nil
}

func (t *TableUseCase) ReadTable(ctx context.Context, req *tablePb.ReadTableRequest) (*tablePb.ReadTableReply, error) {
	table, err := t.tableRepo.GetTable(ctx, req.TableId)
	if err != nil {
		return nil, err
	}

	reply := &tablePb.ReadTableReply{Table: covert.TableCovert.ToPbTable(table)}

	return reply, nil
}

func (t *TableUseCase) ReadTablesByApps(ctx context.Context, req *tablePb.ReadTablesByAppsRequest) (*tablePb.ReadTablesByAppsReply, error) {
	tablesMap, err := t.tableRepo.GetTablesByApps(ctx, req.AppIds)
	if err != nil {
		return nil, err
	}

	reply := &tablePb.ReadTablesByAppsReply{}
	for appId, tables := range tablesMap {
		reply.AppsTables = append(reply.AppsTables, &tablePb.AppTables{
			AppId:  appId,
			Tables: covert.TableCovert.ToPbTables(tables),
		})
	}

	return reply, nil
}

func (t *TableUseCase) GetOrgTables(ctx context.Context, req *tablePb.ReadOrgTablesRequest) (*tablePb.ReadOrgTablesReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	tables, err := t.tableRepo.GetTablesByOrgId(ctx, ch.OrgId)
	if err != nil {
		return nil, err
	}

	return &tablePb.ReadOrgTablesReply{Tables: covert.TableCovert.ToPbTables(tables)}, nil
}

func (t *TableUseCase) ReadSummeryTableId(ctx context.Context, req *tablePb.ReadSummeryTableIdRequest) (*tablePb.ReadSummeryTableIdReply, error) {
	ch := meta.GetCommonHeaderFromCtx(ctx)
	tableId, err := t.tableRepo.GetSummeryTableId(ctx, ch.OrgId)
	if err != nil {
		return nil, err
	}

	return &tablePb.ReadSummeryTableIdReply{TableId: tableId}, nil
}

// reportEvent 上报事件
func (t *TableUseCase) reportEvent(ctx context.Context, orgId, appId, tableId, userId int64, eventType msgPb.EventType, newValue interface{}, oldValue interface{}) error {
	goroutine.SafeRun(func() {
		newCtx := context.Background()
		appInfo, _ := t.appRepo.GetAppInfoByAppId(newCtx, &appvo.GetAppInfoByAppIdReq{
			AppId: appId,
			OrgId: orgId,
		})
		var projectId int64
		if appInfo != nil {
			projectId = appInfo.Data.ProjectId
		}

		tableEvent := &projectvo.TableEvent{}
		tableEvent.OrgId = orgId
		tableEvent.AppId = appId
		tableEvent.ProjectId = projectId
		tableEvent.TableId = tableId
		tableEvent.UserId = userId
		if newValue != nil {
			tableEvent.New = newValue
		}
		if oldValue != nil {
			tableEvent.Old = oldValue
		}

		_, err := t.projectRepo.ReportTableEvent(newCtx, eventType, "", tableEvent)
		if err != nil {
			log.Errorf("[reportEvent] ReportTableEvent, err: %v", err)
		}
	}, t.log)

	return nil
}
