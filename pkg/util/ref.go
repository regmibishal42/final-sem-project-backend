package util

import "reflect"

func Ref[T any](t T) *T {
	return &t
}

func Deref[T any](t *T) T {
	refType := reflect.TypeOf(t)
	if t == nil {
		switch refType.String() {
		case "*string":
			return any("").(T)
		case "*int":
			return any(0).(T)
		case "*bool":
			return any(false).(T)
		}
	}
	return *t
}
