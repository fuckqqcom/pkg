package opt

// https://github.com/eddieowens/opts/blob/main/opts.go 源码出处

// Opt 是一个类型别名，用于定义可以修改 T 类型选项的函数。
type Opt[T any] func(*T)

// Option 是一个工厂接口，用于创建 T 类型的默认选项。
type Option[T any] interface {
	// Option 返回 T 类型的默认值
	Option() T
}

// Bind 使用默认工厂来构造一个默认的选项实例，并运行传入的 Opt 修改器。
// 它可以接受多个 Opt 函数参数，并将它们应用到默认选项实例上。
func Bind[T any](opts ...Opt[T]) T {
	// 创建一个 T 类型的默认实例，假设 Option 接口的 Option 方法已返回一个默认值
	var a T

	// 如果 T 实现了 Option 接口，则使用其 Option 方法来初始化
	if opt, ok := any(&a).(Option[T]); ok {
		a = opt.Option()
	}

	// 将所有传入的 Opt 函数应用到默认选项实例上
	Apply(&a, opts...)
	return a
}

// Apply 将多个 Opt 函数应用到给定的选项实例。
func Apply[T any](o *T, opts ...Opt[T]) {
	for _, v := range opts {
		v(o)
	}
}
