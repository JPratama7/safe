package safe

import (
	refgo "github.com/goccy/go-reflect"
)

func Checker(val refgo.Value) (res bool) {
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case refgo.Chan, refgo.Func, refgo.Interface, refgo.Map, refgo.Ptr, refgo.Slice, refgo.UnsafePointer:
		res = !val.IsNil()
		return
	case refgo.Struct, refgo.Array:
		res = val.Interface() != refgo.Zero(val.Type()).Interface()
		return
	default:
		res = !val.IsZero()
		return
	}
}
