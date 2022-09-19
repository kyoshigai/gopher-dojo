package main

import (
	"errors"
	"reflect"
)

var errCannotSetString = errors.New("cannot set string")

func setString(ps interface{}, s string) error {
	v := reflect.ValueOf(ps)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.String {
		return errCannotSetString
	}
	v.Elem().Set(reflect.ValueOf(s))
	return nil
}
