package rule

import (
	"fmt"
	"github.com/fuckqqcom/pkg/convert"
	"github.com/huandu/go-sqlbuilder"
	"github.com/spf13/cast"
)

/*
	照搬 https://github.com/jzero-io/jzero-contrib/tree/main/condition

*/

type Op string

func (o Op) String() string {
	return string(o)
}

const (
	E         Op = "="
	NE        Op = "!="
	GT        Op = ">"
	LT        Op = "<"
	GTE       Op = ">="
	LTE       Op = "<="
	In        Op = "IN"
	NotIn     Op = "NOT IN"
	Like      Op = "LIKE"
	NotLike   Op = "NOT LIKE"
	Limit     Op = "LIMIT"
	Offset    Op = "OFFSET"
	Between   Op = "BETWEEN"
	OrderBy   Op = "ORDER BY"
	FindInSet Op = "FIND_IN_SET"
)

type Rule struct {
	Key string

	skip     bool
	SkipFunc func() bool

	// Or condition
	Or         bool
	OrOps      []Op
	OrKeys     []string
	orVals     []any
	OrValsFunc func() []any

	// And condition
	Op      Op
	val     any
	ValFunc func() any
}

func NewRule(rules ...Rule) []Rule {
	return rules
}

func buildExpr(cond *sqlbuilder.Cond, key string, operator Op, value any) string {
	switch operator {
	case E:
		return cond.Equal(key, value)
	case NE:
		return cond.NotEqual(key, value)
	case GT:
		return cond.GreaterThan(key, value)
	case LT:
		return cond.LessThan(key, value)
	case GTE:
		return cond.GreaterEqualThan(key, value)
	case LTE:
		return cond.LessEqualThan(key, value)
	case In:
		if len(convert.ReflectSlice(value)) > 0 {
			return cond.In(key, convert.ReflectSlice(value)...)
		}
	case NotIn:
		if len(convert.ReflectSlice(value)) > 0 {
			return cond.NotIn(key, convert.ReflectSlice(value)...)
		}
	case Like:
		return cond.Like(key, value)
	case NotLike:
		return cond.NotLike(key, value)
	case Between:
		values := convert.ReflectSlice(value)
		if len(values) == 2 {
			return cond.Between(key, values[0], values[1])
		}
	}
	return ""
}

func whereClause(rules ...Rule) *sqlbuilder.WhereClause {
	clause := sqlbuilder.NewWhereClause()
	cond := sqlbuilder.NewCond()

	for _, r := range rules {
		// Skip logic
		if r.SkipFunc != nil {
			r.skip = r.SkipFunc()
		}
		if r.skip {
			continue
		}

		// OR condition handling
		if r.Or {
			if r.OrValsFunc != nil {
				r.orVals = r.OrValsFunc()
			}
			var expr []string
			for i, key := range r.OrKeys {
				if or := buildExpr(cond, key, r.OrOps[i], r.orVals[i]); or != "" {
					expr = append(expr, or)
				}
			}
			if len(expr) > 0 {
				clause.AddWhereExpr(cond.Args, cond.Or(expr...))
			}
		} else {
			// Non-OR condition handling
			if r.ValFunc != nil {
				r.val = r.ValFunc()
			}
			if expr := buildExpr(cond, r.Key, r.Op, r.val); expr != "" {
				clause.AddWhereExpr(cond.Args, expr)
			}
		}
	}
	return clause
}

func Select(builder sqlbuilder.SelectBuilder, rules ...Rule) sqlbuilder.SelectBuilder {
	clause := whereClause(rules...)
	for _, r := range rules {
		if r.SkipFunc != nil {
			r.skip = r.SkipFunc()
		}
		if r.skip {
			continue
		}
		if r.ValFunc != nil {
			r.val = r.ValFunc()
		}
		switch r.Op {
		case Limit:
			builder.Limit(cast.ToInt(r.val))
		case Offset:
			builder.Offset(cast.ToInt(r.val))
		case OrderBy:
			if len(convert.ReflectSlice(r.val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convert.ReflectSlice(r.val))...)
			}
		case FindInSet:
			builder.Where(fmt.Sprintf("FIND_IN_SET(%s, %s)", builder.Var(r.val), builder.Var(r.Key)))
		}
	}

	if clause != nil {
		builder = *builder.AddWhereClause(clause)
	}

	return builder
}

func Update(builder sqlbuilder.UpdateBuilder, rules ...Rule) sqlbuilder.UpdateBuilder {
	clause := whereClause(rules...)
	for _, r := range rules {
		if r.SkipFunc != nil {
			r.skip = r.SkipFunc()
		}
		if r.skip {
			continue
		}
		if r.ValFunc != nil {
			r.val = r.ValFunc()
		}
		switch r.Op {
		case Limit:
			builder.Limit(cast.ToInt(r.val))
		case OrderBy:
			if len(convert.ReflectSlice(r.val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convert.ReflectSlice(r.val))...)
			}
		}
	}
	if clause != nil {
		builder = *builder.AddWhereClause(clause)
	}

	return builder
}

func Delete(builder sqlbuilder.DeleteBuilder, rules ...Rule) sqlbuilder.DeleteBuilder {
	clause := whereClause(rules...)
	for _, r := range rules {
		if r.SkipFunc != nil {
			r.skip = r.SkipFunc()
		}
		if r.skip {
			continue
		}
		if r.ValFunc != nil {
			r.val = r.ValFunc()
		}
		switch r.Op {
		case Limit:
			builder.Limit(cast.ToInt(r.val))
		case OrderBy:
			if len(convert.ReflectSlice(r.val)) > 0 {
				builder.OrderBy(cast.ToStringSlice(convert.ReflectSlice(r.val))...)
			}
		}
	}
	if clause != nil {
		builder = *builder.AddWhereClause(clause)
	}
	return builder
}
