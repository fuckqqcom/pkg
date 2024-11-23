package rulex

//https://github.com/jzero-io/jzero-contrib/blob/main/condition/chain.go 源码出处
import (
	"github.com/fuckqqcom/pkg/optx"
)

type Chain struct {
	rules []Rule
}

func (c Chain) Rule() []Rule {
	return c.rules
}

type ChainOptions struct {
	skip     bool
	SkipFunc func() bool

	ValFunc func() any
	val     any

	orVals     []any
	OrValsFunc func() []any
}

func (opts ChainOptions) Options() ChainOptions {
	return ChainOptions{}
}

func WithSkip(skip bool) optx.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.skip = skip
	}
}

func WithSkipFunc(skipFunc func() bool) optx.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.SkipFunc = skipFunc
	}
}

func WithValFunc(valFunc func() any) optx.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.ValFunc = valFunc
	}
}

func WithOrVals(orVals []any) optx.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.orVals = orVals
	}
}

func WithOrValsFunc(orValsFunc func() []any) optx.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.OrValsFunc = orValsFunc
	}
}

func NewChain() Chain {
	return Chain{}
}

func NewChainRules(rules ...Rule) Chain {
	return Chain{rules: rules}
}

func (c Chain) add(field string, op Op, val any, opts ...optx.Opt[ChainOptions]) Chain {
	o := optx.Bind(opts...)
	c.rules = append(c.rules, Rule{
		Key:        field,
		Op:         op,
		val:        val,
		skip:       o.skip,
		SkipFunc:   o.SkipFunc,
		ValFunc:    o.ValFunc,
		orVals:     o.orVals,
		OrValsFunc: o.OrValsFunc,
	})
	return c
}

func (c Chain) E(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, E, value, opts...)
}

func (c Chain) NE(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, NE, value, opts...)
}

func (c Chain) GT(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, GT, value, opts...)
}

func (c Chain) LT(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, LT, value, opts...)
}

func (c Chain) GTE(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, GTE, value, opts...)
}

func (c Chain) LTE(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, LTE, value, opts...)
}

func (c Chain) Like(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Like, value, opts...)
}

func (c Chain) NotLike(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, NotLike, value, opts...)
}

func (c Chain) In(field string, values any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, In, values, opts...)
}

func (c Chain) NotIn(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, NotIn, value, opts...)
}

func (c Chain) Between(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Between, value, opts...)
}

func (c Chain) Or(fields []string, values []any, opts ...optx.Opt[ChainOptions]) Chain {
	o := optx.Bind(opts...)
	c.rules = append(c.rules, Rule{
		Or:         true,
		val:        o.val,
		OrKeys:     fields,
		skip:       o.skip,
		SkipFunc:   o.SkipFunc,
		ValFunc:    o.ValFunc,
		orVals:     o.orVals,
		OrValsFunc: o.OrValsFunc,
	})
	return c
}

func (c Chain) OrderBy(value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add("", OrderBy, value, opts...)
}

func (c Chain) Limit(value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add("", Limit, value, opts...)
}

func (c Chain) Offset(value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add("", Offset, value, opts...)
}

func (c Chain) Page(page, pageSize int, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add("", Offset, (page-1)*pageSize, opts...).add("", Limit, pageSize, opts...)
}

/*
	set操作
*/

func (c Chain) SetIncr(field string, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Incr, "", opts...)
}

func (c Chain) SetDecr(field string, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Decr, "", opts...)
}

func (c Chain) SetAssign(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Assign, value, opts...)
}

func (c Chain) SetAdd(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Add, value, opts...)
}
func (c Chain) SetSub(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Sub, value, opts...)
}

func (c Chain) SetMul(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Mul, value, opts...)
}

func (c Chain) SetDiv(field string, value any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Div, value, opts...)
}
