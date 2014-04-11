package util

import (
	"reflect"
)

func TypeIsStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func TypeIsFunc(t reflect.Type) bool {
	return t.Kind() == reflect.Func
}

func TypeIsSlice(t reflect.Type) bool {
	return t.Kind() == reflect.Slice
}

func TypeIsMap(t reflect.Type) bool {
	return t.Kind() == reflect.Map
}

func TypeIsPtrToMap(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Ptr:
		pointingToType := t.Elem()
		return TypeIsMap(pointingToType)
	}
	return false
}

func TypeIsPtrToStruct(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Ptr:
		pointingToType := t.Elem()
		return TypeIsStruct(pointingToType)
	}
	return false
}

func IsZero(v reflect.Value) bool {
	return v.Kind() == reflect.Invalid
}

func IsStruct(v reflect.Value) bool {
	return v.Kind() == reflect.Struct
}

func IsFunc(v reflect.Value) bool {
	return v.Kind() == reflect.Func
}

func IsSlice(v reflect.Value) bool {
	return v.Kind() == reflect.Slice
}

func IsMap(v reflect.Value) bool {
	return v.Kind() == reflect.Map
}

func IsPtr(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr
}

func IsPtrToMap(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr:
		pointingToObj := v.Elem()
		return IsMap(pointingToObj)
	}
	return false
}

func IsPtrToStruct(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr:
		pointingToObj := v.Elem()
		return IsStruct(pointingToObj)
	}
	return false
}

func IsPtrStructOrStruct(v reflect.Value) bool {
	return IsPtrToStruct(v) || IsStruct(v)
}

func IsSliceOfString(v reflect.Value) bool {
	if !IsSlice(v) {
		return false
	}

	for idx := 0; idx < v.Len(); idx++ {
		_, ok := v.Index(idx).Interface().(string)
		if !ok {
			return false
		}
	}
	return true
}

func CanConvert(v reflect.Value, t reflect.Type) (val reflect.Value, canConvert bool) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	converted := v.Convert(t)
	return converted, true
}
