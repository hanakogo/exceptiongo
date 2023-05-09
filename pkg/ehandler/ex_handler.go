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
	exutil.HandleRecoverException(func(exception *etype.Exception) {
		e.OnHandle(exception)
	})
}
