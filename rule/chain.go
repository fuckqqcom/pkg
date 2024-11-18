package rule

//https://github.com/jzero-io/jzero-contrib/blob/main/condition/chain.go 源码出处
import (
	"github.com/fuckqqcom/pkg/opts"
)

// 操作符定义

// ChainOperatorOpts 是一个选项结构体，表示链式操作符的配置
type ChainOperatorOpts struct {
	Skip         bool
	SkipFunc     func() bool
	ValueFunc    func() any
	OrValuesFunc func() []any
}

func (opts ChainOperatorOpts) DefaultOptions() ChainOperatorOpts {
	return ChainOperatorOpts{}
}

func WithSkip(skip bool) opt.Opt[ChainOperatorOpts] {
	return func(c *ChainOperatorOpts) {
		c.Skip = skip
	}
}

func WithSkipFunc(skipFunc func() bool) opt.Opt[ChainOperatorOpts] {
	return func(c *ChainOperatorOpts) {
		c.SkipFunc = skipFunc
	}
}

func WithValue(value any) opt.Opt[ChainOperatorOpts] {
	return func(c *ChainOperatorOpts) {
		c.ValueFunc = func() any {
			return value
		}
	}
}

func WithOrValues(orValues []any) opt.Opt[ChainOperatorOpts] {
	return func(c *ChainOperatorOpts) {
		c.OrValuesFunc = func() []any {
			return orValues
		}
	}
}

// Chain 表示一个条件链，可以链接多个查询条件
type Chain struct {
	rules []Rule
}

func NewChain() Chain {
	return Chain{}
}

// AddRule 是通用的条件添加方法
func (c Chain) AddRule(key string, op Op, val any, opts ...opt.Opt[ChainOperatorOpts]) Chain {
	o := opt.Bind(opts...)
	rule := Rule{
		Key:      key,
		Op:       op,
		Val:      val,
		Skip:     o.Skip,
		SkipFunc: o.SkipFunc,
		ValFunc:  o.ValueFunc,
	}
	c.rules = append(c.rules, rule)
	return c
}

// 以下是更新后的条件方法，操作符名称改为简写：

// E 添加等于条件
func (c Chain) E(key string, val any, op ...opt.Opt[ChainOperatorOpts]) Chain {
	return c.AddRule(key, E, val, op...)
}

// NE 添加不等于条件
func (c Chain) NE(key string, val any, op ...opt.Opt[ChainOperatorOpts]) Chain {
	return c.AddRule(key, NE, val, op...)
}

// GT 添加大于条件
func (c Chain) GT(key string, val any, op ...opt.Opt[ChainOperatorOpts]) Chain {
	return c.AddRule(key, GT, val, op...)
}

// LT 添加小于条件
func (c Chain) LT(key string, val any, op ...opt.Opt[ChainOperatorOpts]) Chain {
	return c.AddRule(key, LT, val, op...)
}

// GTE 添加大于等于条件
func (c Chain) GTE(key string, val any, op ...opt.Opt[ChainOperatorOpts]) Chain {
	return c.AddRule(key, GTE, val, op...)
}

// LTE 添加小于等于条件
func (c Chain) LTE(key string, val any, op ...opt.Opt[ChainOperatorOpts]) Chain {
	return c.AddRule(key, LTE, val, op...)
}

// ToRules 返回所有的规则
func (c Chain) Bind() []Rule {
	return c.rules
}
