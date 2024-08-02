package helpers

import "reflect"

func IsSliceOrArray(value interface{}) bool {
	v := reflect.ValueOf(value)
	vType := v.Kind()
	return vType == reflect.Array || vType == reflect.Slice
}
