package utils

import "reflect"

func StructToMap(input interface{}) map[string]string {
	v := reflect.ValueOf(input)
	values := make(map[string]string)
	for i := 0; i < v.NumField(); i++ {
		values[v.Type().Field(i).Name] = v.Field(i).String()
	}
	return values
}
