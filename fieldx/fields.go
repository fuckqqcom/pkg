package fieldx

import (
	"github.com/fuckqqcom/pkg/convertx"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
)

func CropProtoFields(obj proto.Message, includeFields []string, excludeFields []string, merge bool) map[string]interface{} {
	fieldsAttrs := obj.ProtoReflect().Descriptor().Fields()
	for i := 0; i < fieldsAttrs.Len(); i++ {
		name := fieldsAttrs.Get(i).TextName()
		if merge && !convertx.ContainIgnoreCase(name, excludeFields) {
			includeFields = append(includeFields, name)
		} else if includeFields == nil {
			includeFields = append(includeFields, name)
		}
	}
	data := make(map[string]interface{}, len(includeFields))
	for _, field := range includeFields {
		data[field] = convertx.AnyToStr(obj.ProtoReflect().Get(fieldsAttrs.ByName(protoreflect.Name(field))).Interface())
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
			if !convertx.ContainIgnoreCase(name, excludeFields) {
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
		if !convertx.ContainIgnoreCase(fieldName, fields) {
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
