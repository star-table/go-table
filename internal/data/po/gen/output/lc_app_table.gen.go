// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package output

import (
	"context"

	"github.com/star-table/go-table/internal/data/po/gen/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newLcAppTable(db *gorm.DB) lcAppTable {
	_lcAppTable := lcAppTable{}

	_lcAppTable.lcAppTableDo.UseDB(db)
	_lcAppTable.lcAppTableDo.UseModel(&model.LcAppTable{})

	tableName := _lcAppTable.lcAppTableDo.TableName()
	_lcAppTable.ALL = field.NewField(tableName, "*")
	_lcAppTable.ID = field.NewInt64(tableName, "id")
	_lcAppTable.CreateAt = field.NewTime(tableName, "create_at")
	_lcAppTable.UpdateAt = field.NewTime(tableName, "update_at")
	_lcAppTable.OrgID = field.NewInt64(tableName, "org_id")
	_lcAppTable.AppID = field.NewInt64(tableName, "app_id")
	_lcAppTable.Config = field.NewString(tableName, "config")
	_lcAppTable.Creator = field.NewInt64(tableName, "creator")
	_lcAppTable.Updater = field.NewInt64(tableName, "updater")
	_lcAppTable.DelFlag = field.NewInt32(tableName, "del_flag")
	_lcAppTable.Name = field.NewString(tableName, "name")
	_lcAppTable.SummeryFlag = field.NewInt32(tableName, "summery_flag")
	_lcAppTable.ColumnFlag = field.NewInt32(tableName, "column_flag")
	_lcAppTable.AutoScheduleFlag = field.NewInt32(tableName, "auto_schedule_flag")
	_lcAppTable.AutoScheduleValid = field.NewInt32(tableName, "auto_schedule_valid")

	_lcAppTable.fillFieldMap()

	return _lcAppTable
}

type lcAppTable struct {
	lcAppTableDo lcAppTableDo

	ALL               field.Field
	ID                field.Int64
	CreateAt          field.Time
	UpdateAt          field.Time
	OrgID             field.Int64
	AppID             field.Int64
	Config            field.String
	Creator           field.Int64
	Updater           field.Int64
	DelFlag           field.Int32
	Name              field.String
	SummeryFlag       field.Int32
	ColumnFlag        field.Int32
	AutoScheduleFlag  field.Int32
	AutoScheduleValid field.Int32

	fieldMap map[string]field.Expr
}

func (l lcAppTable) Table(newTableName string) *lcAppTable {
	l.lcAppTableDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l lcAppTable) As(alias string) *lcAppTable {
	l.lcAppTableDo.DO = *(l.lcAppTableDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *lcAppTable) updateTableName(table string) *lcAppTable {
	l.ALL = field.NewField(table, "*")
	l.ID = field.NewInt64(table, "id")
	l.CreateAt = field.NewTime(table, "create_at")
	l.UpdateAt = field.NewTime(table, "update_at")
	l.OrgID = field.NewInt64(table, "org_id")
	l.AppID = field.NewInt64(table, "app_id")
	l.Config = field.NewString(table, "config")
	l.Creator = field.NewInt64(table, "creator")
	l.Updater = field.NewInt64(table, "updater")
	l.DelFlag = field.NewInt32(table, "del_flag")
	l.Name = field.NewString(table, "name")
	l.SummeryFlag = field.NewInt32(table, "summery_flag")
	l.ColumnFlag = field.NewInt32(table, "column_flag")
	l.AutoScheduleFlag = field.NewInt32(table, "auto_schedule_flag")
	l.AutoScheduleValid = field.NewInt32(table, "auto_schedule_valid")

	l.fillFieldMap()

	return l
}

func (l *lcAppTable) WithContext(ctx context.Context) *lcAppTableDo {
	return l.lcAppTableDo.WithContext(ctx)
}

func (l lcAppTable) TableName() string { return l.lcAppTableDo.TableName() }

func (l *lcAppTable) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *lcAppTable) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 14)
	l.fieldMap["id"] = l.ID
	l.fieldMap["create_at"] = l.CreateAt
	l.fieldMap["update_at"] = l.UpdateAt
	l.fieldMap["org_id"] = l.OrgID
	l.fieldMap["app_id"] = l.AppID
	l.fieldMap["config"] = l.Config
	l.fieldMap["creator"] = l.Creator
	l.fieldMap["updater"] = l.Updater
	l.fieldMap["del_flag"] = l.DelFlag
	l.fieldMap["name"] = l.Name
	l.fieldMap["summery_flag"] = l.SummeryFlag
	l.fieldMap["column_flag"] = l.ColumnFlag
	l.fieldMap["auto_schedule_flag"] = l.AutoScheduleFlag
	l.fieldMap["auto_schedule_valid"] = l.AutoScheduleValid
}

