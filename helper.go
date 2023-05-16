package safe

import (
	"fmt"
	"github.com/goccy/go-reflect"
	"math"
	refdef "reflect"
	"unsafe"
)

func reflectValue(val refdef.Value) (res bool) {
	valDef := refdef.Value{}
	if val.Kind() != refdef.Pointer {
		valDef = refdef.Zero(val.Type())
	}

	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case refdef.Chan, refdef.Func, refdef.Pointer, refdef.UnsafePointer, refdef.Interface:
		res = !val.IsNil()
		return
	case refdef.Array, refdef.Struct:
		res = val.Interface() != valDef.Interface()
		return
	case refdef.Slice, refdef.Map:
		res = val.Len() > 0
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
		panic(fmt.Errorf("unsupported type a %v", val.Kind()))
	}
}

func IsNotEmpty(data any) bool {
	switch val := data.(type) {
	case func(), *interface{}, unsafe.Pointer:
		return val != nil
	case bool:
		return val
	case int, int8, int16, int32, int64:
		return val != 0
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return val != 0
	case float32, float64:
		conv := val.(float64)
		return math.Float64bits(conv) != 0
	case complex64, complex128:
		conv := val.(complex128)
		return math.Float64bits(real(conv)) != 0 && math.Float64bits(imag(conv)) != 0
	case string:
		return val != ""
	default:
		refl := reflect.ToReflectValue(reflect.ValueNoEscapeOf(val))
		return reflectValue(refl)
	}
}
