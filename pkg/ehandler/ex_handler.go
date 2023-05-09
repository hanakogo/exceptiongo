package ehandler

import (
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"github.com/ohanakogo/exceptiongo/pkg/exutil"
)

type ExceptionHandler struct {
	OnHandle func(*etype.Exception)
}

func (e *ExceptionHandler) Handle(ex *etype.Exception) {
	e.OnHandle(ex)
}

func (e *ExceptionHandler) GlobalHandle() {
	defer exutil.HandleRecoverException(func(exception *etype.Exception) {
		e.OnHandle(exception)
	})
	if r := recover(); r != nil {
		panic(r)
	}
}
