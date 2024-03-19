package pointx

func SetPointer[T comparable](value T) *T {
	return &value
}

func GetPointer[T comparable](value *T) T {
	return *value
}
