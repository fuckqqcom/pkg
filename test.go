package main

// import (
//
//	"fmt"
//	"github.com/fuckqqcom/pkg/convert"
//	"github.com/huandu/go-sqlbuilder"
//	"github.com/spf13/cast"
//
// )
type Op string

func (o Op) String() string {
	return string(o)
}

const (
	E       Op = "="
	NE      Op = "!="
	GT      Op = ">"
	LT      Op = "<"
	GTE     Op = ">="
	LTE     Op = "<="
	In      Op = "IN"
	NotIn   Op = "NOT IN"
	Like    Op = "LIKE"
	NotLike Op = "NOT LIKE"
	Limit   Op = "LIMIT"
	Offset  Op = "OFFSET"
	Between Op = "BETWEEN"
	OrderBy Op = "ORDER BY"
)

type Rule struct {
	Key string

	Skip     bool
	SkipFunc func() bool

	// Or condition
	Or         bool
	OrOps      []Op
	OrKeys     []string
	OrVals     []any
	OrValsFunc func() []any

	// And condition
	Op      Op
	Val     any
	ValFunc func() any
}

func New(rules ...Rule) []Rule {
	return rules
}

//
//// buildCondition 构建单一条件表达式
//func buildCondition(cond *sqlbuilder.Cond, key string, op Op, value any) string {
//	switch op {
//	case E:
//		return cond.Equal(key, value)
//	case NE:
//		return cond.NotEqual(key, value)
//	case GT:
//		return cond.GreaterThan(key, value)
//	case LT:
//		return cond.LessThan(key, value)
//	case GTE:
//		return cond.GreaterEqualThan(key, value)
//	case LTE:
//		return cond.LessEqualThan(key, value)
//	case In:
//		return cond.In(key, convert.ReflectSlice(value)...) // 使用 convert.ReflectSlice 替代 castx.ToSlice
//	case NotIn:
//		return cond.NotIn(key, convert.ReflectSlice(value)...) // 使用 convert.ReflectSlice 替代 castx.ToSlice
//	case Like:
//		return cond.Like(key, value)
//	case NotLike:
//		return cond.NotLike(key, value)
//	case Between:
//		valueSlice := convert.ReflectSlice(value) // 使用 convert.ReflectSlice 替代 castx.ToSlice
//		if len(valueSlice) == 2 {
//			return cond.Between(key, valueSlice[0], valueSlice[1])
//		}
//	}
//	return ""
//}
//
//// processConditions 处理规则条件，将 Or 和 And 条件拼接
//func processConditions(conditions []Rule, cond *sqlbuilder.Cond) []string {
//	var exprs []string
//	for _, c := range conditions {
//		if c.SkipFunc != nil && c.SkipFunc() {
//			continue
//		}
//
//		// 处理 Or 条件
//		if c.Or {
//			// 如果存在 OrValsFunc，更新 OrVals
//			if c.OrValsFunc != nil {
//				c.OrVals = c.OrValsFunc()
//			}
//
//			// 对于 Or 条件，构建每个字段与值的条件，并用 OR 拼接
//			var orExprs []string
//			for i, field := range c.OrKeys {
//				// 使用对应的操作符处理每个字段
//				expr := buildCondition(cond, field, c.OrOps[i], c.OrVals[i])
//				if expr != "" {
//					orExprs = append(orExprs, expr)
//				}
//			}
//
//			// 只在 OrExprs 不为空时加入
//			if len(orExprs) > 0 {
//				// 使用 cond.Or(exprs...) 来自动处理 OR 条件的拼接
//				// sqlbuilder 会自动添加括号
//				cond.Or(orExprs...) // 不再使用额外的括号拼接
//			}
//		} else {
//			// 处理 And 条件
//			if c.ValFunc != nil {
//				c.Val = c.ValFunc()
//			}
//			expr := buildCondition(cond, c.Key, c.Op, c.Val)
//			if expr != "" {
//				exprs = append(exprs, expr)
//			}
//		}
//	}
//	return exprs
//}
//
//// buildWhereClause 构建 SQL where 子句
//func buildWhereClause(conditions ...Rule) *sqlbuilder.WhereClause {
//	clause := sqlbuilder.NewWhereClause()
//	cond := sqlbuilder.NewCond()
//
//	exprs := processConditions(conditions, cond)
//	if len(exprs) > 0 {
//		// 使用 AND 拼接各个条件
//		clause.AddWhereExpr(cond.Args, cond.And(exprs...))
//	}
//
//	return clause
//}
//
//// ApplySelect 应用 SELECT 查询条件
//func ApplySelect(sb *sqlbuilder.SelectBuilder, conditions ...Rule) {
//	clause := buildWhereClause(conditions...)
//	for _, c := range conditions {
//		if c.SkipFunc != nil && c.SkipFunc() {
//			continue
//		}
//		if c.ValFunc != nil {
//			c.Val = c.ValFunc()
//		}
//		switch c.Op {
//		case Limit:
//			sb.Limit(cast.ToInt(c.Val))
//		case Offset:
//			sb.Offset(cast.ToInt(c.Val))
//		case OrderBy:
//			if len(convert.ReflectSlice(c.Val)) > 0 {
//				sb.OrderBy(cast.ToStringSlice(convert.ReflectSlice(c.Val))...)
//			}
//		}
//	}
//	if clause != nil {
//		sb = sb.AddWhereClause(clause)
//	}
//}
//
//// ApplyUpdate 应用 UPDATE 查询条件
//func ApplyUpdate(sb *sqlbuilder.UpdateBuilder, conditions ...Rule) {
//	clause := buildWhereClause(conditions...)
//	for _, c := range conditions {
//		if c.SkipFunc != nil && c.SkipFunc() {
//			continue
//		}
//		if c.ValFunc != nil {
//			c.Val = c.ValFunc()
//		}
//		switch c.Op {
//		case Limit:
//			sb.Limit(cast.ToInt(c.Val))
//		case OrderBy:
//			if len(convert.ReflectSlice(c.Val)) > 0 {
//				sb.OrderBy(cast.ToStringSlice(convert.ReflectSlice(c.Val))...)
//			}
//		}
//	}
//	if clause != nil {
//		sb = sb.AddWhereClause(clause)
//	}
//}
//
//func main() {
//	sqlbuilder.DefaultFlavor = sqlbuilder.MySQL
//
//	//var values []any
//	//values = append(values, []int{24, 48}, []int{170, 175})
//	//
//	//cds := New(Rule{
//	//	Key: "name",
//	//	Op:  E,
//	//	Val: "jaronnie",
//	//}, Rule{
//	//	Or:     true,
//	//	OrKeys: []string{"age", "height"},
//	//	OrOps:  []Op{Between, Between},
//	//	OrVals: values,
//	//})
//	//
//	//sb := sqlbuilder.NewSelectBuilder().Select("name", "age", "height").From("user")
//	//ApplySelect(sb, cds...)
//	//
//	//sql, args := sb.Build()
//	//fmt.Println(sql)
//	//fmt.Println(args)
//
//	var values []any
//	values = append(values, []int{24, 48}, []int{170, 175})
//	cds := New(Rule{
//		SkipFunc: func() bool {
//			return true
//		},
//		Key: "name",
//		Op:  E,
//		Val: "jaronnie",
//		ValFunc: func() any {
//			return "jaronnie2"
//		},
//	}, Rule{
//		Or:     true,
//		OrKeys: []string{"age", "height"},
//		OrOps:  []Op{Between, Between},
//		OrVals: values,
//		OrValsFunc: func() []any {
//			return []any{[]int{24, 49}, []int{170, 176}}
//		},
//	})
//	clause := buildWhereClause(cds...)
//
//	fmt.Println(clause)
//	statement, args := clause.Build()
//	fmt.Println(statement)
//	fmt.Println(args)
//}
