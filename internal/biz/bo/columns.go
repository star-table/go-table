package bo

import (
	pb "github.com/star-table/interface/golang/table/v1"
)

type Column struct {
	TableId    int64
	ColumnId   string
	ColumnType string
	Schema     *pb.Column
}

type Columns struct {
	AppId   int64
	TableId int64
	Name    string
	Columns []*pb.Column
}

type GroupSelectOption struct {
	Id       int64 // 当前选项的id
	ParentId int64 // 父类
}

type ReferenceSetting struct {
	RelateColumnId string // 引用的关联列
	AggFunc        string // 处理方法
}

type ReferenceColumnInfo struct {
	RelateColumnId    string // 关联列
	ReferenceColumnId string // 真正引用的数据列
	OriginColumnId    string // 原始的引用列
	AggFunc           string // 处理函数
	IsRelate          bool   // 是否是关联列
}

type RelateColumnInfo struct {
	Name  string
	AppId int64
}
