package util

import (
	. "github.com/smartystreets/goconvey/convey"

	"fmt"
	"reflect"
	"testing"
)

type AStruct struct {
	Exported        string
	unexported      string
	NestedStruct    BStruct
	PtrNestedStruct *BStruct
}

type BStruct struct {
	Exported     string
	unexported   string
	NestedStruct *CStruct
}

type CStruct struct {
	Exported string
}

func getStruct1() *AStruct {
	return &AStruct{
		"v1",
		"v2",
		BStruct{"b1", "b2", &CStruct{"c1"}},
		&BStruct{"b3", "b4", &CStruct{"c1"}},
	}
}

func TestStructIndex(t *testing.T) {
	struct1 := getStruct1()
	struct1Val := reflect.ValueOf(struct1)
	Convey("When calling StructIndex", t, func() {

		check := func(key string, shouldExist bool, expected interface{}) {
			val, exists := StructIndex(key, struct1Val)
			So(exists, ShouldEqual, shouldExist)
			if shouldExist {
				So(val.Interface(), ShouldDeepEqual, expected)
			}
		}

		Convey("You should be able to get an exported string field", func() {
			check("Exported", true, "v1")
		})
		Convey("You should not be able to get an unexported field", func() {
			So(func() { _, _ = StructIndex("unexported", struct1Val) }, ShouldPanic)
		})
		Convey("You should be able to get an exported struct field", func() {
			check("NestedStruct", true, struct1.NestedStruct)
		})
		Convey("You should be able to get an exported ptr to struct field", func() {
			check("PtrNestedStruct", true, struct1.PtrNestedStruct)
		})
		Convey("Exists should return false for a non-existent field", func() {
			check("Nonexistent", false, "")
		})
		Convey("It should panic on a non-struct input", func() {
			So(func() { _, _ = StructIndex("field", reflect.ValueOf("")) }, ShouldPanic)
		})

	})
}

func TestNestedStructIndex(t *testing.T) {
	struct1 := getStruct1()
	struct1Val := reflect.ValueOf(struct1)
	Convey("When calling NestedStructIndex", t, func() {

		check := func(key string, shouldExist bool, expected interface{}) {
			val, exists := NestedStructIndex(key, struct1Val)
			So(exists, ShouldEqual, shouldExist)
			if shouldExist {
				So(safeInterface(val), ShouldDeepEqual, expected)
			}
		}

		Convey("You should be able to get a top-level field", func() {
			check("Exported", true, struct1.Exported)
		})
		Convey("You should be able to get a 2nd-level field from a struct", func() {
			check("NestedStruct.Exported", true, struct1.NestedStruct.Exported)
		})
		Convey("You should be able to get a 2nd-level field from a ptr to a struct", func() {
			check("PtrNestedStruct.Exported", true, struct1.PtrNestedStruct.Exported)
		})
		Convey("Should panic trying to get a 2nd-level unexported field from a struct", func() {
			So(func() { _, _ = StructIndex("NestedField.unexported", reflect.ValueOf("")) }, ShouldPanic)
		})
		Convey("Should panic trying to get a 2nd-level unexported field from a ptr to a struct", func() {
			So(func() { _, _ = StructIndex("PtrNestedField.unexported", reflect.ValueOf("")) }, ShouldPanic)
		})
		Convey("Should panic trying to get a nested field from a non-struct", func() {
			So(func() { _, _ = StructIndex("Exported.Nonexistent", reflect.ValueOf("")) }, ShouldPanic)
		})
		Convey("You should be able to get a 3rd-level field from a struct", func() {
			check("NestedStruct.NestedStruct.Exported", true, struct1.NestedStruct.NestedStruct.Exported)
		})
		Convey("You should be able to get a 3rd-level field from a ptr to a struct", func() {
			check("PtrNestedStruct.NestedStruct.Exported", true, struct1.PtrNestedStruct.NestedStruct.Exported)
		})
		Convey("You should get exists=false for nested fields if the outer struct is nil", func() {
			struct1.PtrNestedStruct = nil
			check("PtrNestedStruct", true, struct1.PtrNestedStruct)
			check("PtrNestedStruct.Exported", true, nil)
		})

	})
}

