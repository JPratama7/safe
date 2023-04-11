package safe

import "github.com/goccy/go-reflect"

func Checker(val reflect.Value) (res bool) {
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		res = !val.IsNil()
		break
	case reflect.Array:
		res = val.Len() > 0
		break
	case reflect.Struct, reflect.String:
		res = val.Interface() != reflect.Zero(val.Type()).Interface()
		break
	default:
		res = !val.IsZero()
		break
	}
	return res
}
