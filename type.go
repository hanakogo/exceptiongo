package exceptiongo

import (
	"github.com/hanakogo/hanakoutilgo"
	"reflect"
)

// Compare check basic type is T or not
func Compare[T any](e *Exception) bool {
	return TypeOf[T]() == e.Type()
}

// TypeOf get basic data type of T
func TypeOf[T any]() reflect.Type {
	return hanakoutilgo.ActualTypeOf[T]()
}
