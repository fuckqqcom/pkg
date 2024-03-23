package metadatax

import (
	"github.com/fuckqqcom/pkg/convertx"
	"reflect"
)

func GetTagNames(obj interface{}, excludes []string) (fields []string) {
	s := reflect.TypeOf(obj).Elem()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i).Tag.Get("json")
		if field == "" || convertx.ContainIgnoreCase(field, excludes) == true {
			continue
		}
		fields = append(fields, field)
	}
	return
}
