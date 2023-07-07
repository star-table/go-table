package validator

import (
	"strconv"

	"github.com/star-table/go-common/pkg/errors"
	"github.com/star-table/go-table/internal/biz/bo"
	commonPb "github.com/star-table/interface/golang/common/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

type AbstractValidator struct{}

// 检查：不可写字段
func (v *AbstractValidator) validateWritable(column *bo.Column, value *structpb.Value) error {
	if !column.Schema.Writable && value != nil {
		return errors.Ignore(commonPb.ErrorRowDataValidateFailed("%s字段不允许写入", column.Schema.Label))
	}
	return nil
}

// 检查：required字段
func (v *AbstractValidator) validateRequired(column *bo.Column, value *structpb.Value) error {
	var props *structpb.Struct
	if column.Schema.GetField() != nil && column.Schema.GetField().GetProps() != nil {
		props = column.Schema.GetField().GetProps()
	}
	if props != nil && len(props.GetFields()) != 0 {
		if r, ok := props.GetFields()["required"]; ok && r.GetBoolValue() == true {
			if err := v._validateRequired(column, value); err != nil {
				return err
			}
		} else if r, ok := props.GetFields()["require"]; ok && r.GetBoolValue() == true {
			if err := v._validateRequired(column, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func (v *AbstractValidator) _validateRequired(column *bo.Column, value *structpb.Value) error {
	e := commonPb.ErrorRowDataValidateFailed("%s字段必填", column.Schema.Label)
	if value == nil {
		return errors.Ignore(e)
	}
	switch value.Kind.(type) {
	case *structpb.Value_StringValue:
		if value.GetStringValue() == "" {
			return errors.Ignore(e)
		}
	case *structpb.Value_ListValue:
		if len(value.GetListValue().Values) == 0 {
			return errors.Ignore(e)
		}
	case *structpb.Value_StructValue:
		if len(value.GetStructValue().Fields) == 0 {
			return errors.Ignore(e)
		}
	}
	return nil
}

// 检查：数字类型
func (v *AbstractValidator) validateNumber(column *bo.Column, value *structpb.Value) error {
	e := commonPb.ErrorRowDataValidateFailed("%s字段必须为数字", column.Schema.Label)
	switch value.Kind.(type) {
	case *structpb.Value_NumberValue:
		return nil
	case *structpb.Value_StringValue:
		if _, err := strconv.ParseFloat(value.GetStringValue(), 64); err == nil {
			return nil
		} else {
			return errors.Ignore(e)
		}
	default:
		return errors.Ignore(e)
	}
	return nil
}
