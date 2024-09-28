package metadata

import (
	"github.com/fuckqqcom/pkg/convert"
	"reflect"
)

func GetTagNames(obj interface{}, excludes []string) (fields []string) {
	s := reflect.TypeOf(obj).Elem()
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i).Tag.Get("json")
		if field == "" || convert.ContainIgnoreCase(field, excludes) == true {
			continue
		}
		fields = append(fields, field)
	}
	return
}
