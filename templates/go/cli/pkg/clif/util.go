package clif

import "reflect"

// interfaceToSlice converts an interface that is of underlying type
// slice, array, or map to a slice of interfaces instead of a single interface.
// Panics if the interface passed is not a slice, array, or map.
func interfaceToSlice(ifc interface{}) []interface{} {
	s := reflect.ValueOf(ifc)
	k := s.Kind()

	ret := make([]interface{}, s.Len())

	isSliceOrArray := k == reflect.Slice || k == reflect.Array
	if isSliceOrArray {

		for i:=0; i<s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
		}

		return ret
	}

	isMap := k == reflect.Map
	if isMap {

		iter := s.MapRange()
		count := 0
		for iter.Next() {
			val := iter.Value()
			ret[count] = val.Interface()
			count++
		}

		return ret
	}

	panic("interfaceToSlice() only takes a slice, array, or map")
}