package safe

import (
	"github.com/goccy/go-reflect"
	refdef "reflect"
)

func NotEmpty(data any) (res bool) {
	val := reflect.ToReflectValue(reflect.ValueNoEscapeOf(data))
	typeOf := val.Type()
	valDef := refdef.Zero(typeOf)
	//fmt.Printf("value Croot zeroed : %+v \n", valDef)
	//fmt.Printf("value Croot : %+v \n", val)
	//fmt.Printf("value Croot compare : %+v \n", val.Equal(valDef))
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case refdef.Chan, refdef.Slice, refdef.Map, refdef.Func, refdef.Pointer, refdef.UnsafePointer, refdef.Interface:
		res = !val.IsNil()
		return
	case refdef.Array, refdef.Struct:
		res = val.Interface() != valDef.Interface()
		return
	case refdef.String:
		res = val != valDef
		return
	default:
		res = !val.IsZero()
		return
		//default:
		//	return
	}
}
