package rulex

//https://github.com/jzero-io/jzero-contrib/blob/main/condition/chain.go 源码出处
import (
	"github.com/ettle/strcase"
	"github.com/fuckqqcom/pkg/optx"
)

type Chain struct {
	keyFunc func(string) string // 新增 KeyFunc 字段
	rules   []Rule
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

func NewChain(fs ...func(string) string) Chain {
	var keyFunc func(string) string
	if len(fs) > 0 {
		keyFunc = fs[0]
	} else {
		keyFunc = strcase.ToSnake
	}
	return Chain{keyFunc: keyFunc}
}

func NewChainRules(rules ...Rule) Chain {
	return Chain{rules: rules}
}

func (c Chain) add(field string, op Op, val any, opts ...optx.Opt[ChainOptions]) Chain {
	o := optx.Bind(opts...)
	c.rules = append(c.rules, Rule{
		Key:        c.keyFunc(field),
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

func (c Chain) E(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, E, val, opts...)
}

func (c Chain) NE(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, NE, val, opts...)
}

func (c Chain) GT(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, GT, val, opts...)
}

func (c Chain) LT(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, LT, val, opts...)
}

func (c Chain) GTE(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, GTE, val, opts...)
}

func (c Chain) LTE(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, LTE, val, opts...)
}

func (c Chain) Like(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Like, val, opts...)
}

func (c Chain) NotLike(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, NotLike, val, opts...)
}

func (c Chain) In(field string, vals any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, In, vals, opts...)
}

func (c Chain) NotIn(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, NotIn, val, opts...)
}

func (c Chain) Between(field string, val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add(field, Between, val, opts...)
}

// func (c Chain) Or(fields []string, values []any, opts ...optx.Opt[ChainOptions]) Chain {
func (c Chain) Or(fields []string, ops []Op, vals []any, opts ...optx.Opt[ChainOptions]) Chain {
	o := optx.Bind(opts...)

	if o.OrValsFunc != nil {
		o.orVals = o.OrValsFunc()
	}

	var keys []string
	for _, field := range fields {
		keys = append(keys, c.keyFunc(field))
	}
	c.rules = append(c.rules, Rule{
		Or:         true,
		OrKeys:     keys,
		OrOps:      ops,
		skip:       o.skip,
		SkipFunc:   o.SkipFunc,
		ValFunc:    o.ValFunc,
		orVals:     vals,
		OrValsFunc: o.OrValsFunc,
	})
	return c
}

func (c Chain) OrderBy(val any, opts ...optx.Opt[ChainOptions]) Chain {
	o := optx.Bind(opts...)
	if o.ValFunc != nil {
		val = o.ValFunc()
	}
	switch vals := val.(type) {
	case map[string]any:
		for k, v := range vals {
			c.rules = append(c.rules, Rule{
				Key:      c.keyFunc(k),
				Op:       OrderBy,
				val:      v,
				skip:     o.skip,
				SkipFunc: o.SkipFunc,
			})
		}
	case map[string]string:
		for k, v := range vals {
			c.rules = append(c.rules, Rule{
				Key:      c.keyFunc(k),
				Op:       OrderBy,
				val:      v,
				skip:     o.skip,
				SkipFunc: o.SkipFunc,
			})
		}
	default:
		c.rules = append(c.rules, Rule{
			Key:      c.keyFunc(""),
			Op:       OrderBy,
			val:      val,
			skip:     o.skip,
			SkipFunc: o.SkipFunc,
		})
		c.add("", OrderBy, val, opts...)
	}
	return c
}

func (c Chain) Limit(val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add("", Limit, val, opts...)
}

func (c Chain) Offset(val any, opts ...optx.Opt[ChainOptions]) Chain {
	return c.add("", Offset, val, opts...)
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
