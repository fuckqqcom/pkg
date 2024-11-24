package rulex

import (
	"fmt"
	"github.com/ettle/strcase"
	"github.com/huandu/go-sqlbuilder"
	"testing"
)

func TestChain_ToRules(t *testing.T) {
	sb := sqlbuilder.NewSelectBuilder().Select("name", "age").From("user")

	chain := NewChain(strcase.ToSnake).
		E("createTime", "value1", WithSkip(false)).
		E("field2", "value2", WithValFunc(func() any {
			return "100"
		})).
		OrderBy("createTime", " desc").
		OrderBy("sort", " desc")
	s := Select(sb, chain.Rule()...)
	sql, args := s.Build()
	fmt.Println(sql, args)
	chain = chain.SetIncr("key1", WithSkipFunc(func() bool {
		return false
	}))
	builder := sqlbuilder.NewUpdateBuilder()
	b := Update(builder, chain.Rule()...)
	sql, args = b.Build()
	fmt.Println(sql, args)

	//chain = chain.NE("1", "value")
	//builder = Select(*sb, chain.Build()...)
	//
	//sql, args = builder.Build()
	//fmt.Println(sql)
	//fmt.Println(args)
}
