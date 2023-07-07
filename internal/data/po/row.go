package po

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"

	"github.com/star-table/go-table/internal/data/facade/vo/permissionvo/appauth"

	pb "github.com/star-table/interface/golang/table/v1"

	"github.com/spf13/cast"

	"github.com/star-table/go-common/pkg/encoding"
	"github.com/star-table/go-common/utils/unsafe"

	"github.com/star-table/go-table/internal/data/consts"
)

func NewRow() *Row {
	r := &Row{UserIds: []int64{}, DeptIds: []int64{}, UserIdsMap: map[int64]struct{}{},
		DeptIdsMap: map[int64]struct{}{}, memberColumns: map[string]struct{}{}, relateColumns: map[string]struct{}{},
		RelateIssueIds: map[int64]struct{}{}, OriginRelateIssueIds: map[string]map[string][]int64{}}

	r.Buf.WriteRune(consts.ArrCharLeft)

	return r
}

type Row struct {
	UserIds              []int64                       `gorm:"-"`
	DeptIds              []int64                       `gorm:"-"`
	Buf                  bytes.Buffer                  `gorm:"-"`
	MaxUpdateTime        time.Time                     `gorm:"-"`
	RowCount             int                           `gorm:"-"`
	RelateIssueIds       map[int64]struct{}            `gorm:"-"` // 获取关联的id
	OriginRelateIssueIds map[string]map[string][]int64 `gorm:"-"`
	UserIdsMap           map[int64]struct{}            `gorm:"-"`
	DeptIdsMap           map[int64]struct{}            `gorm:"-"`

	relateColumns             map[string]struct{}
	memberColumns             map[string]struct{}
	appAuthData               *appauth.GetAppAuthData
	tableIdStr                string
	needChangeId              bool
	needCheckColumnPermission bool
}

func (r *Row) TableName() string {
	return "lc_data"
}

func (r *Row) AddColumnData(memberColumns []*pb.Column, relateColumnIds []string, needChangeId bool) {
	r.needChangeId = needChangeId

	for _, column := range memberColumns {
		r.memberColumns[column.Name] = struct{}{}
	}

	for _, column := range relateColumnIds {
		r.relateColumns[column] = struct{}{}
	}
}

func (r *Row) AddAuthData(tableId int64, appAuthData string) {
	if appAuthData != "" {
		r.tableIdStr = cast.ToString(tableId)
		r.appAuthData = &appauth.GetAppAuthData{}
		_ = encoding.GetJsonCodec().Unmarshal(unsafe.StringBytes(appAuthData), r.appAuthData)
		r.needCheckColumnPermission = !r.appAuthData.HasAppRootPermission && !r.appAuthData.HasAllFieldAuthOfTable(r.tableIdStr)
	}
}

func (r *Row) SetUserDeptIds() {
	for id := range r.UserIdsMap {
		if id != 0 {
			r.UserIds = append(r.UserIds, id)
		}
	}
	for id := range r.DeptIdsMap {
		if id != 0 {
			r.DeptIds = append(r.DeptIds, id)
		}
	}
}

