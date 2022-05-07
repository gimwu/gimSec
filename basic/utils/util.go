package utils

import (
	"fmt"
	"reflect"
)

func ToStringSlice(params interface{}) *[]string {
	res := make([]string, 0)
	switch reflect.TypeOf(params).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(params)
		for i := 0; i < s.Len(); i++ {
			index := fmt.Sprintf("%s", s.Index(i))
			res = append(res, index)
		}
	}
	return &res
}
