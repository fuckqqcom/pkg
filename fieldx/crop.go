package fieldx

import (
	"fmt"
	"github.com/fuckqqcom/pkg/convertx"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
	"strings"
)

func FilterProtoField(obj proto.Message, includeFields []string, excludeFields []string, merge bool) map[string]any {

	/*
		先判断 includeFields和excludeFields共有的字段
	*/
	fieldMap := map[string]bool{}
	for _, v := range includeFields {
		if !convertx.ContainIgnoreCase(v, excludeFields) {
			fieldMap[v] = false
		}
	}
	/*
			 includeFields [1,2,3,6]
			 excludeFields [3,4,5]
		     all: 1,2,3,4,5,6,7

			merge:
				false notRemoveFields: 1,2,6
				true  notRemoveFields: 1,2,6,7

	*/
	protoFieldMap := map[string]bool{}
	fieldsAttrs := obj.ProtoReflect().Descriptor().Fields()
	for i := 0; i < fieldsAttrs.Len(); i++ {
		name := fieldsAttrs.Get(i).TextName()
		if merge && !convertx.ContainIgnoreCase(name, excludeFields) {
			fieldMap[name] = true
			continue
		}
		protoFieldMap[name] = true

	}
	//去重和删除重复元素
	for key, val := range fieldMap {
		if val {
			continue
		}
		if ok := protoFieldMap[key]; !ok {
			delete(fieldMap, key)
		}
	}

	data := make(map[string]any)
	for field, _ := range fieldMap {
		data[field] = convertx.AnyToStr(obj.ProtoReflect().Get(fieldsAttrs.ByName(protoreflect.Name(field))).Interface())
	}
	return data
}

func CropProtoMessages(objs []proto.Message, includeFields []string, excludeFields []string) []map[string]any {
	var data []map[string]any

	// 使用 map 来存储排除的字段，避免重复遍历 excludeFields
	excludeMap := make(map[string]struct{}, len(excludeFields))
	for _, field := range excludeFields {
		excludeMap[field] = struct{}{}
	}

	// 转换 includeFields 为 set 形式，避免在循环中频繁使用 append
	includeMap := make(map[string]struct{}, len(includeFields))
	for _, field := range includeFields {
		includeMap[field] = struct{}{}
	}

	for _, obj := range objs {
		fieldsAttrs := obj.ProtoReflect().Descriptor().Fields()

		// 过滤出需要的字段，并构建信息
		info := make(map[string]interface{}, len(includeFields))

		for i := 0; i < fieldsAttrs.Len(); i++ {
			name := fieldsAttrs.Get(i).TextName()

			// 如果字段被排除，跳过
			if _, excluded := excludeMap[name]; excluded {
				continue
			}

			// 如果字段在 includeFields 中，添加到结果 map
			if _, included := includeMap[name]; included || len(includeFields) == 0 {
				info[name] = obj.ProtoReflect().Get(fieldsAttrs.ByName(protoreflect.Name(name))).Interface()
			}
		}

		// 将信息添加到最终结果中
		data = append(data, info)
	}

	return data
}

func CropObjFields(obj any, fields []string) {
	object := reflect.ValueOf(obj)
	elems := object.Elem()
	typeOfType := elems.Type()

	// 将需要忽略的字段名称转为map，便于快速查找
	ignoreFields := make(map[string]struct{}, len(fields))
	for _, fieldName := range fields {
		ignoreFields[strings.ToLower(fieldName)] = struct{}{}
	}

	// 遍历结构体字段并重置
	for i := 0; i < elems.NumField(); i++ {
		field := elems.Field(i)
		fieldType := typeOfType.Field(i)
		fieldName := fieldType.Name

		if _, found := ignoreFields[strings.ToLower(fieldName)]; !found {
			resetFieldValue(field)
		}
	}
}
func resetFieldValue(field reflect.Value) {
	// 如果字段不可设置，直接返回
	if !field.CanSet() {
		return
	}

	// 使用一个映射来简化类型的设置
	defaultValues := map[reflect.Kind]func(reflect.Value){
		reflect.String:  func(v reflect.Value) { v.SetString("") },
		reflect.Int:     func(v reflect.Value) { v.SetInt(0) },
		reflect.Int8:    func(v reflect.Value) { v.SetInt(0) },
		reflect.Int16:   func(v reflect.Value) { v.SetInt(0) },
		reflect.Int32:   func(v reflect.Value) { v.SetInt(0) },
		reflect.Int64:   func(v reflect.Value) { v.SetInt(0) },
		reflect.Uint:    func(v reflect.Value) { v.SetUint(0) },
		reflect.Uint8:   func(v reflect.Value) { v.SetUint(0) },
		reflect.Uint16:  func(v reflect.Value) { v.SetUint(0) },
		reflect.Uint32:  func(v reflect.Value) { v.SetUint(0) },
		reflect.Uint64:  func(v reflect.Value) { v.SetUint(0) },
		reflect.Float32: func(v reflect.Value) { v.SetFloat(0) },
		reflect.Float64: func(v reflect.Value) { v.SetFloat(0) },
		reflect.Bool:    func(v reflect.Value) { v.SetBool(false) },
	}

	// 获取字段的类型并设置其值
	resetFunc, ok := defaultValues[field.Kind()]
	if ok {
		resetFunc(field)
	} else {
		// 如果没有匹配到类型，抛出错误
		panic(fmt.Sprintf("Unhandled type: %v", field.Kind()))
	}
}
