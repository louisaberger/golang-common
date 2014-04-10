package util

import (
	"fmt"
	"reflect"
)

func MapHasKey(m interface{}, key string) bool {
	v := reflect.ValueOf(m)
	if !IsMap(v) {
		panic(fmt.Sprintf("Cannot call MapHasKey on a non-map %#v of kind %#v", m, v.Kind().String()))
	}

	for _, k := range v.MapKeys() {
		if reflect.DeepEqual(k.Interface(), key) {
			return true
		}
	}
	return false
}
