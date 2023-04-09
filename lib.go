package safe

import "github.com/goccy/go-reflect"

func Checker(val reflect.Value) (res bool) {
	switch val.Kind() {
	case reflect.Chan, reflect.Slice, reflect.Map:
		res = !val.IsNil()
		break
	case reflect.Array, reflect.String:
		res = val.Len() > 0
		break
	default:
		res = val.IsValid() && !val.IsZero()
		break
	}
	return res
}
