package convertx

import (
	"github.com/duke-git/lancet/v2/convertor"
	"math/rand"
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

func MapVal[T comparable](m map[T]T) (val []any) {
	for _, v := range val {
		val = append(val, v)
	}
	return
}
