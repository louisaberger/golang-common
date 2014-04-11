package util

import (
	"fmt"
	"reflect"
	"strings"
)

func StructIndex(key string, v reflect.Value) (val reflect.Value, exists bool) {

	if !IsPtrStructOrStruct(v) {
		panic(fmt.Sprintf("Cannot call StructIndex on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	if FirstCharUpper(key) != key {
		panic(fmt.Sprintf("Cannot get a non-exported field=%v from struct.", key))
	}

	if !StructFieldExists(key, v) {
		return
	}

	return StructIndirect(v).FieldByName(key), true
}

func NestedStructIndex(key string, v reflect.Value) (val reflect.Value, exists bool) {
	obj := v

	// sanitize the key input
	key = strings.TrimSuffix(key, ".")

	for strings.Index(key, ".") != -1 {
		if !IsPtrStructOrStruct(obj) {
			panic(fmt.Sprintf("Cannot call NestedStructIndex on a non-struct %#v of kind %#v", obj, obj.Kind().String()))
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
		if TypeIsPtrToStruct(StructFieldType(keyBeforeDot, obj)) && !nestedObj.Elem().IsValid() {
			return
		}
		obj = nestedObj
	}

	return StructIndex(key, obj)
}

func StructFieldExists(key string, v reflect.Value) bool {
	if !IsPtrStructOrStruct(v) {
		panic(fmt.Sprintf("Cannot call StructFieldExists on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	_, ok := StructIndirect(v).Type().FieldByName(key)
	return ok
}

func StructFieldType(key string, v reflect.Value) (t reflect.Type) {
	if !IsPtrStructOrStruct(v) {
		panic(fmt.Sprintf("Cannot call StructFieldType on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	structField, ok := StructIndirect(v).Type().FieldByName(key)
	if !ok {
		panic(fmt.Sprintf("Key=%v is not a field in struct=%v", key, v))
	}
	return structField.Type
}

func StructIndirect(v reflect.Value) (strct reflect.Value) {
	if !IsPtrStructOrStruct(v) {
		panic(fmt.Sprintf("Cannot call StructIndirect on a non-struct %#v of kind %#v", v, v.Kind().String()))
	}

	switch {
	case IsPtrToStruct(v):
		return v.Elem()
	case IsStruct(v):
		return v
	}

	return
}
