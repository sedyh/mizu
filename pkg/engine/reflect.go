package engine

import (
	"reflect"
)

func assertKind(v any, k reflect.Kind, fail func()) {
	value := reflect.ValueOf(v)
	if value.Kind() != k {
		fail()
	}
}

//func initAnyType[T any]() *T {
//	var x T
//
//	t := reflect.TypeOf(x)
//	v := reflect.New(t)
//
//	if t.Kind() == reflect.Struct {
//		walkInitStruct(t, v.Elem())
//	}
//
//	return v.Interface().(*T)
//}
//
//func walkInitStruct(t reflect.Type, v reflect.Value) {
//	for i := 0; i < v.NumField(); i++ {
//		f := v.Field(i)
//		ft := t.Field(i)
//		switch ft.Type.Kind() {
//		case reflect.Map:
//			f.Set(reflect.MakeMap(ft.Type))
//		case reflect.Slice:
//			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
//		case reflect.Chan:
//			f.Set(reflect.MakeChan(ft.Type, 0))
//		case reflect.Struct:
//			walkInitStruct(ft.Type, f)
//		case reflect.Ptr:
//			fv := reflect.New(ft.Type.Elem())
//			walkInitStruct(ft.Type.Elem(), fv.Elem())
//			f.Set(fv)
//		default:
//		}
//	}
//}
