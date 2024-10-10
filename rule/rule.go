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
	Skip     bool
	SkipFunc func() bool

	//or condition
	Or         bool
	OrOps      []Op
	OrKeys     []string
	OrVals     []any
	OrValsFunc func() []any

	// and condition
	Key     string
	Op      Op
	Val     any
	ValFunc func() any
}

func New(rules ...Rule) []Rule {
	return rules
}

func ApplySelect(builder *sqlbuilder.SelectBuilder, rules ...Rule) {
	for _, rule := range rules {
		if rule.SkipFunc != nil {
			rule.Skip = rule.SkipFunc()
		}
		if rule.Skip {
			continue
		}
		if rule.Or {
			if rule.OrValsFunc != nil {
				rule.OrVals = rule.OrValsFunc()
			}
			var expr []string
			for i, field := range rule.OrKeys {
				switch Op(strings.ToUpper(string(rule.OrOps[i]))) {
				case Equal:
					expr = append(expr, builder.Equal(field, rule.OrVals[i]))
				case NotEqual:
					expr = append(expr, builder.NotEqual(field, rule.OrVals[i]))
				case GreaterThan:
					expr = append(expr, builder.GreaterThan(field, rule.OrVals[i]))
				case LessThan:
					expr = append(expr, builder.LessThan(field, rule.OrVals[i]))
				case GreaterEqualThan:
					expr = append(expr, builder.GreaterEqualThan(field, rule.OrVals[i]))
				case LessEqualThan:
					expr = append(expr, builder.LessEqualThan(field, rule.OrVals[i]))
				case In:
					if len(convert.ReflectSlice(rule.OrVals[i])) > 0 {
						expr = append(expr, builder.In(field, convert.ReflectSlice(rule.OrVals[i])...))
					}
				case NotIn:
					if len(convert.ReflectSlice(rule.OrVals[i])) > 0 {
						expr = append(expr, builder.NotIn(field, convert.ReflectSlice(rule.OrVals[i])...))
					}
				case Like:
					expr = append(expr, builder.Like(field, rule.OrVals[i]))
				case NotLike:
					expr = append(expr, builder.NotLike(field, rule.OrVals[i]))
				case Between:
					value := convert.ReflectSlice(rule.OrVals[i])
					if len(value) == 2 {
						expr = append(expr, builder.Between(field, value[0], value[1]))
					}
				}
			}
			builder.Where(builder.Or(expr...))
		} else {
			if rule.ValFunc != nil {
				rule.Val = rule.ValFunc()
			}
			switch Op(strings.ToUpper(string(rule.Op))) {
			case Equal:
				builder.Where(builder.Equal(rule.Key, rule.Val))
			case NotEqual:
				builder.Where(builder.NotEqual(rule.Key, rule.Val))
			case GreaterThan:
				builder.Where(builder.GreaterThan(rule.Key, rule.Val))
			case LessThan:
				builder.Where(builder.LessThan(rule.Key, rule.Val))
			case GreaterEqualThan:
				builder.Where(builder.GreaterEqualThan(rule.Key, rule.Val))
			case LessEqualThan:
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
	}
}

func ApplyUpdate(builder *sqlbuilder.UpdateBuilder, rules ...Rule) {
	for _, rule := range rules {
		if rule.SkipFunc != nil {
			rule.Skip = rule.SkipFunc()
		}
		if rule.Skip {
			continue
		}
		if rule.Or {
			if rule.OrValsFunc != nil {
				rule.OrVals = rule.OrValsFunc()
			}
			var expr []string
			for i, field := range rule.OrKeys {
				switch Op(strings.ToUpper(string(rule.OrOps[i]))) {
				case Equal:
					expr = append(expr, builder.Equal(field, rule.OrVals[i]))
				case NotEqual:
					expr = append(expr, builder.NotEqual(field, rule.OrVals[i]))
				case GreaterThan:
					expr = append(expr, builder.GreaterThan(field, rule.OrVals[i]))
				case LessThan:
					expr = append(expr, builder.LessThan(field, rule.OrVals[i]))
				case GreaterEqualThan:
					expr = append(expr, builder.GreaterEqualThan(field, rule.OrVals[i]))
				case LessEqualThan:
					expr = append(expr, builder.LessEqualThan(field, rule.OrVals[i]))
				case In:
					if len(convert.ReflectSlice(rule.OrVals[i])) > 0 {
						expr = append(expr, builder.In(field, convert.ReflectSlice(rule.OrVals[i])...))
					}
				case NotIn:
					if len(convert.ReflectSlice(rule.OrVals[i])) > 0 {
						expr = append(expr, builder.NotIn(field, convert.ReflectSlice(rule.OrVals[i])...))
					}
				case Like:
					expr = append(expr, builder.Like(field, rule.OrVals[i]))
				case NotLike:
					expr = append(expr, builder.NotLike(field, rule.OrVals[i]))
				case Between:
					value := convert.ReflectSlice(rule.OrVals[i])
					if len(value) == 2 {
						expr = append(expr, builder.Between(field, value[0], value[1]))
					}
				}
			}
			builder.Where(builder.Or(expr...))
		} else {
			if rule.ValFunc != nil {
				rule.Val = rule.ValFunc()
			}
			switch Op(strings.ToUpper(string(rule.Op))) {
			case Equal:
				builder.Where(builder.Equal(rule.Key, rule.Val))
			case NotEqual:
				builder.Where(builder.NotEqual(rule.Key, rule.Val))
			case GreaterThan:
				builder.Where(builder.GreaterThan(rule.Key, rule.Val))
			case LessThan:
				builder.Where(builder.LessThan(rule.Key, rule.Val))
			case GreaterEqualThan:
				builder.Where(builder.GreaterEqualThan(rule.Key, rule.Val))
			case LessEqualThan:
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
	}
}

func ApplyDelete(builder *sqlbuilder.DeleteBuilder, rules ...Rule) {
	for _, rule := range rules {
		if rule.SkipFunc != nil {
			rule.Skip = rule.SkipFunc()
		}
		if rule.Skip {
			continue
		}
		if rule.Or {
			if rule.OrValsFunc != nil {
				rule.OrVals = rule.OrValsFunc()
			}
			var expr []string
			for i, field := range rule.OrKeys {
				switch Op(strings.ToUpper(string(rule.OrOps[i]))) {
				case Equal:
					expr = append(expr, builder.Equal(field, rule.OrVals[i]))
				case NotEqual:
					expr = append(expr, builder.NotEqual(field, rule.OrVals[i]))
				case GreaterThan:
					expr = append(expr, builder.GreaterThan(field, rule.OrVals[i]))
				case LessThan:
					expr = append(expr, builder.LessThan(field, rule.OrVals[i]))
				case GreaterEqualThan:
					expr = append(expr, builder.GreaterEqualThan(field, rule.OrVals[i]))
				case LessEqualThan:
					expr = append(expr, builder.LessEqualThan(field, rule.OrVals[i]))
				case In:
					if len(convert.ReflectSlice(rule.OrVals[i])) > 0 {
						expr = append(expr, builder.In(field, convert.ReflectSlice(rule.OrVals[i])...))
					}
				case NotIn:
					if len(convert.ReflectSlice(rule.OrVals[i])) > 0 {
						expr = append(expr, builder.NotIn(field, convert.ReflectSlice(rule.OrVals[i])...))
					}
				case Like:
					expr = append(expr, builder.Like(field, rule.OrVals[i]))
				case NotLike:
					expr = append(expr, builder.NotLike(field, rule.OrVals[i]))
				case Between:
					value := convert.ReflectSlice(rule.OrVals[i])
					if len(value) == 2 {
						expr = append(expr, builder.Between(field, value[0], value[1]))
					}
				}
			}
			builder.Where(builder.Or(expr...))
		} else {
			if rule.ValFunc != nil {
				rule.Val = rule.ValFunc()
			}
			switch Op(strings.ToUpper(string(rule.Op))) {
			case Equal:
				builder.Where(builder.Equal(rule.Key, rule.Val))
			case NotEqual:
				builder.Where(builder.NotEqual(rule.Key, rule.Val))
			case GreaterThan:
				builder.Where(builder.GreaterThan(rule.Key, rule.Val))
			case LessThan:
				builder.Where(builder.LessThan(rule.Key, rule.Val))
			case GreaterEqualThan:
				builder.Where(builder.GreaterEqualThan(rule.Key, rule.Val))
			case LessEqualThan:
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
	}
}
