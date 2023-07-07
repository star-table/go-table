package cache

import (
	"fmt"
	"time"
)

var (
	DefaultDuration = time.Hour * 24 * 7 // 默认缓存7天
)

var (
	prefix               = "less_code_table:"
	tableSchemasKey      = prefix + "table_schema:%d"
	columnDescriptionKey = prefix + "column_description:%d"
	appTablesKey         = prefix + "app_table:%d" // app对应的table
	orgColumnsKey        = prefix + "org_columns:%d"
	summeryTableKey      = prefix + "summery_table:%d"
)

func GetTableSchemasKey(tableId int64) string {
	return fmt.Sprintf(tableSchemasKey, tableId)
}

func GetAppTablesKey(appId int64) string {
	return fmt.Sprintf(appTablesKey, appId)
}

func GetColumnDescriptionKey(tableId int64) string {
	return fmt.Sprintf(columnDescriptionKey, tableId)
}

func GetOrgColumnsKey(orgId int64) string {
	return fmt.Sprintf(orgColumnsKey, orgId)
}

func GetSummeryTableKey(orgId int64) string {
	return fmt.Sprintf(summeryTableKey, orgId)
}
