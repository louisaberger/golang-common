package util

import (
	"fmt"
	"reflect"
	"strings"
)

func StructIndex(key string, v reflect.Value) (val reflect.Value, exists bool) {

	if !IsPtrToStruct(v) && !IsStruct(v) {
		panic(fmt.Sprintf("Cannot call StructIndex on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	if FirstCharUpper(key) != key {
		panic(fmt.Sprintf("Cannot get a non-exported field=%v from struct.", key))
	}

	if !StructFieldExists(key, v) {
		return
	}

	switch {
	case IsPtrToStruct(v):
		return v.Elem().FieldByName(key), true
	case IsStruct(v):
		return v.FieldByName(key), true
	default:
		return
	}
}

func StructFieldExists(key string, v reflect.Value) bool {
	if !IsPtrToStruct(v) && !IsStruct(v) {
		panic(fmt.Sprintf("Cannot call StructFieldExists on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	var ok bool
	switch {
	case IsPtrToStruct(v):
		_, ok = v.Elem().Type().FieldByName(key)
	case IsStruct(v):
		_, ok = v.Type().FieldByName(key)
	}
	return ok
}

func NestedStructIndex(key string, v reflect.Value) (val reflect.Value, exists bool) {
	obj := v

	// sanitize the key input
	key = strings.TrimSuffix(key, ".")

	for strings.Index(key, ".") != -1 {
		if !IsPtrToStruct(obj) && !IsStruct(obj) {
			panic(fmt.Sprintf("Cannot call NestedStructIndex on a non-struct %#v of kind %#v", v, v.Kind().String()))
		}

		dotIndex := strings.Index(key, ".")
		keyBeforeDot := key[:dotIndex]
		key = key[dotIndex+1:]

		if FirstCharUpper(keyBeforeDot) != keyBeforeDot {
			panic(fmt.Sprintf("Cannot get a non-exported field=%v from struct.", keyBeforeDot))
		}

		var nestedObj reflect.Value
		nestedObj, exists = NestedStructIndex(keyBeforeDot, obj)
		if !exists {
			return
		}

		if TypeIsPtrToStruct(structFieldType(keyBeforeDot, obj)) && !nestedObj.Elem().IsValid() {
			return
		}
		obj = nestedObj
	}

	return StructIndex(key, obj)
}

func structFieldType(key string, v reflect.Value) (t reflect.Type) {
	if !IsPtrToStruct(v) && !IsStruct(v) {
		panic(fmt.Sprintf("Cannot call structFieldType on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	var structField reflect.StructField
	var ok bool
	switch {
	case IsPtrToStruct(v):
		structField, ok = v.Elem().Type().FieldByName(key)
	case IsStruct(v):
		structField, ok = v.Type().FieldByName(key)
	}
	if !ok {
		panic(fmt.Sprintf("Key=%v is not a field in struct=%v", key, v))
	}
	return structField.Type
}
