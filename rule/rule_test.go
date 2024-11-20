package rule

import (
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"testing"
)

func TestSelect(t *testing.T) {
	sb := sqlbuilder.NewSelectBuilder().Select("name", "age", "height").From("user")
	var values []any
	values = append(values, []int{24, 48}, []int{170, 175})
	cds := NewRule(Rule{
		SkipFunc: func() bool {
			return true
		},
		Key: "name",
		Op:  E,
		Val: "jaronnie",
		ValFunc: func() any {
			return "jaronnie2"
		},
	}, Rule{
		Or:     true,
		OrKeys: []string{"age", "height"},
		OrOps:  []Op{Between, Between},
		OrVals: values,
		OrValsFunc: func() []any {
			return []any{[]int{24, 49}, []int{170, 176}}
		},
	},
		Rule{
			Key: "name",
			Op:  FindINSet,
			Val: "jaronnie",
		},
		Rule{
			Op:  Limit,
			Val: 10,
		},
	)

	builder := Select(*sb, cds...)

	statement, args := builder.Build()
	fmt.Println(statement)
	fmt.Println(args)
}
