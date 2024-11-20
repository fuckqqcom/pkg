package rule

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"testing"
)

func TestChain_ToRules(t *testing.T) {
	sb := sqlbuilder.NewSelectBuilder().Select("name", "age").From("user")

	chain := NewChain().
		E("field1", "value1", WithSkip(true)).
		E("field2", "value2").
		OrderBy("create_time desc").
		OrderBy("sort desc")
	builder := Select(*sb, chain.Build()...)
	sql, args := builder.Build()
	fmt.Println(sql)
	fmt.Println(args)

	chain = chain.NE("1", "value")
	builder = Select(*sb, chain.Build()...)

	sql, args = builder.Build()
	fmt.Println(sql)
	fmt.Println(args)
}