func (l lcAppTable) clone(db *gorm.DB) lcAppTable {
	l.lcAppTableDo.ReplaceDB(db)
	return l
}

type lcAppTableDo struct{ gen.DO }

func (l lcAppTableDo) Debug() *lcAppTableDo {
	return l.withDO(l.DO.Debug())
}

func (l lcAppTableDo) WithContext(ctx context.Context) *lcAppTableDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l lcAppTableDo) Clauses(conds ...clause.Expression) *lcAppTableDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l lcAppTableDo) Returning(value interface{}, columns ...string) *lcAppTableDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l lcAppTableDo) Not(conds ...gen.Condition) *lcAppTableDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l lcAppTableDo) Or(conds ...gen.Condition) *lcAppTableDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l lcAppTableDo) Select(conds ...field.Expr) *lcAppTableDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l lcAppTableDo) Where(conds ...gen.Condition) *lcAppTableDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l lcAppTableDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *lcAppTableDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l lcAppTableDo) Order(conds ...field.Expr) *lcAppTableDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l lcAppTableDo) Distinct(cols ...field.Expr) *lcAppTableDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l lcAppTableDo) Omit(cols ...field.Expr) *lcAppTableDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l lcAppTableDo) Join(table schema.Tabler, on ...field.Expr) *lcAppTableDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l lcAppTableDo) LeftJoin(table schema.Tabler, on ...field.Expr) *lcAppTableDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l lcAppTableDo) RightJoin(table schema.Tabler, on ...field.Expr) *lcAppTableDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l lcAppTableDo) Group(cols ...field.Expr) *lcAppTableDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l lcAppTableDo) Having(conds ...gen.Condition) *lcAppTableDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l lcAppTableDo) Limit(limit int) *lcAppTableDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l lcAppTableDo) Offset(offset int) *lcAppTableDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l lcAppTableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *lcAppTableDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l lcAppTableDo) Unscoped() *lcAppTableDo {
	return l.withDO(l.DO.Unscoped())
}

func (l lcAppTableDo) Create(values ...*model.LcAppTable) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l lcAppTableDo) CreateInBatches(values []*model.LcAppTable, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l lcAppTableDo) Save(values ...*model.LcAppTable) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l lcAppTableDo) First() (*model.LcAppTable, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LcAppTable), nil
	}
}

func (l lcAppTableDo) Take() (*model.LcAppTable, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LcAppTable), nil
	}
}

func (l lcAppTableDo) Last() (*model.LcAppTable, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LcAppTable), nil
	}
}

func (l lcAppTableDo) Find() ([]*model.LcAppTable, error) {
	result, err := l.DO.Find()
	return result.([]*model.LcAppTable), err
}

func (l lcAppTableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LcAppTable, err error) {
	buf := make([]*model.LcAppTable, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l lcAppTableDo) FindInBatches(result *[]*model.LcAppTable, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l lcAppTableDo) Attrs(attrs ...field.AssignExpr) *lcAppTableDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l lcAppTableDo) Assign(attrs ...field.AssignExpr) *lcAppTableDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l lcAppTableDo) Joins(field field.RelationField) *lcAppTableDo {
	return l.withDO(l.DO.Joins(field))
}

func (l lcAppTableDo) Preload(field field.RelationField) *lcAppTableDo {
	return l.withDO(l.DO.Preload(field))
}

func (l lcAppTableDo) FirstOrInit() (*model.LcAppTable, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LcAppTable), nil
	}
}

func (l lcAppTableDo) FirstOrCreate() (*model.LcAppTable, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LcAppTable), nil
	}
}

func (l lcAppTableDo) FindByPage(offset int, limit int) (result []*model.LcAppTable, count int64, err error) {
	if limit <= 0 {
		count, err = l.Count()
		return
	}

	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l lcAppTableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l *lcAppTableDo) withDO(do gen.Dao) *lcAppTableDo {
	l.DO = *do.(*gen.DO)
	return l
}