package validator

import (
	"github.com/star-table/go-table/internal/biz/bo"
	tablePb "github.com/star-table/interface/golang/table/v1"
)

var Validators map[tablePb.ColumnType]*Validator

type Validator interface {
	Validate(column *bo.Column, value interface{}) error
}
