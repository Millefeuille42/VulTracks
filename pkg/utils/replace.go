package utils

import (
	"reflect"
)

// Replace fills the first struct with the listed second structs fields if they are not empty
func Replace(t1, t2 interface{}, fields ...string) {
	e1 := reflect.ValueOf(t1).Elem()
	e2 := reflect.ValueOf(t2).Elem()

	if e1.Kind() != reflect.Struct || e2.Kind() != reflect.Struct {
		return
	}

	for _, field := range fields {
		if !e2.FieldByName(field).IsZero() {
			e1.FieldByName(field).Set(e2.FieldByName(field))
		}
	}
}
