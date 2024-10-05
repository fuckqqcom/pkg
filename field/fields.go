package field

import (
	"github.com/fuckqqcom/pkg/convert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
)

func CropProtoFields(obj proto.Message, includeFields []string, excludeFields []string, merge bool) map[string]interface{} {

	/*
		先判断 includeFields和excludeFields共有的字段
	*/
	fieldMap := map[string]bool{}
	for _, v := range includeFields {
		if !convert.ContainIgnoreCase(v, excludeFields) {
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
		if merge && !convert.ContainIgnoreCase(name, excludeFields) {
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

	data := make(map[string]interface{})
	for field, _ := range fieldMap {
		data[field] = convert.AnyToStr(obj.ProtoReflect().Get(fieldsAttrs.ByName(protoreflect.Name(field))).Interface())
	}
	return data
}

func CropProtoMessages(objs []proto.Message, includeFields []string, excludeFields []string) []map[string]interface{} {
	//data := make([]map[string]interface{}, 1)
	var data []map[string]interface{}
	for _, obj := range objs {
		fieldsAttrs := obj.ProtoReflect().Descriptor().Fields()
		for i := 0; i < fieldsAttrs.Len(); i++ {
			name := fieldsAttrs.Get(i).TextName()
			if !convert.ContainIgnoreCase(name, excludeFields) {
				includeFields = append(includeFields, name)
			}
		}
		info := make(map[string]interface{}, cap(includeFields))
		for _, field := range includeFields {
			info[field] = obj.ProtoReflect().Get(fieldsAttrs.ByName(protoreflect.Name(field))).Interface()
		}
		data = append(data, info)
	}
	return data
}

func CropObjFields(obj any, fields []string) {
	object := reflect.ValueOf(obj)
	elems := object.Elem()
	typeOfType := elems.Type()
	for i := 0; i < elems.NumField(); i++ {
		field := elems.Field(i)
		fieldName := typeOfType.Field(i).Name
		if !convert.ContainIgnoreCase(fieldName, fields) {
			kind(field)
		}
	}
}

func kind(field reflect.Value) {
	switch field.Kind() {
	case reflect.String:
		field.SetString("")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		field.SetInt(0)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		field.SetUint(0)
	case reflect.Float32, reflect.Float64:
		field.SetFloat(0)
	case reflect.Bool:
		field.SetBool(false)
	default:
		panic("unhandled default case")
	}
}
