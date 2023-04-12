package safe

import (
	"github.com/goccy/go-reflect"
)

func Checker(val reflect.Value) (res bool) {
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		res = !val.IsNil()
		return
	case reflect.Struct, reflect.Array:
		res = val.Interface() != reflect.Zero(val.Type()).Interface()
		return
	case reflect.String:
		res = val != reflect.Zero(val.Type())
		return
	default:
		res = !val.IsZero()
		return
	}
}
