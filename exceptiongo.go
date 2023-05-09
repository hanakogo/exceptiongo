package exceptiongo

import (
	"fmt"
	"github.com/ohanakogo/exceptiongo/internal/_exutil"
	"github.com/ohanakogo/exceptiongo/pkg/ehandler"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"github.com/ohanakogo/exceptiongo/pkg/exutil"
	"github.com/ohanakogo/ohanakoutilgo"
)

func Throw(exception *etype.Exception) {
	_exutil.InternalThrow(exception)
}

func TryHandle[T any](do func()) {
	defer exutil.HandleRecoverException(func(exception *etype.Exception) {
		switch {
		case exception.Compare(ohanakoutilgo.TypeOf[T]()):
			do()
		default:
			Throw(exception)
		}
	})
	if r := recover(); r != nil {
		panic(r)
	}
}

func NewExceptionF[T any](format string, a ...any) *etype.Exception {
	var err error
	switch len(a) {
	case 0:
		err = fmt.Errorf(format)
	default:
		err = fmt.Errorf(format, a)
	}
	return etype.InternalException[T](err)
}

func NewException[T any](message string) *etype.Exception {
	return etype.InternalException[T](fmt.Errorf(message))
}

func ToException[T any](err error) *etype.Exception {
	if err == nil {
		return nil
	}
	return etype.InternalException[T](err)
}

func NewDefaultExceptionHandler() *ehandler.ExceptionHandler {
	return NewExceptionHandler(func(e *etype.Exception) {
		panic(e)
	})
}

func NewExceptionHandler(handle func(*etype.Exception)) (e *ehandler.ExceptionHandler) {
	e = new(ehandler.ExceptionHandler)
	e.OnHandle = handle
	return
}
