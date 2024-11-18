package field

import (
	"fmt"
	"reflect"
)

// Field 结构体定义
type Field struct {
	Key      string
	Val      any
	SkipFunc func() bool
	ValFunc  func() any
	next     *Field
}

func NewField() *Field {
	return &Field{}
}

// Option 类型用于对 Field 进行配置
type Option func(*Field)

// SetVal 设置 Field 的值，并接受可选的配置参数 opts，用于定制化行为
func (f *Field) SetVal(key string, val any, opts ...Option) *Field {
	f.Key = key
	f.Val = val

	for _, opt := range opts {
		opt(f)
	}

	// 返回当前 Field 支持链式调用
	if f.next != nil {
		return f.next.SetVal(f.next.Key, f.next.Val, opts...)
	}

	return f
}

func Skip(skipFunc func() bool) Option {
	return func(f *Field) {
		f.SkipFunc = skipFunc
	}
}

func ValFunc(valFunc func() any) Option {
	return func(f *Field) {
		f.ValFunc = valFunc
	}
}

func (f *Field) ApplyTo(obj any) (errs []error) {
	vals := reflect.ValueOf(obj).Elem()

	for field := f; field != nil; field = field.next {
		if field.SkipFunc != nil && field.SkipFunc() {
			continue
		}

		if field.ValFunc != nil {
			field.Val = field.ValFunc()
		}

		val := vals.FieldByName(field.Key)

		if !val.IsValid() {
			errs = append(errs, fmt.Errorf("field %s not found", field.Key))
			continue
		}

		if val.CanSet() {
			val.Set(reflect.ValueOf(field.Val))
		}
	}

	return
}

// Next 方法创建并返回下一个 Field，用于链式调用
func (f *Field) Next() *Field {
	field := &Field{}
	f.next = field
	return field
}
