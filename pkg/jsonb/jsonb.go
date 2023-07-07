package jsonb

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"

	"go.uber.org/zap/buffer"

	"github.com/spf13/cast"

	tablePb "github.com/star-table/interface/golang/table/v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewQuery(condition *tablePb.Condition) *QueryExpression {
	return &QueryExpression{
		Condition: condition,
	}
}

type QueryExpression struct {
	Condition *tablePb.Condition
}

func (query *QueryExpression) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		query.toSql(builder, stmt, query.Condition)
	}
}

func (query *QueryExpression) toSql(builder clause.Builder, stmt *gorm.Statement, condition *tablePb.Condition) {
	values, err := query.changeValue(condition.Column, condition.Value, condition.Type)
	if err != nil {
		return
	}

	switch condition.Type {
	case tablePb.ConditionType_equal:
		if len(values) == 0 {
			_, _ = builder.WriteString(" 1 = 2 ")
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s = ", condition.Column))
		stmt.AddVar(builder, values[0])
	case tablePb.ConditionType_un_equal:
		if len(values) == 0 {
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s != ", condition.Column))
		stmt.AddVar(builder, values[0])
	case tablePb.ConditionType_in:
		if len(values) == 0 {
			_, _ = builder.WriteString(" 1 = 2 ")
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s in ", condition.Column))
		stmt.AddVar(builder, values)
	case tablePb.ConditionType_not_in:
		if len(values) == 0 {
			_, _ = builder.WriteString(" 1 = 2 ")
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s not in ", condition.Column))
		stmt.AddVar(builder, values)
	case tablePb.ConditionType_not_null:
		_, _ = builder.WriteString(fmt.Sprintf(` %s is not null and %s != '[]' and %s != '""' `, condition.Column, condition.Column, condition.Column))
	case tablePb.ConditionType_is_null:
		_, _ = builder.WriteString(fmt.Sprintf(` %s is null or %s = '[]' or %s = '""' `, condition.Column, condition.Column, condition.Column))
	case tablePb.ConditionType_values_in:
		_, _ = builder.WriteString(fmt.Sprintf("ARRAY(SELECT jsonb_array_elements_text(%s)) && ARRAY[%s]", query.getCoalesceColumnArray(condition.Column), strings.Join(cast.ToStringSlice(values), ",")))
	case tablePb.ConditionType_gt:
		if len(values) == 0 {
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s > ", condition.Column))
		stmt.AddVar(builder, values[0])
	case tablePb.ConditionType_gte:
		if len(values) == 0 {
			_, _ = builder.WriteString(" 1 = 2 ")
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s >= ", condition.Column))
		stmt.AddVar(builder, values[0])
	case tablePb.ConditionType_lt:
		if len(values) == 0 {
			_, _ = builder.WriteString(" 1 = 2 ")
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s < ", condition.Column))
		stmt.AddVar(builder, values[0])
	case tablePb.ConditionType_lte:
		if len(values) == 0 {
			_, _ = builder.WriteString(" 1 = 2 ")
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s <= ", condition.Column))
		stmt.AddVar(builder, values[0])
	case tablePb.ConditionType_between:
		if len(values) != 2 {
			_, _ = builder.WriteString(" 1 = 2 ")
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s is not null ", condition.Column))
		_, _ = builder.WriteString(fmt.Sprintf(" and  %s >= ", condition.Column))
		stmt.AddVar(builder, values[0])
		_, _ = builder.WriteString(fmt.Sprintf(" and %s <= ", condition.Column))
		stmt.AddVar(builder, values[1])
	case tablePb.ConditionType_like:
		if len(values) == 0 {
			_, _ = builder.WriteString(" 1 = 2 ")
			return
		}
		_, _ = builder.WriteString(fmt.Sprintf(" %s like ", condition.Column))
		stmt.AddVar(builder, "%"+cast.ToString(values[0])+"%")
	case tablePb.ConditionType_or:
		query.toSqls(builder, stmt, condition.Conditions, "or")
	case tablePb.ConditionType_and:
		query.toSqls(builder, stmt, condition.Conditions, "and")
	case tablePb.ConditionType_raw_sql:
		_, _ = builder.WriteString(condition.Column)
	default:
		log.Errorf("[toSql] unsupported type:%v", condition.Type.String())
	}
}

func (query *QueryExpression) writeInOptimize(condition *tablePb.Condition, builder clause.Builder, values []interface{}) {
	if len(values) > 1 {
		_, _ = builder.WriteString(" ( ")
	}

	jsonFmt := query.getOptimizeJson([]*tablePb.Condition{condition})
	for i, value := range values {
		_, _ = builder.WriteString(fmt.Sprintf(jsonFmt, cast.ToString(value)))
		if i != len(values)-1 {
			_, _ = builder.WriteString(" or ")
		}
	}

	if len(values) > 1 {
		_, _ = builder.WriteString(" ) ")
	}
}

// getOptimizeJson fmt.Sprintf("\"data\" :: jsonb -> '%s' \"%s\"", c, c)
// 同一个and内的equal做json合并优化
func (query *QueryExpression) getOptimizeJson(conditions []*tablePb.Condition) string {
	bf := buffer.Buffer{}
	_, _ = bf.WriteString(" data @> '{")
	for i, condition := range conditions {
		sp := strings.Split(condition.Column, "'")
		_, _ = bf.WriteString("\"" + sp[1] + "\"")
		// 最多支持两层 ->aaa->bbb的优化
		if len(sp) >= 3 && sp[2] != "" {
			_, _ = bf.WriteString(":{\"" + sp[2] + "\":%s}")
		} else {
			_, _ = bf.WriteString(":%s")
		}

		if i != len(conditions)-1 {
			_, _ = bf.WriteString(", ")
		}
	}

	_, _ = bf.WriteString("}' ")

	return bf.String()
}

func (query *QueryExpression) checkIsIndexOptimize(column string) bool {
	if strings.Contains(column, "->") {
		return true
	}

	return false
}

func (query *QueryExpression) toSqls(builder clause.Builder, stmt *gorm.Statement, conditions []*tablePb.Condition, conditionType string) {
	if len(conditions) == 0 {
		return
	}

	_, _ = builder.WriteString(" ( ")
	for i, condition := range conditions {
		query.toSql(builder, stmt, condition)
		if i != len(conditions)-1 {
			_, _ = builder.WriteString(" " + conditionType + " ")
		}
	}
	_, _ = builder.WriteString(" ) ")
}

func (query *QueryExpression) toAndSqls(builder clause.Builder, stmt *gorm.Statement, conditions []*tablePb.Condition) {
	if len(conditions) == 0 {
		return
	}

	otherConditions := make([]*tablePb.Condition, 0, len(conditions)/2)
	equalConditions := make([]*tablePb.Condition, 0, len(conditions)/2)
	equalValues := make([]interface{}, 0, len(conditions)/2)
	for _, condition := range conditions {
		if condition.Type == tablePb.ConditionType_equal && query.checkIsIndexOptimize(condition.Column) {
			equalConditions = append(equalConditions, condition)
			values, _ := query.changeValue(condition.Column, condition.Value, condition.Type)
			equalValues = append(equalValues, values...)
		} else {
			otherConditions = append(otherConditions, condition)
		}
	}

	if len(equalConditions) > 0 {
		jsonFmt := query.getOptimizeJson(equalConditions)
		_, _ = builder.WriteString(fmt.Sprintf(jsonFmt, equalValues...))
	}

	if len(otherConditions) > 0 {
		_, _ = builder.WriteString(" and ")
		query.toSqls(builder, stmt, otherConditions, "and")
	}
}

func (query *QueryExpression) changeValue(column, value string, conditionType tablePb.ConditionType) ([]interface{}, error) {
	if value == "" {
		return nil, nil
	}

	var list []interface{}
	dec := json.NewDecoder(strings.NewReader(value))
	dec.UseNumber()
	err := dec.Decode(&list)
	for i := range list {
		switch conditionType {
		case tablePb.ConditionType_values_in:
			list[i] = "'" + cast.ToString(list[i]) + "'"
		default:
			// 只有json的列才加上双引号
			if s, ok := list[i].(string); ok && strings.Contains(column, "->") {
				list[i] = "\"" + s + "\""
			}
		}
	}

	return list, err
}

func (query *QueryExpression) getCoalesceColumnArray(column string) string {
	if !strings.Contains(column, "->") {
		return column
	}
	columnText := strings.Replace(column, "->", "->>", 1)

	return fmt.Sprintf("case when %s is null then '[]'::jsonb else %s end", columnText, column)
}
