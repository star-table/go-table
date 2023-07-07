package consts

const (
	// 操作类型, 1 update, 2 delete, 3 insert
	ExecutorUpdate = iota + 1
	ExecutorDelete
	ExecutorInsert
)

// set值 类型
const (
	CommonFields = iota + 1
	JsonbFields
)

// 查询字段
const (
	ColumnData = "data"
	TableType = "schema"
)