// Scan 自定义扫描数据
func (r *Row) Scan(values []interface{}, columns []string) {
	i := 0
	if r.RowCount != 0 {
		r.Buf.WriteRune(consts.SpiltChar)
	}
	r.RowCount++
	r.Buf.WriteRune(consts.ObjCharLeft)
	issueId := ""
	dataId := ""
	for idx, column := range columns {
		if values[idx] == nil {
			continue
		}
		// 外部不需要这个字段
		if column == consts.ColumnIdCollaborators {
			continue
		}

		var value interface{}
		if v, ok := values[idx].(*interface{}); ok {
			value = *v
		}
		if value == nil {
			continue
		}

		if column == consts.ColumnIdData {
			if bts, ok := value.([]byte); ok {
				if len(bts) <= 2 {
					continue
				}
			} else if str, ok2 := value.(string); ok2 {
				if len(str) <= 2 {
					continue
				}
			}
		}

		if i != 0 {
			r.Buf.WriteRune(consts.SpiltChar)
		}
		if column != consts.ColumnData {
			r.Buf.WriteRune(consts.Mark)
			// 极星遗留恶心逻辑
			if r.needChangeId {
				if column == consts.ColumnId {
					r.Buf.WriteString(consts.ColumnDataId)
				} else if column == consts.ColumnIdIssueId {
					r.Buf.WriteString(consts.ColumnId)
				} else {
					r.Buf.WriteString(column)
				}
			} else {
				r.Buf.WriteString(column)
			}

			r.Buf.WriteRune(consts.Mark)
			r.Buf.WriteRune(consts.EqualChar)
		}
		switch it := value.(type) {
		case []byte, string:
			var s string
			if bts, ok := it.([]byte); ok {
				s = unsafe.BytesString(bts)
			} else {
				s = it.(string)
			}
			if len(s) == 0 || s == consts.Null {
				r.Buf.WriteString(consts.Null)
			} else if column == consts.ColumnData {
				if len(r.memberColumns) > 0 || len(r.relateColumns) > 0 || r.needCheckColumnPermission {
					isDelete := false
					m := make(map[string]interface{}, 10)
					_ = encoding.GetJsonCodec().Unmarshal(unsafe.StringBytes(s), &m)
					for s2, i := range m {
						if _, ok := r.memberColumns[s2]; ok {
							if is, ok2 := i.([]interface{}); ok2 {
								for _, i2 := range is {
									if i3, ok3 := i2.(string); ok3 {
										if len(i3) >= 2 && i3[0] == 'U' {
											i4, _ := strconv.ParseInt(i3[2:], 10, 64)
											r.UserIdsMap[i4] = struct{}{}
										} else {
											i4, _ := strconv.ParseInt(i3, 10, 64)
											r.DeptIdsMap[i4] = struct{}{}
										}
									}
								}
							}
						}
						if _, ok := r.relateColumns[s2]; ok {
							r.collectRelateDataIds(dataId, s2, m[s2])
						}
						// 没有权限的字段删除
						if r.needCheckColumnPermission && !r.appAuthData.HasFieldViewAuth(r.tableIdStr, s2) {
							delete(m, s2)
							isDelete = true
						}
					}
					if isDelete {
						bts, _ := json.Marshal(m)
						if len(bts) > 2 {
							r.Buf.Write(bts[1 : len(bts)-1])
						} else {
							// 由于前面写了逗号，所以这个地方一定要写一个值
							r.Buf.WriteString("\"a\":0")
						}
					} else {
						r.Buf.WriteString(s[1 : len(s)-1])
					}
				} else {
					r.Buf.WriteString(s[1 : len(s)-1])
				}
			} else {
				if s[0] == consts.ArrCharLeft {
					_, _ = r.Buf.WriteString(s)
					if _, ok := r.memberColumns[column]; ok {
						ls := make([]interface{}, 0)
						_ = encoding.GetJsonCodec().Unmarshal(unsafe.StringBytes(s), &ls)
						for _, l := range ls {
							if s2, ok2 := l.(string); ok2 {
								if len(s2) >= 2 && s2[0] == 'U' {
									i4, _ := strconv.ParseInt(s2[2:], 10, 64)
									r.UserIdsMap[i4] = struct{}{}
								} else {
									i4, _ := strconv.ParseInt(s2, 10, 64)
									r.DeptIdsMap[i4] = struct{}{}
								}
							}
						}
					}
				} else if s[0] == consts.ObjCharLeft {
					if _, ok := r.relateColumns[column]; ok {
						linkM := make(map[string]interface{}, 2)
						_ = encoding.GetJsonCodec().Unmarshal(unsafe.StringBytes(s), &linkM)
						r.collectRelateDataIds(dataId, column, linkM)
					}

					if s[1] != consts.Mark && s[1] != consts.ObjCharRight {
						_, _ = r.Buf.WriteRune(consts.Mark)
						_, _ = r.Buf.WriteString(s)
						_, _ = r.Buf.WriteRune(consts.Mark)
					} else {
						_, _ = r.Buf.WriteString(s)
					}
				} else {
					if len(s) >= 2 && s[0] == consts.Mark && s[len(s)-1] == consts.Mark {
						_, _ = r.Buf.WriteString(s)
					} else if s == consts.TrueStr || s == consts.FalseStr {
						_, _ = r.Buf.WriteString(s)
					} else {
						fs, err := strconv.ParseFloat(s, 64)
						if err == nil && fs < consts.MaxFloat64ToInt64 {
							_, _ = r.Buf.WriteString(s)
						} else {
							_, _ = r.Buf.WriteRune(consts.Mark)
							_, _ = r.Buf.WriteString(s)
							_, _ = r.Buf.WriteRune(consts.Mark)
						}
					}
				}
			}
		case int64:
			ns := strconv.FormatInt(it, 10)
			if column == consts.ColumnIdIssueId {
				issueId = ns
			}
			if column == consts.ColumnId {
				dataId = ns
			}
			if column == consts.ColumnIdCreator || column == consts.ColumnIdUpdator {
				r.UserIdsMap[it] = struct{}{}
			}

			if column == consts.ColumnId || column == consts.ColumnIdCreator || column == consts.ColumnIdUpdator ||
				column == consts.ColumnIdAppId || column == consts.ColumnIdTableId {
				_, _ = r.Buf.WriteRune(consts.Mark)
				_, _ = r.Buf.WriteString(ns)
				_, _ = r.Buf.WriteRune(consts.Mark)
			} else {
				_, _ = r.Buf.WriteString(ns)
			}
		case int:
			ns := strconv.Itoa(it)
			_, _ = r.Buf.WriteString(ns)
		case int32:
			ns := strconv.Itoa(int(it))
			_, _ = r.Buf.WriteString(ns)
		case float32:
			ns := strconv.FormatFloat(float64(it), 'f', -1, 32)
			_, _ = r.Buf.WriteString(ns)
		case float64:
			ns := strconv.FormatFloat(it, 'f', -1, 64)
			_, _ = r.Buf.WriteString(ns)
		case time.Time:
			if column == consts.ColumnIdUpdateTime && it.Unix() > r.MaxUpdateTime.Unix() {
				r.MaxUpdateTime = it
			}
			_, _ = r.Buf.WriteRune(consts.Mark)
			r.Buf.WriteString(it.Format(consts.DateFormat))
			_, _ = r.Buf.WriteRune(consts.Mark)
		default:
			_, _ = r.Buf.WriteRune(consts.Mark)
			r.Buf.WriteString(cast.ToString(it))
			_, _ = r.Buf.WriteRune(consts.Mark)
		}
		i++
	}

	// 最后再把issueId拼上，因为里面把issueId改为了id
	if r.needChangeId && issueId != "" {
		r.Buf.WriteRune(consts.SpiltChar)
		r.Buf.WriteRune(consts.Mark)
		r.Buf.WriteString(consts.ColumnIdIssueId)
		r.Buf.WriteRune(consts.Mark)
		r.Buf.WriteRune(consts.EqualChar)
		r.Buf.WriteString(issueId)
	}

	r.Buf.WriteRune(consts.ObjCharRight)
}

// 收集关联的数据id
func (r *Row) collectRelateDataIds(dataId, column string, data interface{}) {
	if m, ok := data.(map[string]interface{}); ok {
		for _, value := range m {
			if list, ok2 := value.([]interface{}); ok2 {
				for _, idString := range list {
					id := cast.ToInt64(idString)
					if id != 0 {
						r.RelateIssueIds[id] = struct{}{}
						if r.OriginRelateIssueIds[dataId] == nil {
							r.OriginRelateIssueIds[dataId] = make(map[string][]int64, 1)
						}
						r.OriginRelateIssueIds[dataId][column] = append(r.OriginRelateIssueIds[dataId][column], id)
					}
				}
			}
		}
	}
}
