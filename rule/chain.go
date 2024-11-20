package rule

//https://github.com/jzero-io/jzero-contrib/blob/main/condition/chain.go 源码出处
import (
	"github.com/fuckqqcom/pkg/opts"
)

type Chain struct {
	rules []Rule
}

type ChainOptions struct {
	Skip     bool
	SkipFunc func() bool

	ValueFunc func() any

	OrValuesFunc func() []any
}

func (opts ChainOptions) Options() ChainOptions {
	return ChainOptions{}
}

func WithSkip(skip bool) opt.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.Skip = skip
	}
}

func WithSkipFunc(skipFunc func() bool) opt.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.SkipFunc = skipFunc
	}
}

func WithValue(value any) opt.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.ValueFunc = func() any {
			return value
		}
	}
}

func WithOrValues(orValues []any) opt.Opt[ChainOptions] {
	return func(c *ChainOptions) {
		c.OrValuesFunc = func() []any {
			return orValues
		}
	}
}

func NewChain() Chain {
	return Chain{}
}

func NewChainRules(rules ...Rule) Chain {
	return Chain{rules: rules}
}

func (c Chain) add(field string, op Op, value any, opts ...opt.Opt[ChainOptions]) Chain {
	o := opt.Bind(opts...)
	c.rules = append(c.rules, Rule{
		Key:        field,
		Op:         op,
		Val:        value,
		Skip:       o.Skip,
		SkipFunc:   o.SkipFunc,
		OrValsFunc: o.OrValuesFunc,
	})
	return c
}

func (c Chain) E(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, E, value, opts...)
}

func (c Chain) NE(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, NE, value, opts...)
}

func (c Chain) GT(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, GT, value, opts...)
}

func (c Chain) LT(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, LT, value, opts...)
}

func (c Chain) GTE(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, GTE, value, opts...)
}

func (c Chain) LTE(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, LTE, value, opts...)
}

func (c Chain) Like(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, Like, value, opts...)
}

func (c Chain) NotLike(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, NotLike, value, opts...)
}

func (c Chain) In(field string, values any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, In, values, opts...)
}

func (c Chain) NotIn(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, NotIn, value, opts...)
}

func (c Chain) Between(field string, value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add(field, Between, value, opts...)
}

func (c Chain) Or(fields []string, values []any, opts ...opt.Opt[ChainOptions]) Chain {
	o := opt.Bind(opts...)
	c.rules = append(c.rules, Rule{
		Or:         true,
		OrKeys:     fields,
		OrVals:     values,
		Skip:       o.Skip,
		SkipFunc:   o.SkipFunc,
		OrValsFunc: o.OrValuesFunc,
	})
	return c
}

func (c Chain) OrderBy(value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add("", OrderBy, value, opts...)
}

func (c Chain) Limit(value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add("", Limit, value, opts...)
}

func (c Chain) Offset(value any, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add("", Offset, value, opts...)
}

func (c Chain) Page(page, pageSize int, opts ...opt.Opt[ChainOptions]) Chain {
	return c.add("", Offset, (page-1)*pageSize, opts...).add("", Limit, pageSize, opts...)
}

func (c Chain) Build() []Rule {
	return c.rules
}
