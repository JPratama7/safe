package safe

import (
	"github.com/goccy/go-reflect"
	refdef "reflect"
)

func NotEmpty(data any) (res bool) {
	val := reflect.ToReflectValue(reflect.ValueNoEscapeOf(data))
	valdef := refdef.Zero(val.Type()).Interface()
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Chan, reflect.Slice, reflect.Map:
		res = !val.IsNil()
		return
	case reflect.Array, reflect.Struct:
		res = val.Interface() != valdef
		return
	case reflect.String:
		res = val != refdef.Zero(val.Type())
		return
	default:
		res = !val.IsZero()
		return
	}
}
