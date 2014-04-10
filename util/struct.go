package util

import (
	"fmt"
	"reflect"
)

func StructIndex(key string, strc reflect.Value) (val reflect.Value, exists bool) {

	if !IsStruct(strc) {
		panic(fmt.Sprintf("Cannot call StructIndex on a non-struct %#v of kind %#v", strc, strc.Kind().String()))
	}

	if !StructFieldExists(key, strc) {
		return
	}

	fieldValue := strc.FieldByName(key)
	return fieldValue, true
}

func StructFieldExists(key string, strc reflect.Value) bool {
	if !IsStruct(strc) {
		panic(fmt.Sprintf("Cannot call StructFieldExists on a non-struct %#v of kind %#v", strc, strc.Kind().String()))
	}

	_, ok := strc.Type().FieldByName(key)
	return ok
}
