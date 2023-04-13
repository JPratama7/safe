package safe

import (
	"github.com/goccy/go-reflect"
	refdef "reflect"
)

func NotEmpty(data any) (res bool) {
	val := reflect.ToReflectValue(reflect.ValueNoEscapeOf(data))
	typeOf := val.Type()
	valDef := refdef.Zero(typeOf)
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Chan, reflect.Slice, reflect.Map:
		res = !val.IsNil()
		return
	case reflect.Array, reflect.Struct:
		res = val.Interface() != valDef.Interface()
		return
	//case reflect.String:
	//	res = val != valDef
	//	return
	default:
		res = val != valDef
		return
		//default:
		//	res = !val.IsZero()
		//	return
	}
}
