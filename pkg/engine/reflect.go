package engine

import (
	"reflect"
)

// structFieldTypes returns exported fields of the struct
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

// typeName returns the type if there is one or "anonymous" if it is not
func typeName(t reflect.Type) string {
	name := t.Name()
	if name != "" {
		return name
	}
	return "anonymous"
}
