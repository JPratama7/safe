package safe

import (
	"math"
)

//func NotEmpty(data any) (res bool) {
//	val := reflect.ToReflectValue(reflect.ValueNoEscapeOf(data))
//	typeOf := val.Type()
//	valDef := refdef.Zero(typeOf)
//	if !val.IsValid() {
//		return
//	}
//	switch val.Kind() {
//	case refdef.Chan, refdef.Slice, refdef.Map, refdef.Func, refdef.Pointer, refdef.UnsafePointer, refdef.Interface:
//		res = !val.IsNil()
//		return
//	case refdef.Array, refdef.Struct:
//		res = val.Interface() != valDef.Interface()
//		return
//	case refdef.Bool:
//		res = val.Bool()
//		return
//	case refdef.Int, refdef.Int8, refdef.Int16, refdef.Int32, refdef.Int64:
//		res = val.Int() != 0
//		return
//	case refdef.Uint, refdef.Uint8, refdef.Uint16, refdef.Uint32, refdef.Uint64, refdef.Uintptr:
//		res = val.Uint() != 0
//		return
//	case refdef.Float32, refdef.Float64:
//		res = math.Float64bits(val.Float()) != 0
//		return
//	case refdef.Complex64, refdef.Complex128:
//		c := val.Complex()
//		res = math.Float64bits(real(c)) != 0 && math.Float64bits(imag(c)) != 0
//		return
//	case refdef.String:
//		res = val != valDef
//		return
//	default:
//		panic(fmt.Errorf("unsupported type %T, a %s", data, val.Kind()))
//	}
//}

func NotEmpty(data any) bool {
	switch data := data.(type) {
	case string:
		return data != ""
	case int, int8, int16, int32, int64:
		return data != 0
	case uint, uint8, uint16, uint32, uint64, uintptr:
		return data != 0
	case float32, float64:
		return math.Float64bits(data.(float64)) != 0
	case complex64, complex128:
		c := data.(complex128)
		return math.Float64bits(real(c)) != 0 && math.Float64bits(imag(c)) != 0
	case bool:
		return data
	case []any:
		return len(data) > 0
	case map[any]any:
		return len(data) > 0
	case nil:
		return false
	default:
		return true
	}
}
