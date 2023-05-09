package etype

import (
	"github.com/ohanakogo/ohanakoutilgo"
	"reflect"
)

type Exception struct {
	error
	kind reflect.Type
}

func InternalException[T any](err error) *Exception {
	return &Exception{
		error: err,
		kind:  ohanakoutilgo.TypeOf[T](),
	}
}

func (e *Exception) Type() reflect.Type {
	if e == nil {
		return nil
	}
	return e.kind
}

func (e *Exception) Compare(p reflect.Type) bool {
	return e.Type() == p
}
