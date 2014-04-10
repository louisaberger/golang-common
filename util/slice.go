package util

import (
	"fmt"
	"reflect"
)

// @return true if elt is in slice.
// panics if slice is not of Kind reflect.Slice
func SliceContains(slice, elt interface{}) bool {
	if slice == nil {
		return false
	}
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Cannot call SliceContains on a non-slice %#v of kind %#v", slice, v.Kind().String()))
	}
	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), elt) {
			return true
		}
	}
	return false
}

// @return number of instances of 'elt' in 'slice'.
// panics if slice is not of Kind reflect.Slice
func SliceCount(slice, elt interface{}) int {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Cannot call SliceCount on a non-slice %#v of kind %#v", slice, v.Kind().String()))
	}
	counter := 0
	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), elt) {
			counter++
		}
	}
	return counter
}

// @return true if s1 and s2 have exactly the same elements, in any order.
// panics if s1 or s2 are not of Kind reflect.Slice
func SliceUnorderedEqual(s1, s2 interface{}) bool {
	v1 := reflect.ValueOf(s1)
	v2 := reflect.ValueOf(s2)

	if !v1.IsValid() || !v2.IsValid() {
		return !v1.IsValid() && !v2.IsValid()
	}

	if v1.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Cannot call SliceUnorderedEqual on a non-slice %#v of kind %#v", s1, v1.Kind().String()))
	}
	if v2.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Cannot call SliceUnorderedEqual on a non-slice %#v of kind %#v", s2, v2.Kind().String()))
	}

	if v1.Len() != v2.Len() {
		return false
	}
	for i := 0; i < v1.Len(); i++ {
		elt := v1.Index(i).Interface()
		s1Count := SliceCount(s1, elt)
		s2Count := SliceCount(s2, elt)
		if s1Count != s2Count {
			return false
		}
	}
	return true
}

// @return slice with the element at idx removed.
// panics if slice is not of Kind reflect.Slice
func SliceRemove(slice interface{}, idx int) (modSlice interface{}) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Cannot call SliceRemove on a non-slice %#v of kind %#v", slice, v.Kind().String()))
	}

	if idx >= v.Len() {
		panic(fmt.Sprintf("Index=%v for SliceRemove is out of bounds (slice=%v has len=%v)", idx, slice, v.Len()))
	}

	return reflect.AppendSlice(v.Slice(0, idx), v.Slice(idx+1, v.Len())).Interface()
}

// @return slice with all instances of elt removed.
// panics if slice is not of Kind reflect.Slice
func SliceRemoveElt(slice interface{}, elt interface{}) (modSlice interface{}) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Cannot call SliceRemoveElt on a non-slice %#v of kind %#v", slice, v.Kind().String()))
	}

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), elt) {
			v = reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len()))
			i-- // have removed the previous elt at 'i'
		}
	}

	return v.Interface()
}
