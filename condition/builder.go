package condition

import (
	"github.com/fuckqqcom/pkg/convertx"
	"github.com/huandu/go-sqlbuilder"
	"github.com/spf13/cast"
	"strings"
)

type Operator string

func (o Operator) String() string {
	return string(o)
}

const (
	Equal            Operator = "="
	E                Operator = "="
	NotEqual         Operator = "!="
	NE               Operator = "!="
	GreaterThan      Operator = ">"
	GT               Operator = ">"
	LessThan         Operator = "<"
	LT               Operator = "<"
	GreaterEqualThan Operator = ">="
	GTE              Operator = ">="
	LessEqualThan    Operator = "<="
	LTE              Operator = "<="
	In               Operator = "IN"
	NotIn            Operator = "NOT IN"
	Like             Operator = "LIKE"
	NotLike          Operator = "NOT LIKE"
	Limit            Operator = "LIMIT"
	Offset           Operator = "OFFSET"
	Between          Operator = "BETWEEN"
	OrderBy          Operator = "ORDER BY"
)

type Condition struct {
	Skip bool

	Field    string
	Operator Operator
	Value    any
}

func New(conditions ...Condition) []Condition {
	return conditions
}

func Select(builder sqlbuilder.SelectBuilder, conditions ...Condition) sqlbuilder.SelectBuilder {
	for _, cond := range conditions {
		if cond.Skip {
			continue
		}
		switch Operator(strings.ToUpper(string(cond.Operator))) {
		case E:
			builder.Where(builder.Equal(cond.Field, cond.Value))
		case NE:
			builder.Where(builder.NotEqual(cond.Field, cond.Value))
		case GT:
			builder.Where(builder.GreaterThan(cond.Field, cond.Value))
		case LT:
			builder.Where(builder.LessThan(cond.Field, cond.Value))
		case GTE:
			builder.Where(builder.GreaterEqualThan(cond.Field, cond.Value))
		case LTE:
			builder.Where(builder.LessEqualThan(cond.Field, cond.Value))
		case In:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.Where(builder.In(cond.Field, convertx.ReflectSlice(cond.Value)...))
			}
		case NotIn:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.Where(builder.NotIn(cond.Field, convertx.ReflectSlice(cond.Value)...))
			}
		case Like:
			builder.Where(builder.Like(cond.Field, cond.Value))
		case NotLike:
			builder.Where(builder.NotLike(cond.Field, cond.Value))
		case Limit:
			builder.Limit(cast.ToInt(cond.Value))
		case Offset:
			builder.Offset(cast.ToInt(cond.Value))
		case Between:
			value := convertx.ReflectSlice(cond.Value)
			if len(value) == 2 {
				builder.Where(builder.Between(cond.Field, value[0], value[1]))
			}
		case OrderBy:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convertx.ReflectSlice(cond.Value))...)
			}
		}
	}
	return builder
}

func Update(builder sqlbuilder.UpdateBuilder, conditions ...Condition) sqlbuilder.UpdateBuilder {
	for _, cond := range conditions {
		if cond.Skip {
			continue
		}
		switch Operator(strings.ToUpper(string(cond.Operator))) {
		case E:
			builder.Where(builder.Equal(cond.Field, cond.Value))
		case NE:
			builder.Where(builder.NotEqual(cond.Field, cond.Value))
		case GT:
			builder.Where(builder.GreaterThan(cond.Field, cond.Value))
		case LT:
			builder.Where(builder.LessThan(cond.Field, cond.Value))
		case GTE:
			builder.Where(builder.GreaterEqualThan(cond.Field, cond.Value))
		case LTE:
			builder.Where(builder.LessEqualThan(cond.Field, cond.Value))
		case In:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.Where(builder.In(cond.Field, convertx.ReflectSlice(cond.Value)...))
			}
		case NotIn:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.Where(builder.NotIn(cond.Field, convertx.ReflectSlice(cond.Value)...))
			}
		case Like:
			builder.Where(builder.Like(cond.Field, cond.Value))
		case NotLike:
			builder.Where(builder.NotLike(cond.Field, cond.Value))
		case Limit:
			builder.Limit(cast.ToInt(cond.Value))
		case Between:
			value := convertx.ReflectSlice(cond.Value)
			if len(value) == 2 {
				builder.Where(builder.Between(cond.Field, value[0], value[1]))
			}
		case OrderBy:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convertx.ReflectSlice(cond.Value))...)
			}
		}
	}
	return builder
}

func Delete(builder sqlbuilder.DeleteBuilder, conditions ...Condition) sqlbuilder.DeleteBuilder {
	for _, cond := range conditions {
		if cond.Skip {
			continue
		}
		switch Operator(strings.ToUpper(string(cond.Operator))) {
		case E:
			builder.Where(builder.Equal(cond.Field, cond.Value))
		case NE:
			builder.Where(builder.NotEqual(cond.Field, cond.Value))
		case GT:
			builder.Where(builder.GreaterThan(cond.Field, cond.Value))
		case LT:
			builder.Where(builder.LessThan(cond.Field, cond.Value))
		case GTE:
			builder.Where(builder.GreaterEqualThan(cond.Field, cond.Value))
		case LTE:
			builder.Where(builder.LessEqualThan(cond.Field, cond.Value))
		case In:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.Where(builder.In(cond.Field, convertx.ReflectSlice(cond.Value)...))
			}
		case NotIn:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.Where(builder.NotIn(cond.Field, convertx.ReflectSlice(cond.Value)...))
			}
		case Like:
			builder.Where(builder.Like(cond.Field, cond.Value))
		case NotLike:
			builder.Where(builder.NotLike(cond.Field, cond.Value))
		case Limit:
			builder.Limit(cast.ToInt(cond.Value))
		case Between:
			value := convertx.ReflectSlice(cond.Value)
			if len(value) == 2 {
				builder.Where(builder.Between(cond.Field, value[0], value[1]))
			}
		case OrderBy:
			if len(convertx.ReflectSlice(cond.Value)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convertx.ReflectSlice(cond.Value))...)
			}
		}
	}
	return builder
}
