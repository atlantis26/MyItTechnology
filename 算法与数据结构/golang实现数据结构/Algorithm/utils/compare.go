package utils

import (
	"reflect"
)

func Compare(a interface{}, b interface{}) (ret int) {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()

	if aType != bType {
		panic("cannot compare different type params")
	}
	switch a.(type){
	case int:
		if a.(int) > b.(int) {
			ret = 1
		} else if a.(int) < b.(int) {
			ret = -1
		} else {
			ret = 0
		}
	case string:
		if a.(string) > b.(string) {
			ret = 1
		} else if a.(string) < b.(string) {
			ret = -1
		} else {
			ret = 0
		}
	case float64:
		if a.(float64) > b.(float64) {
			ret = 1
		} else if a.(float64) < b.(float64) {
			ret = -1
		} else {
			ret = 0
		}
	default:
		panic("unsupported type params")
	}
	return ret
}
