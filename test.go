package main

import (
	"fmt"
	"github.com/ettle/strcase"
	"reflect"
)

// Field 结构体定义
type Field struct {
	key      string
	val      any
	skipFunc func() bool
	valFunc  func() any
	nodes    []*Field // 改为切片存储节点，以支持链式调用
}

// NewField 创建新的 Field 实例
func NewField() *Field {
	return &Field{}
}

// SetVal 设置 Field 的值，并接受可选的配置参数 optx，用于定制化行为
func (f *Field) SetVal(key string, val any, opts ...Option) *Field {
	f.key = strcase.ToPascal(key) // 将 key 转换为 PascalCase 格式
	f.val = val

	// 应用所有传入的配置选项
	for _, opt := range opts {
		opt(f)
	}

	// 如果有下一个字段，返回下一个 Field 实例
	fmt.Println("f.nodes--->", f.nodes)
	return f.next()
}

// next 方法用于返回下一个 Field 节点
func (f *Field) next() *Field {
	node := &Field{key: f.key, val: f.val, skipFunc: f.skipFunc, valFunc: f.valFunc}
	f.nodes = append(f.nodes, node) // 添加到链表中
	return f
}

// SkipFunc 设置 SkipFunc，决定是否跳过此字段
func SkipFunc(skipFunc func() bool) Option {
	return func(f *Field) {
		f.skipFunc = skipFunc
	}
}

// ValFunc 设置 ValFunc，决定如何动态计算该字段的值
func ValFunc(valFunc func() any) Option {
	return func(f *Field) {
		f.valFunc = valFunc
	}
}

// Bind 将 Field 应用到目标对象上，返回错误列表
func (f *Field) Bind(obj any) (errs []error) {
	vals := reflect.ValueOf(obj).Elem()

	// 遍历当前 Field 链表，逐个处理
	for _, field := range f.nodes {
		// 如果 SkipFunc 返回 true，则跳过此字段
		if field.skipFunc != nil && field.skipFunc() {
			continue
		}

		// 如果 ValFunc 存在，则通过它计算字段的值
		if field.valFunc != nil {
			field.val = field.valFunc()
		}

		// 通过反射获取目标对象中对应的字段
		val := vals.FieldByName(field.key)

		// 如果字段无效，返回错误
		if !val.IsValid() {
			errs = append(errs, fmt.Errorf("fieldx %s not found", field.key))
			continue
		}

		// 如果字段可设置，赋值
		if val.CanSet() {
			val.Set(reflect.ValueOf(field.val))
		}
	}

	return
}

// Option 类型用于对 Field 进行配置
type Option func(*Field)

func main() {
	type Person struct {
		Name        string
		Age         int
		Email       string
		PhoneNumber string
	}

	p := &Person{Age: 60}

	// 链式调用 SetVal 和 Bind，注意 SkipFunc 和 ValFunc 的配置
	NewField().SetVal("name", "John", SkipFunc(func() bool {
		return true // 不跳过 Name 字段
	})).
		SetVal("email", "john@example.com",
			SkipFunc(func() bool {
				return true // 不跳过 Email 字段
			}),
			ValFunc(func() any {
				return "new-email@example.com" // 通过 ValFunc 修改 Email 字段值
			}),
		).SetVal("phoneNumber", "1190",
		SkipFunc(func() bool {
			return true // 跳过 PhoneNumber 字段
		}),
		ValFunc(func() any {
			return "1190" // 该字段不会被设置，因为跳过了
		}),
	).Bind(p)

	// 输出结果: { John 60 new-email@example.com }
	fmt.Println(p)
}
