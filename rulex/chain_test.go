package rulex

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"testing"
)

func TestChain_ToRules(t *testing.T) {
	//sb := sqlbuilder.NewSelectBuilder().Select("name", "age").From("user")
	//
	//chain := NewChain().
	//	E("field1", "value1", WithSkip(true)).
	//	E("field2", "value2", WithValFunc(func() any {
	//		return "100"
	//	})).
	//	OrderBy("create_time desc").
	//	OrderBy("sort desc")
	//builder := Select(*sb, chain.Rule()...)
	//sql, args := builder.Build()
	//fmt.Println(sql)
	//fmt.Println(args)
	builder := sqlbuilder.NewUpdateBuilder()
	chain := NewChain().E("id", "1").SetIncr("key1").SetDecr("key2")
	sb := Update(builder, chain.Rule()...)
	sql, args := sb.Build()
	fmt.Println(sql, args)
	//chain = chain.NE("1", "value")
	//builder = Select(*sb, chain.Build()...)
	//
	//sql, args = builder.Build()
	//fmt.Println(sql)
	//fmt.Println(args)
}
