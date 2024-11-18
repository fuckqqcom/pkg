package field

import (
	"fmt"
	"reflect"
)

// Field 结构体定义
type Field struct {
	Key      string
	val      any
	skipFunc func() bool
	valFunc  func() any
	next     *Field
}

// NewField 创建新的 Field 实例
func NewField() *Field {
	return &Field{}
}

// SetVal 设置 Field 的值，并接受可选的配置参数 opts，用于定制化行为
func (f *Field) SetVal(key string, val any, opts ...Option) *Field {
	f.Key = key
	f.val = val

	// 应用所有传入的配置选项
	for _, opt := range opts {
		opt(f)
	}

	// 如果有 next 字段，继续设置下一个 Field
	if f.next != nil {
		return f.next.SetVal(f.next.Key, f.next.val, opts...)
	}

	// 返回当前 Field 支持链式调用
	return f
}

// SkipFunc 设置 SkipFunc，决定是否跳过此字段
func (f *Field) SkipFunc(skipFunc func() bool) Option {
	return func(f *Field) {
		f.skipFunc = skipFunc
	}
}

// ValFunc 设置 ValFunc，决定如何动态计算该字段的值
func (f *Field) ValFunc(valFunc func() any) Option {
	return func(f *Field) {
		f.valFunc = valFunc
	}
}

// Bind 将 Field 应用到目标对象上，返回错误列表
func (f *Field) Bind(obj any) (errs []error) {
	vals := reflect.ValueOf(obj).Elem()

	// 遍历当前 Field 链表，逐个处理
	for field := f; field != nil; field = field.next {
		// 如果 SkipFunc 返回 true，则跳过此字段
		if field.skipFunc != nil && field.skipFunc() {
			continue
		}

		// 如果 ValFunc 存在，则通过它计算字段的值
		if field.valFunc != nil {
			field.val = field.valFunc()
		}

		// 通过反射获取目标对象中对应的字段
		val := vals.FieldByName(field.Key)

		// 如果字段无效，返回错误
		if !val.IsValid() {
			errs = append(errs, fmt.Errorf("field %s not found", field.Key))
			continue
		}

		// 如果字段可设置，赋值
		if val.CanSet() {
			val.Set(reflect.ValueOf(field.val))
		}
	}

	return errs
}

// Next 方法创建并返回下一个 Field，用于链式调用
func (f *Field) Next() *Field {
	field := &Field{}
	f.next = field
	return field
}

// Option 类型用于对 Field 进行配置
type Option func(*Field)
