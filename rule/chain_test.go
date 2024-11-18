package rule

import (
	"fmt"
	"testing"
)

func TestChain_ToRules(t *testing.T) {
	c := NewChain().
		E("age", 30, WithSkip(true), WithValue(30)).
		GT("salary", 5000)

	// 打印所有规则
	for _, rule := range c.Bind() {
		fmt.Printf("Key: %s, Operator: %s, Value: %v\n", rule.Key, rule.Op, rule.Val)
	}
}
