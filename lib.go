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
	case reflect.String:
		res = val != reflect.Zero(val.Type())
		break
	case reflect.Struct:
		res = val != reflect.Zero(val.Type())
		break
	default:
		res = !val.IsZero()
		break
	}
	return res
}
