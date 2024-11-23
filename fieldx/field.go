package fieldx

import (
	"fmt"
	"github.com/ettle/strcase"
	"reflect"
)

type Field struct {
	key        string
	val        any
	skipFunc   func() bool
	valFunc    func() any
	nodes      []*Field        // 改为切片存储节点，以支持链式调用
	keys       []string        //存储所有处理的key
	skipKeys   []string        //存储跳过的key
	ignoreKeys map[string]bool //
	errs       []error
}

// NewField 创建新的 Field 实例
func NewField() *Field {
	return &Field{ignoreKeys: make(map[string]bool)}
}

// SetVal 设置 Field 的值，并接受可选的配置参数 optx，用于定制化行为
func (f *Field) SetVal(key string, val any, opts ...Option) *Field {
	f.key = strcase.ToPascal(key) // 将 key 转换为 PascalCase 格式
	f.val = val
	f.keys = append(f.keys, f.key)
	// 应用所有传入的配置选项
	for _, opt := range opts {
		opt(f)
	}
	return f.next()
}

// next 方法用于返回下一个 Field 节点
func (f *Field) next() *Field {
	node := &Field{key: f.key, val: f.val, skipFunc: f.skipFunc, valFunc: f.valFunc}
	f.nodes = append(f.nodes, node) // 添加到数组中
	return f
}

// WithSkipFunc WithValFunc 设置 SkipFunc，决定是否跳过此字段
func WithSkipFunc(skipFunc func() bool) Option {
	return func(f *Field) {
		f.skipFunc = skipFunc
	}
}

// WithValFunc 设置 ValFunc，决定如何动态计算该字段的值
func WithValFunc(valFunc func() any) Option {
	return func(f *Field) {
		f.valFunc = valFunc
	}
}

func (f *Field) SetIgnoreKey(keys []string) *Field {
	for _, key := range keys {
		f.ignoreKeys[strcase.ToPascal(key)] = true
	}
	return f
}

// Bind 将 Field 应用到目标对象上，返回错误列表
func (f *Field) Bind(obj any) *Field {
	vals := reflect.ValueOf(obj).Elem()

	// 遍历当前 Field 链表，逐个处理
	for _, field := range f.nodes {
		// 如果 SkipFunc 返回 true，则跳过此字段
		if field.skipFunc != nil && field.skipFunc() {
			f.skipKeys = append(f.skipKeys, field.key)
			continue
		}
		if f.ignoreKeys[field.key] {
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
			f.errs = append(f.errs, fmt.Errorf("fieldx %s not found", field.key))
			continue
		}

		// 如果字段可设置，赋值
		if val.CanSet() {
			val.Set(reflect.ValueOf(field.val))
		}
	}

	return f
}

// Check 检查要忽略的key等信息 返回true表示没有要执行的更新操作的字段
func (f *Field) Check() bool {

	excludeField := make(map[string]bool)
	// 将 skipKeys 和 ignoreKey 中的元素添加到排除列表
	for _, key := range f.skipKeys {
		excludeField[key] = true
	}
	for key := range f.ignoreKeys {
		excludeField[key] = true
	}
	// 创建一个新的切片存储计算结果
	var result []string
	for _, key := range f.keys {
		if !excludeField[key] {
			result = append(result, key)
			break
		}
	}
	return len(result) == 0
}

// Option 类型用于对 Field 进行配置
type Option func(*Field)
