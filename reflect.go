package mizu

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

func mapTypesToInterface(types []reflect.Type) []interface{} {
	res := make([]interface{}, 0, 2)
	for _, t := range types {
		if !exported(t) {
			continue
		}
		//fmt.Println("loose", t.Name())
		//tt := reflect.ValueOf(t).Interface()
		//fmt.Println("geeet", reflect.ValueOf(tt).Type().Name())
		res = append(res, t)
	}
	return res
}

func exported(v interface{}) bool {
	return reflect.ValueOf(v).CanInterface()
}
