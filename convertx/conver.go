package convertx

import (
	"github.com/duke-git/lancet/v2/convertor"
)

func AnyToInt64(value any) (int64, error) {
	return convertor.ToInt(value)
}

func AnyToArr[T comparable](arr []T) []any {
	var t []any
	for _, item := range arr {
		t = append(t, item)
	}
	return t
}

func ToBytes(value any) ([]byte, error) {
	return convertor.ToBytes(value)
}
