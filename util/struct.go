package util

import (
	"fmt"
	"reflect"
	"strings"
)

// @param obj should be the value of a ptr to a struct
// @return reflect.Value of the 'key' field of 'obj'
func StructIndex(key string, obj reflect.Value) reflect.Value {

	if !IsPtrToStruct(obj) {
		panic(fmt.Sprintf("Cannot call StructIndex on a non-ptr to struct %#v of kind %#v", obj, obj.Kind().String()))
	}

	pointingTo := obj.Elem()
	if !fieldExists(key, pointingTo) {
		return reflect.ValueOf(nil)
	}
	// get the actual field value
	fieldValue := pointingTo.FieldByName(strings.Title(key))
	return fieldValue
}

func fieldExists(key string, v reflect.Value) bool {
	_, ok := v.Type().FieldByName(strings.Title(key))
	return ok
}
