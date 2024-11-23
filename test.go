package main

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

type Field struct {
	key      string      // 当前字段的名称
	val      any         // 当前字段的值
	skipFunc func() bool // 判断是否跳过该字段的操作
	valFunc  func() any  // 动态计算的值函数
	keys     []string    // 记录处理过的 key

	builder *sqlbuilder.UpdateBuilder
}
type Option func(*Field)

// NewField 创建一个新的 Field
func NewField(builder *sqlbuilder.UpdateBuilder) *Field {
	return &Field{
		builder: builder,
	}
}

// Set 用于设置字段的值
func (f *Field) Set(assignment []string, opts ...Option) *Field {
	f.builder.Set(assignment...)
	return f
}

// SetMore
func (f *Field) SetMore(assignment []string, opts ...Option) *Field {
	f.builder.SetMore(assignment...)
	return f
}

type Op string

func (o Op) String() string {
	return string(o)
}

const (
	Incr   Op = "Incr"
	Decr   Op = "Decr"
	Assign Op = "Assign"
	Add    Op = "Add"
	Sub    Op = "Sub"
	Mul    Op = "Mul"
	Div    Op = "Div"
)

func buildExpr(builder *sqlbuilder.UpdateBuilder, key string, operator Op, value any) string {
	switch operator {
	case Incr:
		return builder.Incr(key)
	case Decr:
		return builder.Decr(key)
	case Assign:
		return builder.Assign(key, value)
	case Add:
		return builder.Add(key, value)
	case Sub:
		return builder.Sub(key, value)
	case Mul:
		return builder.Mul(key, value)
	case Div:
		return builder.Div(key, value)
	default:
		return ""
	}
}

func main() {
	// Build a SQL to select a user from database.
	//sb := sqlbuilder.NewSelectBuilder().Select("name", "level").From("users")
	//sb.Where(
	//	sb.Equal("id", 1234),
	//)
	//fmt.Println(sb)

	ub := sqlbuilder.Update("users")
	ub.Set(
		ub.Div("level", "1"),
	)

	// Set the WHERE clause of UPDATE to the WHERE clause of SELECT.
	//ub.WhereClause = sb.WhereClause
	fmt.Println(ub)

	// Output:
	// SELECT name, level FROM users WHERE id = ?
	// UPDATE users SET level = level + ? WHERE id = ?
}
