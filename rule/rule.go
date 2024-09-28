package rule

import (
	"github.com/fuckqqcom/pkg/convert"
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

type Rule struct {
	Skip bool

	Key string
	Op  Op
	Val any
}

func New(rules ...Rule) []Rule {
	return rules
}

func Select(builder sqlbuilder.SelectBuilder, rules ...Rule) sqlbuilder.SelectBuilder {
	for _, rule := range rules {
		if rule.Skip {
			continue
		}
		switch Op(strings.ToUpper(string(rule.Op))) {
		case E:
			builder.Where(builder.Equal(rule.Key, rule.Val))
		case NE:
			builder.Where(builder.NotEqual(rule.Key, rule.Val))
		case GT:
			builder.Where(builder.GreaterThan(rule.Key, rule.Val))
		case LT:
			builder.Where(builder.LessThan(rule.Key, rule.Val))
		case GTE:
			builder.Where(builder.GreaterEqualThan(rule.Key, rule.Val))
		case LTE:
			builder.Where(builder.LessEqualThan(rule.Key, rule.Val))
		case In:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.Where(builder.In(rule.Key, convert.ReflectSlice(rule.Val)...))
			}
		case NotIn:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.Where(builder.NotIn(rule.Key, convert.ReflectSlice(rule.Val)...))
			}
		case Like:
			builder.Where(builder.Like(rule.Key, rule.Val))
		case NotLike:
			builder.Where(builder.NotLike(rule.Key, rule.Val))
		case Limit:
			builder.Limit(cast.ToInt(rule.Val))
		case Offset:
			builder.Offset(cast.ToInt(rule.Val))
		case Between:
			value := convert.ReflectSlice(rule.Val)
			if len(value) == 2 {
				builder.Where(builder.Between(rule.Key, value[0], value[1]))
			}
		case OrderBy:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convert.ReflectSlice(rule.Val))...)
			}
		}
	}
	return builder
}

func Update(builder sqlbuilder.UpdateBuilder, rules ...Rule) sqlbuilder.UpdateBuilder {
	for _, rule := range rules {
		if rule.Skip {
			continue
		}
		switch Op(strings.ToUpper(string(rule.Op))) {
		case E:
			builder.Where(builder.Equal(rule.Key, rule.Val))
		case NE:
			builder.Where(builder.NotEqual(rule.Key, rule.Val))
		case GT:
			builder.Where(builder.GreaterThan(rule.Key, rule.Val))
		case LT:
			builder.Where(builder.LessThan(rule.Key, rule.Val))
		case GTE:
			builder.Where(builder.GreaterEqualThan(rule.Key, rule.Val))
		case LTE:
			builder.Where(builder.LessEqualThan(rule.Key, rule.Val))
		case In:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.Where(builder.In(rule.Key, convert.ReflectSlice(rule.Val)...))
			}
		case NotIn:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.Where(builder.NotIn(rule.Key, convert.ReflectSlice(rule.Val)...))
			}
		case Like:
			builder.Where(builder.Like(rule.Key, rule.Val))
		case NotLike:
			builder.Where(builder.NotLike(rule.Key, rule.Val))
		case Limit:
			builder.Limit(cast.ToInt(rule.Val))
		case Between:
			value := convert.ReflectSlice(rule.Val)
			if len(value) == 2 {
				builder.Where(builder.Between(rule.Key, value[0], value[1]))
			}
		case OrderBy:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convert.ReflectSlice(rule.Val))...)
			}
		}
	}
	return builder
}

func Delete(builder sqlbuilder.DeleteBuilder, rules ...Rule) sqlbuilder.DeleteBuilder {
	for _, rule := range rules {
		if rule.Skip {
			continue
		}
		switch Op(strings.ToUpper(string(rule.Op))) {
		case E:
			builder.Where(builder.Equal(rule.Key, rule.Val))
		case NE:
			builder.Where(builder.NotEqual(rule.Key, rule.Val))
		case GT:
			builder.Where(builder.GreaterThan(rule.Key, rule.Val))
		case LT:
			builder.Where(builder.LessThan(rule.Key, rule.Val))
		case GTE:
			builder.Where(builder.GreaterEqualThan(rule.Key, rule.Val))
		case LTE:
			builder.Where(builder.LessEqualThan(rule.Key, rule.Val))
		case In:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.Where(builder.In(rule.Key, convert.ReflectSlice(rule.Val)...))
			}
		case NotIn:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.Where(builder.NotIn(rule.Key, convert.ReflectSlice(rule.Val)...))
			}
		case Like:
			builder.Where(builder.Like(rule.Key, rule.Val))
		case NotLike:
			builder.Where(builder.NotLike(rule.Key, rule.Val))
		case Limit:
			builder.Limit(cast.ToInt(rule.Val))
		case Between:
			value := convert.ReflectSlice(rule.Val)
			if len(value) == 2 {
				builder.Where(builder.Between(rule.Key, value[0], value[1]))
			}
		case OrderBy:
			if len(convert.ReflectSlice(rule.Val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convert.ReflectSlice(rule.Val))...)
			}
		}
	}
	return builder
}
