package data

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cast"
	"github.com/star-table/go-table/internal/biz/bo"
	"github.com/star-table/go-table/internal/data/consts"
	tablePb "github.com/star-table/interface/golang/table/v1"
)

func (r *rowRepo) CheckIsAppCollaborator(ctx context.Context, orgId, appId, userId int64) (bool, error) {
	var c int64
	userIdStr := fmt.Sprintf("%s%d", consts.UserPrefix, userId)
	sql := fmt.Sprintf(
		`SELECT COUNT(1) AS c FROM lc_data WHERE "orgId"=%d AND "appId"=%d AND collaborators && ARRAY['%s'] AND "recycleFlag"=2`,
		orgId, appId, userIdStr)
	err := r.getDb(tablePb.DbType_master).WithContext(ctx).Raw(sql).Scan(&c).Error
	return c > 0, err
}

func (r *rowRepo) GetUserAppCollaboratorColumns(ctx context.Context, orgId, appId, userId int64) ([]*bo.CollaboratorColumn, error) {
	userIdStr := fmt.Sprintf("%s%d", consts.UserPrefix, userId)
	type Result struct {
		TableId  int64  `gorm:"column:tableId"`
		ColumnId string `gorm:"column:columnId"`
	}
	var res []*Result
	sql := fmt.Sprintf(
		`SELECT DISTINCT "tableId", collaborators[UNNEST(ARRAY_POSITIONS(collaborators, '%s'))+1] "columnId" FROM lc_data WHERE "orgId"=%d AND "appId"=%d AND collaborators && ARRAY['%s'] AND "recycleFlag"=2`,
		userIdStr, orgId, appId, userIdStr)
	err := r.getDb(tablePb.DbType_master).WithContext(ctx).Raw(sql).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	var ccs []*bo.CollaboratorColumn
	for _, re := range res {
		ccs = append(ccs, &bo.CollaboratorColumn{
			Id:       userIdStr,
			TableId:  re.TableId,
			ColumnId: re.ColumnId,
		})
	}
	return ccs, err
}

func (r *rowRepo) GetAppCollaboratorColumns(ctx context.Context, orgId, appId int64) ([]*bo.CollaboratorColumn, error) {
	var res []*bo.CollaboratorColumn
	sql := fmt.Sprintf(
		`SELECT DISTINCT b[1] id, "tableId", b[2] "columnId" FROM (SELECT "tableId","issueId"*100+(t1.n+1)/2 a, ARRAY_AGG(t1.c) b FROM lc_data, UNNEST(collaborators) WITH ORDINALITY t1(c, n) WHERE "orgId"=%d AND "appId"=%d AND "recycleFlag"=2 GROUP BY "tableId", a) t2`,
		orgId, appId)
	err := r.getDb(tablePb.DbType_master).WithContext(ctx).Raw(sql).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *rowRepo) GetDataCollaborators(ctx context.Context, orgId int64, dataIds []int64) ([]*tablePb.DataCollaborators, error) {
	type Result struct {
		Id            int64
		Collaborators string
	}
	res := make([]*Result, 0)
	sql := fmt.Sprintf(`SELECT id, collaborators FROM lc_data WHERE "orgId"=%d AND "id" IN (%s)`,
		orgId, strings.Join(cast.ToStringSlice(dataIds), ","))
	err := r.getDb(tablePb.DbType_master).WithContext(ctx).Raw(sql).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	var cs []*tablePb.DataCollaborators
	for _, re := range res {
		if len(re.Collaborators) <= 2 {
			continue
		}

		var ids []string
		css := strings.Split(re.Collaborators[1:len(re.Collaborators)-1], ",")
		for i, id := range css {
			if i%2 == 0 {
				ids = append(ids, id)
			}
		}
		cs = append(cs, &tablePb.DataCollaborators{
			DataId: re.Id,
			Ids:    ids,
		})
	}
	return cs, nil
}

func (r *rowRepo) SwitchColumnCollaboratorOn(ctx context.Context, orgId, appId, tableId int64, columnId string) error {
	sql := fmt.Sprintf(`UPDATE lc_data SET collaborators = collaborators_add_column(collaborators, '%s', data->>'%s') WHERE "orgId"=%d AND "appId"=%d AND "tableId"=%d AND jsonb_array_length(data->'%s') > 0;`,
		columnId, columnId, orgId, appId, tableId, columnId)
	return r.getDb(tablePb.DbType_master).WithContext(ctx).Exec(sql).Error
}

func (r *rowRepo) SwitchColumnCollaboratorOff(ctx context.Context, orgId, appId, tableId int64, columnId string) error {
	sql := fmt.Sprintf(`UPDATE lc_data SET collaborators = collaborators_delete_column(collaborators, '%s') WHERE "orgId"=%d AND "appId"=%d AND "tableId"=%d AND collaborators && ARRAY['%s'];`,
		columnId, orgId, appId, tableId, columnId)
	return r.getDb(tablePb.DbType_master).WithContext(ctx).Exec(sql).Error
}

func (r *rowRepo) CopyColumnCollaborator(ctx context.Context, orgId, appId, tableId int64, fromColumnId, toColumnId string) error {
	sql := fmt.Sprintf(`UPDATE lc_data SET collaborators = collaborators_copy_column(collaborators, '%s', '%s') WHERE "orgId"=%d AND "appId"=%d AND "tableId"=%d AND collaborators && ARRAY['%s'];`,
		fromColumnId, toColumnId, orgId, appId, tableId, fromColumnId)
	return r.getDb(tablePb.DbType_master).WithContext(ctx).Exec(sql).Error
}
