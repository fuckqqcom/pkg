package condition

import (
	"github.com/fuckqqcom/pkg/convertx"
	"github.com/huandu/go-sqlbuilder"
	"github.com/spf13/cast"
	"strings"
)

type Op string

func (o Op) String() string {
	return string(o)
}

const (
	Equal            Op = "="
	E                Op = "="
	NotEqual         Op = "!="
	NE               Op = "!="
	GreaterThan      Op = ">"
	GT               Op = ">"
	LessThan         Op = "<"
	LT               Op = "<"
	GreaterEqualThan Op = ">="
	GTE              Op = ">="
	LessEqualThan    Op = "<="
	LTE              Op = "<="
	In               Op = "IN"
	NotIn            Op = "NOT IN"
	Like             Op = "LIKE"
	NotLike          Op = "NOT LIKE"
	Limit            Op = "LIMIT"
	Offset           Op = "OFFSET"
	Between          Op = "BETWEEN"
	OrderBy          Op = "ORDER BY"
)

type Condition struct {
	Skip bool

	Key string
	Op  Op
	Val any
}

func New(conditions ...Condition) []Condition {
	return conditions
}

func Select(builder sqlbuilder.SelectBuilder, conditions ...Condition) sqlbuilder.SelectBuilder {
	for _, cond := range conditions {
		if cond.Skip {
			continue
		}
		switch Op(strings.ToUpper(string(cond.Op))) {
		case E:
			builder.Where(builder.Equal(cond.Key, cond.Val))
		case NE:
			builder.Where(builder.NotEqual(cond.Key, cond.Val))
		case GT:
			builder.Where(builder.GreaterThan(cond.Key, cond.Val))
		case LT:
			builder.Where(builder.LessThan(cond.Key, cond.Val))
		case GTE:
			builder.Where(builder.GreaterEqualThan(cond.Key, cond.Val))
		case LTE:
			builder.Where(builder.LessEqualThan(cond.Key, cond.Val))
		case In:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.Where(builder.In(cond.Key, convertx.ReflectSlice(cond.Val)...))
			}
		case NotIn:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.Where(builder.NotIn(cond.Key, convertx.ReflectSlice(cond.Val)...))
			}
		case Like:
			builder.Where(builder.Like(cond.Key, cond.Val))
		case NotLike:
			builder.Where(builder.NotLike(cond.Key, cond.Val))
		case Limit:
			builder.Limit(cast.ToInt(cond.Val))
		case Offset:
			builder.Offset(cast.ToInt(cond.Val))
		case Between:
			value := convertx.ReflectSlice(cond.Val)
			if len(value) == 2 {
				builder.Where(builder.Between(cond.Key, value[0], value[1]))
			}
		case OrderBy:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convertx.ReflectSlice(cond.Val))...)
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
		switch Op(strings.ToUpper(string(cond.Op))) {
		case E:
			builder.Where(builder.Equal(cond.Key, cond.Val))
		case NE:
			builder.Where(builder.NotEqual(cond.Key, cond.Val))
		case GT:
			builder.Where(builder.GreaterThan(cond.Key, cond.Val))
		case LT:
			builder.Where(builder.LessThan(cond.Key, cond.Val))
		case GTE:
			builder.Where(builder.GreaterEqualThan(cond.Key, cond.Val))
		case LTE:
			builder.Where(builder.LessEqualThan(cond.Key, cond.Val))
		case In:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.Where(builder.In(cond.Key, convertx.ReflectSlice(cond.Val)...))
			}
		case NotIn:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.Where(builder.NotIn(cond.Key, convertx.ReflectSlice(cond.Val)...))
			}
		case Like:
			builder.Where(builder.Like(cond.Key, cond.Val))
		case NotLike:
			builder.Where(builder.NotLike(cond.Key, cond.Val))
		case Limit:
			builder.Limit(cast.ToInt(cond.Val))
		case Between:
			value := convertx.ReflectSlice(cond.Val)
			if len(value) == 2 {
				builder.Where(builder.Between(cond.Key, value[0], value[1]))
			}
		case OrderBy:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convertx.ReflectSlice(cond.Val))...)
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
		switch Op(strings.ToUpper(string(cond.Op))) {
		case E:
			builder.Where(builder.Equal(cond.Key, cond.Val))
		case NE:
			builder.Where(builder.NotEqual(cond.Key, cond.Val))
		case GT:
			builder.Where(builder.GreaterThan(cond.Key, cond.Val))
		case LT:
			builder.Where(builder.LessThan(cond.Key, cond.Val))
		case GTE:
			builder.Where(builder.GreaterEqualThan(cond.Key, cond.Val))
		case LTE:
			builder.Where(builder.LessEqualThan(cond.Key, cond.Val))
		case In:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.Where(builder.In(cond.Key, convertx.ReflectSlice(cond.Val)...))
			}
		case NotIn:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.Where(builder.NotIn(cond.Key, convertx.ReflectSlice(cond.Val)...))
			}
		case Like:
			builder.Where(builder.Like(cond.Key, cond.Val))
		case NotLike:
			builder.Where(builder.NotLike(cond.Key, cond.Val))
		case Limit:
			builder.Limit(cast.ToInt(cond.Val))
		case Between:
			value := convertx.ReflectSlice(cond.Val)
			if len(value) == 2 {
				builder.Where(builder.Between(cond.Key, value[0], value[1]))
			}
		case OrderBy:
			if len(convertx.ReflectSlice(cond.Val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convertx.ReflectSlice(cond.Val))...)
			}
		}
	}
	return builder
}
