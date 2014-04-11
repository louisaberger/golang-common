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

	return structFromPtrOrStruct(v).FieldByName(key), true
}

func StructFieldExists(key string, v reflect.Value) bool {
	if !IsPtrToStruct(v) && !IsStruct(v) {
		panic(fmt.Sprintf("Cannot call StructFieldExists on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	_, ok := structFromPtrOrStruct(v).Type().FieldByName(key)
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

		// nestedObj is a nil ptr to struct
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

	structField, ok := structFromPtrOrStruct(v).Type().FieldByName(key)
	if !ok {
		panic(fmt.Sprintf("Key=%v is not a field in struct=%v", key, v))
	}
	return structField.Type
}

func structFromPtrOrStruct(v reflect.Value) (strct reflect.Value) {
	if !IsPtrToStruct(v) && !IsStruct(v) {
		panic(fmt.Sprintf("Cannot call structFromPtrOrStruct on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	switch {
	case IsPtrToStruct(v):
		return v.Elem()
	case IsStruct(v):
		return v
	}

	return
}
