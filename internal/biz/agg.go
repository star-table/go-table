package biz

import (
	"github.com/spf13/cast"
	"github.com/star-table/go-table/internal/biz/bo"
	"github.com/star-table/go-table/internal/data/consts"
)

const (
	aggTypeValue         = ""
	aggTypeSum           = "sum"
	aggTypeCount         = "count"
	aggTypeDistinct      = "distinct"
	aggTypeCountDistinct = "COUNTDISTINCT"
	aggTypeAvg           = "avg"
	aggTypeMax           = "max"
	aggTypeMin           = "min"
)

type linkInfo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// agg计算
func (r *RowUseCase) aggCalculate(list []map[string]interface{}, referenceColumnInfos []*bo.ReferenceColumnInfo, idToRelateIdsMap map[string]map[string][]int64) map[string]map[string]interface{} {
	result := map[string]map[string]interface{}{}
	dataMap := make(map[int64]map[string]interface{}, len(list))
	for _, data := range list {
		id := cast.ToInt64(data[consts.ColumnIdIssueId])
		dataMap[id] = data
	}

	var tempCalculateValue interface{}
	for originId, idsMap := range idToRelateIdsMap {
		for _, column := range referenceColumnInfos {
			tempValues := make([]interface{}, 0, len(idsMap[column.RelateColumnId]))
			for _, id := range idsMap[column.RelateColumnId] {
				if values, ok := dataMap[id]; ok && values[column.OriginColumnId] != nil && !column.IsRelate {
					if it, ok2 := values[column.OriginColumnId].([]interface{}); ok2 {
						tempValues = append(tempValues, it...)
					} else {
						tempValues = append(tempValues, values[column.OriginColumnId])
					}
				} else if column.IsRelate {
					// 这种是关联列，直接取title
					tempValues = append(tempValues, &linkInfo{Id: id, Name: cast.ToString(values[consts.ColumnIdTitle])})
				}
			}

			if len(tempValues) > 0 {
				switch column.AggFunc {
				case aggTypeSum:
					tempCalculateValue = r.aggCalculateSum(tempValues)
				case aggTypeCount:
					tempCalculateValue = len(tempValues)
				case aggTypeDistinct:
					tempCalculateValue = r.aggCalculateDistinct(tempValues)
				case aggTypeCountDistinct:
					tempCalculateValue = len(r.aggCalculateDistinct(tempValues))
				case aggTypeAvg:
					tempCalculateValue = r.aggCalculateSum(tempValues) / float64(len(tempValues))
				case aggTypeMax:
					tempCalculateValue = r.aggCalculateMax(tempValues)
				case aggTypeMin:
					tempCalculateValue = r.aggCalculateMin(tempValues)
				default:
					tempCalculateValue = tempValues
				}

				if result[originId] == nil {
					result[originId] = make(map[string]interface{}, len(referenceColumnInfos))
				}
				// 关联的数据写入一个特殊的map里面，好让前端判断的时候不需要依赖表头
				if column.IsRelate {
					result[originId][column.OriginColumnId] = map[string]interface{}{"linkNames": tempCalculateValue}
				} else {
					result[originId][column.OriginColumnId] = tempCalculateValue
				}
			}
		}
	}

	return result
}

func (r *RowUseCase) aggCalculateSum(values []interface{}) float64 {
	var result float64
	for _, value := range values {
		result += cast.ToFloat64(value)
	}

	return result
}

func (r *RowUseCase) aggCalculateMax(values []interface{}) interface{} {
	var result interface{}
	for _, value := range values {
		if result == nil || cast.ToString(result) < cast.ToString(value) {
			result = value
		}
	}

	return result
}

func (r *RowUseCase) aggCalculateMin(values []interface{}) interface{} {
	var result interface{}
	for _, value := range values {
		if result == nil || cast.ToString(result) > cast.ToString(value) {
			result = value
		}
	}

	return result
}

func (r *RowUseCase) aggCalculateDistinct(values []interface{}) []interface{} {
	result := make([]interface{}, 0, len(values))
	tempMap := make(map[interface{}]struct{}, len(values))
	for _, value := range values {
		strValue := cast.ToString(value)
		// 特别的结构，这个时候直接忽略
		if strValue == "" {
			result = append(result, value)
		} else {
			if _, ok := tempMap[value]; !ok {
				result = append(result, value)
			}
			tempMap[value] = struct{}{}
		}
	}

	return result
}
