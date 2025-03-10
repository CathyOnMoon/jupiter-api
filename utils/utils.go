package utils

import (
	"fmt"
	"reflect"
)

// StructToValues Helper function to convert struct fields to URL values
func StructToValues(v interface{}) map[string]string {
	val := reflect.ValueOf(v)
	typ := val.Type()

	result := make(map[string]string)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("url")

		if tag != "" {
			if !IsZeroValue(field) {
				switch field.Kind() {
				case reflect.String:
					result[tag] = field.String()
				case reflect.Int, reflect.Int64:
					result[tag] = fmt.Sprintf("%d", field.Int())
				case reflect.Bool:
					result[tag] = fmt.Sprintf("%t", field.Bool())
				case reflect.Slice:
					for j := 0; j < field.Len(); j++ {
						result[fmt.Sprintf("%s[]", tag)] = field.Index(j).String()
					}
				}
			}
		}
	}

	return result
}

func IsZeroValue(field reflect.Value) bool {
	zero := reflect.Zero(field.Type())
	return reflect.DeepEqual(field.Interface(), zero.Interface())
}