func TestSetStructIndex(t *testing.T) {
	struct1 := new(AStruct)
	struct1Val := reflect.ValueOf(struct1)
	Convey("When calling SetStructIndex", t, func() {

		check := func(key string, x interface{}) {
			SetStructIndex(key, reflect.ValueOf(x), struct1Val)
			res, _ := StructIndex(key, reflect.ValueOf(struct1))
			So(safeInterface(res), ShouldDeepEqual, x)
		}

		Convey("You should be able to set an exported string field", func() {
			check("Exported", "v1")
		})
		Convey("You should be able to reset a field", func() {
			check("Exported", "v2")
		})
		Convey("You should not be able to set an unexported field", func() {
			So(func() { SetStructIndex("unexported", reflect.ValueOf(""), struct1Val) }, ShouldPanic)
		})
		Convey("You should be able to set an exported struct field", func() {
			check("NestedStruct", BStruct{"v1", "v2", &CStruct{"v3"}})
		})
		Convey("You should be able to set an exported ptr to struct field", func() {
			check("PtrNestedStruct", &BStruct{"v1", "v2", &CStruct{"v3"}})
		})
		Convey("It should panic on a non-struct input", func() {
			So(func() { SetStructIndex("Exported", reflect.ValueOf(""), reflect.ValueOf("")) }, ShouldPanic)
		})
		Convey("It should panic setting a non-existent field", func() {
			So(func() { SetStructIndex("nonexistent", reflect.ValueOf(""), struct1Val) }, ShouldPanic)
		})

	})
}

func TestSetNestedStructIndex(t *testing.T) {
	struct1 := new(AStruct)
	struct1Val := reflect.ValueOf(struct1)
	Convey("When calling SetNestedStructIndex", t, func() {

		check := func(key string, x interface{}) {
			SetNestedStructIndex(key, reflect.ValueOf(x), struct1Val)
			res, _ := NestedStructIndex(key, reflect.ValueOf(struct1))
			So(safeInterface(res), ShouldDeepEqual, x)
		}

		Convey("You should be able to set a top-level field", func() {
			check("Exported", "v1")
		})
		Convey("You should be able to set a 2nd-level field in an un-set struct", func() {
			SetNestedStructIndex("NestedStruct.Exported", reflect.ValueOf("v1"), struct1Val)
			res, _ := NestedStructIndex("NestedStruct", reflect.ValueOf(struct1))
			So(safeInterface(res), ShouldDeepEqual, BStruct{Exported: "v1"})

		})
		Convey("You should be able to set a 2nd-level field in a nil struct ptr", func() {
			struct1.PtrNestedStruct = nil
			SetNestedStructIndex("PtrNestedStruct.Exported", reflect.ValueOf("v1"), struct1Val)
			res, _ := NestedStructIndex("PtrNestedStruct", reflect.ValueOf(struct1))
			So(safeInterface(res), ShouldDeepEqual, &BStruct{Exported: "v1"})

		})
		// Convey("You should be able to get a 2nd-level field from a ptr to a struct", func() {
		// 	check("PtrNestedStruct.Exported", true, struct1.PtrNestedStruct.Exported)
		// })
		// Convey("Should panic trying to get a 2nd-level unexported field from a struct", func() {
		// 	So(func() { _, _ = StructIndex("NestedField.unexported", reflect.ValueOf("")) }, ShouldPanic)
		// })
		// Convey("Should panic trying to get a 2nd-level unexported field from a ptr to a struct", func() {
		// 	So(func() { _, _ = StructIndex("PtrNestedField.unexported", reflect.ValueOf("")) }, ShouldPanic)
		// })
		// Convey("Should panic trying to get a nested field from a non-struct", func() {
		// 	So(func() { _, _ = StructIndex("Exported.Nonexistent", reflect.ValueOf("")) }, ShouldPanic)
		// })
		// Convey("You should be able to get a 3rd-level field from a struct", func() {
		// 	check("NestedStruct.NestedStruct.Exported", true, struct1.NestedStruct.NestedStruct.Exported)
		// })
		// Convey("You should be able to get a 3rd-level field from a ptr to a struct", func() {
		// 	check("PtrNestedStruct.NestedStruct.Exported", true, struct1.PtrNestedStruct.NestedStruct.Exported)
		// })
		// Convey("You should get exists=false for nested fields if the outer struct is nil", func() {
		// 	struct1.PtrNestedStruct = nil
		// 	check("PtrNestedStruct", true, struct1.PtrNestedStruct)
		// 	check("PtrNestedStruct.Exported", true, nil)
		// })

	})
}

func safeInterface(v reflect.Value) (res interface{}) {
	if !v.IsValid() {
		return nil
	}
	return v.Interface()
}

func ShouldDeepEqual(actual interface{}, expected ...interface{}) string {
	if reflect.DeepEqual(actual, expected[0]) {
		return "" // empty string means the assertion passed
	} else {
		return fmt.Sprintf("Actual=%v(%T) does not deep equal expected=%v(%T)", actual, actual, expected[0], expected[0])
	}
}
