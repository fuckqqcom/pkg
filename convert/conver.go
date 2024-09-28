package convert

import (
	"github.com/duke-git/lancet/v2/convertor"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateStr(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

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
func AnyToStr(value any) string {
	return convertor.ToString(value)
}
func ArrToJoinStr(args ...any) string {
	var s []string
	for _, arg := range args {
		s = append(s, convertor.ToString(arg))
	}
	return strings.Join(s, "_")
}
func ToBytes(value any) ([]byte, error) {
	return convertor.ToBytes(value)
}

func ContainIgnoreCase(target string, arr []string) bool {
	target = strings.ToLower(target)

	for _, item := range arr {
		if strings.ToLower(item) == target {
			return true
		}
	}
	return false
}

func Contains[T comparable](target T, arr []T) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

func MapErrs[T comparable](m map[T]error) (errs []error) {
	for _, v := range errs {
		errs = append(errs, v)
	}
	return
}

func ReflectSlice(i interface{}) []interface{} {
	if i == nil {
		return []interface{}{}
	}

	switch v := i.(type) {
	case []interface{}:
		return v
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]interface{}, s.Len())
		for i := range a {
			a[i] = s.Index(i).Interface()
		}
		return a
	case
		reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.String:
		return []interface{}{i}
	default:
		return []interface{}{}
	}
}
