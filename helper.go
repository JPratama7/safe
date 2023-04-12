package safe

import (
	"github.com/goccy/go-reflect"
	refdef "reflect"
)

func NotEmpty(data any) (res bool) {
	val := reflect.ToReflectValue(reflect.ValueNoEscapeOf(data))
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case refdef.Chan, refdef.Slice, refdef.Map:
		res = !val.IsNil()
		return
	case refdef.Array, refdef.Struct:
		res = val.Interface() != refdef.Zero(val.Type()).Interface()
		return
	case refdef.String:
		res = val != refdef.Zero(val.Type())
	default:
		res = !val.IsZero()
		return
	}
	return
}
