package safe

import (
	"github.com/goccy/go-reflect"
	refdef "reflect"
)

func NotEmpty(data any) (res bool) {
	val := reflect.ToReflectValue(reflect.ValueNoEscapeOf(data))
	typeOf := val.Type()
	valDef := refdef.Zero(typeOf).Interface()
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Chan, reflect.Slice, reflect.Map:
		res = !val.IsNil()
		return
	case reflect.Array, reflect.Struct:
		res = val.Interface() != valDef
		return
	case reflect.String:
		res = val.Len() >= 1
		return
	default:
		res = !val.IsZero()
		return
	}
}
