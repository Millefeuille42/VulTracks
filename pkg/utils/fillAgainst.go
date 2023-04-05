package utils

import (
	"reflect"
)

// FillAgainst fills the first struct with its empty fields filled with the second structs fields
func FillAgainst(t1, t2 interface{}) {
	e1 := reflect.ValueOf(t1).Elem()
	e2 := reflect.ValueOf(t2).Elem()

	if e1.Kind() != reflect.Struct || e2.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < e2.NumField(); i++ {
		if e1.Field(i).IsZero() {
			e1.Field(i).Set(e2.Field(i))
		}
	}
}
