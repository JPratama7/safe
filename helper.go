package safe

import (
	"fmt"
	"github.com/goccy/go-reflect"
	"math"
	refdef "reflect"
)

func NotEmpty(data any) (res bool) {
	val := reflect.ToReflectValue(reflect.ValueNoEscapeOf(data))
	valDef := refdef.Zero(val.Type())
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
	case refdef.Bool:
		res = val.Bool()
		return
	case refdef.Int, refdef.Int8, refdef.Int16, refdef.Int32, refdef.Int64:
		res = val.Int() != 0
		return
	case refdef.Uint, refdef.Uint8, refdef.Uint16, refdef.Uint32, refdef.Uint64, refdef.Uintptr:
		res = val.Uint() != 0
		return
	case refdef.Float32, refdef.Float64:
		res = math.Float64bits(val.Float()) != 0
		return
	case refdef.Complex64, refdef.Complex128:
		c := val.Complex()
		res = math.Float64bits(real(c)) != 0 && math.Float64bits(imag(c)) != 0
		return
	case refdef.String:
		res = val != valDef
		return
	default:
		panic(fmt.Errorf("unsupported type %T, a %s", data, val.Kind()))
	}
}
