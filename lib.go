package safe

import (
	"github.com/goccy/go-reflect"
	"math"
)

func Checker(val reflect.Value) (res bool) {
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		res = !val.IsNil()
		return
	case reflect.Array:
		res = val.Len() > 0
		return
	case reflect.Struct:
		res = val.Interface() != reflect.Zero(val.Type()).Interface()
		return
	case reflect.String:
		res = val != reflect.Zero(val.Type())
		return
	case reflect.Bool:
		res = !val.Bool()
		return
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res = val.Int() != 0
		return
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		res = val.Uint() != 0
		return
	case reflect.Float32, reflect.Float64:
		res = math.Float64bits(val.Float()) != 0
		return
	case reflect.Complex64, reflect.Complex128:
		c := val.Complex()
		res = math.Float64bits(real(c)) != 0 && math.Float64bits(imag(c)) != 0
		return
	default:
		return
	}
}
