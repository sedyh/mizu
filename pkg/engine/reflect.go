package engine

import (
	"reflect"
)

func structFieldTypes(structType reflect.Value) []interface{} {
	res := make([]interface{}, 0, 2)
	for i := 0; i < structType.NumField(); i++ {
		valueField := structType.Field(i)
		if !valueField.CanInterface() {
			continue
		}
		res = append(res, valueField.Interface())
	}
	return res
}

func typeName(t reflect.Type) string {
	name := t.Name()
	if name != "" {
		return name
	}
	return "anonymous"
}
