package common

import (
	"errors"
	"fmt"
	"reflect"
)

// Empty returns an empty value of the given type
func Empty[T any]() T {
	var empty T
	return empty
}

// ToPtr converts a value to a pointer.
func ToPtr[T any](v T) *T {
	return &v
}

// ToValue converts a pointer to a value.
func ToValue[T any](v *T) T {
	if v == nil {
		return Empty[T]()
	}

	return *v
}

// IsEmpty returns true if argument is a empty value
func IsEmpty[T comparable](v T) bool {
	var empty T
	return v == empty
}

// IsEmpty2 反射判断是否为空值
func IsEmpty2[T any](v T) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

// IsNotEmpty returns true if argument is not a empty value
func IsNotEmpty[T comparable](v T) bool {
	var empty T
	return v != empty
}

// CheckStructEmptyFields check struct has at least one field is empty
func CheckStructEmptyFields(s any) error {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return errors.New("only struct is supported")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name

		// 检查字段的值是否为空值
		if IsEmpty2(field.Interface()) {
			return fmt.Errorf("struct field %s can not empty", fieldName)
		}
	}
	return nil
}
