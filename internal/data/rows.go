package data

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/spf13/cast"

	"github.com/star-table/go-table/internal/data/po"

	"github.com/star-table/go-table/internal/data/facade/vo/form"

	"github.com/star-table/go-table/pkg/jsonb"

	"gorm.io/gorm"

	"github.com/star-table/go-table/internal/data/consts"

	tablePb "github.com/star-table/interface/golang/table/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-table/internal/biz"
)

type rowRepo struct {
	data *Data
	form *formRepo
	log  *log.Helper
}

// NewRowRepo .
func NewRowRepo(data *Data, form *formRepo, logger log.Logger) biz.RowRepo {
	return &rowRepo{data: data, form: form, log: log.NewHelper(logger)}
}

func (r *rowRepo) List(ctx context.Context, req *tablePb.ListRequest, queryReq *form.QuerySqlReq,
	memberColumns []*tablePb.Column, relateColumnIds []string) (*po.Row, error) {
	reply, err := r.form.QuerySql(ctx, queryReq)
	if err != nil {
		return nil, err
	}

	log.Infof("[List] QuerySql sql:%v, args:%v", reply.Data.Sql, reply.Data.Args)

	values, err := r.parseValues(reply.Data.Args)
	if err != nil {
		return nil, err
	}

	rows := po.NewRow()
	rows.AddColumnData(memberColumns, relateColumnIds, req.NeedChangeId)
	rows.AddAuthData(queryReq.TableId, req.AppAuthData)
	db := r.getDb(req.DbType).WithContext(ctx)
	err = db.Raw(reply.Data.Sql, values...).Find(rows).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rows.Buf.WriteRune(consts.ArrCharRight)
	rows.SetUserDeptIds()

	if req.NeedTotal {
		splits := strings.Split(reply.Data.Sql, " from ")
		if len(splits) == 2 {
			orderSplit := strings.Split(splits[1], " order by ")
			countSql := "select count(*) as total from " + orderSplit[0]
			result := map[string]interface{}{}
			count := strings.Count(strings.ReplaceAll(countSql, "??", "?"), "?")
			err = db.Raw(countSql, values[0:count]...).Find(result).Error
			if err != nil {
				return nil, errors.WithStack(err)
			}
			rows.RowCount = cast.ToInt(result["total"])
		}
	}

	return rows, nil
}

func (r *rowRepo) ListRaw(ctx context.Context, req *tablePb.ListRawRequest, memberColumns []*tablePb.Column) (*po.Row, error) {
	db := r.getDb(req.DbType).WithContext(ctx)
	page := int(req.Page)
	size := int(req.Size)
	if req.Size > 0 {
		db = db.Limit(size)
	}
	if req.Page > 0 {
		db = db.Offset((page - 1) * size)
	}

	if len(req.Orders) > 0 {
		for _, order := range req.Orders {
			if order.Asc {
				db.Order(order.Column + " ASC")
			} else {
				db.Order(order.Column + " DESC")
			}
		}
	}

	if len(req.Groups) > 0 {
		for _, group := range req.Groups {
			db = db.Group(group)
		}
	}

	rows := po.NewRow()
	rows.AddColumnData(memberColumns, nil, false)
	err := db.Select(req.FilterColumns).Find(rows, jsonb.NewQuery(req.Condition)).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rows.Buf.WriteRune(consts.ArrCharRight)
	rows.SetUserDeptIds()

	return rows, nil
}

func (r *rowRepo) Delete(ctx context.Context, condition *tablePb.Condition) (int64, error) {
	rows := po.NewRow()
	db := r.getDb(tablePb.DbType_master).WithContext(ctx)
	tx := db.Delete(rows, jsonb.NewQuery(condition))
	return tx.RowsAffected, tx.Error
}

func (r *rowRepo) getDb(dbType tablePb.DbType) *gorm.DB {
	db := r.data.postgres[consts.DbTypePrefix+dbType.String()]
	if db == nil {
		db = r.data.postgres[consts.DbTypePrefix+tablePb.DbType_master.String()]
	}

	return db
}

func (r *rowRepo) parseValues(args string) ([]interface{}, error) {
	values := make([]interface{}, 0, 10)

	dec := json.NewDecoder(strings.NewReader(args))
	dec.UseNumber()
	err := dec.Decode(&values)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for i := range values {
		if s, ok := values[i].(map[string]interface{}); ok {
			values[i] = s["value"]
		}
	}

	return values, nil
}
