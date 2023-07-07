package datacentervo

import (
	"fmt"
	"strings"

	"github.com/star-table/go-table/internal/data/consts"

	"github.com/star-table/go-table/internal/data/facade/vo"
)

const (
	ExecuteTypeUpdate = 1
	ExecuteTypeDelete = 2
	ExecuteTypeInsert = 3
)

const (
	ConditionEqual    = "equal"
	ConditionIn       = "in"
	ConditionAnd      = "and"
	ConditionValuesIn = "values_in"
)

const (
	SetTypeNormal = 1
	SetTypeJson   = 2
)

type QueryReq struct {
	From      []Table       `json:"from"`
	Condition LessCondsData `json:"condition"`
	Limit     int           `json:"limit"`
	Offset    int           `json:"offset"`
	Columns   []string      `json:"columns"`
	Orders    []Order       `json:"orders"`
	Groups    []string      `json:"groups"`
}

type Table struct {
	Type     string `json:"type"`
	Schema   string `json:"schema"`
	Database string `json:"database"`
}

type Order struct {
	Column string `json:"column"`
	IsAsc  bool   `json:"is_asc"`
}

type QueryResp struct {
	vo.Err
	Data []map[string]interface{} `json:"data"`
}

type LessCondsData struct {
	// 类型(between,equal,gt,gte,in,like,lt,lte,not_in,not_like,not_null,is_null,all_in,values_in)
	Type string `json:"type"`
	// 字段类型
	FieldType *string `json:"fieldType"`
	// 值
	Value interface{} `json:"value"`
	// 值（数组）
	Values interface{} `json:"values"`
	// 字段id
	Column string `json:"column"`
	// 左值
	Left interface{} `json:"left"`
	// 右值
	Right interface{} `json:"right"`
	// 嵌套
	Conds []*LessCondsData `json:"conds"`
}

type Executor struct {
	// 操作类型, 1 update, 2 delete, 3 insert
	Type      int64          `json:"type"`
	Table     *Table         `json:"table"`
	From      []*Table       `json:"from"`
	Condition *LessCondsData `json:"condition"`
	Sets      []Set          `json:"sets"`
	Columns   []string       `json:"columns"` // 插入的字段
	Values    []interface{}  `json:"values"`  // 插入的值
}

type Set struct {
	TableAlias string      `json:"tableAlias"` // 表别名
	Column     string      `json:"column"`     // 字段
	Value      interface{} `json:"value"`      // 值
	// 类型，1 普通字段，2 jsonb
	Type            int64 `json:"type"`
	WithoutPretreat bool  `json:"withoutPretreat"` // 转义
}

type ExecutorResp struct {
	vo.Err
	Data int64 `json:"data"`
}

type ExecutorsResp struct {
	vo.Err
	Data []int64 `json:"data"`
}

// CreateTableReq 创建表
type CreateTableReq struct {
	Table       string `json:"table"`
	IsSub       bool   `json:"sub"`
	SummaryFlag bool   `json:"summaryFlag"`
}

type CreateTableResp struct {
	vo.Err
}

func NewTable(tableName string) *Table {
	return &Table{Schema: tableName, Type: "schema"}
}

func WrapperJsonColumn(keyPath string) string {
	if _, ok := consts.NotJsonColumnMap[keyPath]; ok {
		return fmt.Sprintf(`"%s"`, keyPath)
	} else {
		jsonPath := "\"data\"::jsonb"
		sps := strings.Split(keyPath, ".")
		for _, sp := range sps {
			jsonPath = jsonPath + "->'" + sp + "'"
		}
		return jsonPath
	}
}

func WrapperJsonColumnAlias(keyPath, alias string) string {
	column := WrapperJsonColumn(keyPath)

	return column + " " + fmt.Sprintf(`"%s"`, alias)
}
