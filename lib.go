package safetypes

import "github.com/goccy/go-reflect"

func isEmpty(x any) bool {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0.0
	case reflect.Bool:
		return !v.Bool()
	default:
		return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
	}
}
