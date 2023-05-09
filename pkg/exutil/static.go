package exutil

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/ohanakogo/exceptiongo/internal/_exutil"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"github.com/ohanakogo/ohanakoutilgo"
	"runtime"
)

func HandleRecoverException(do func(exception *etype.Exception)) {
	if r := recover(); r != nil {
		if !ohanakoutilgo.Is[*etype.Exception](r) {
			_exutil.InternalThrow(etype.InternalException[any](fmt.Errorf("handle failed: \"%v\" is not a Exception, panic with default function", r)))
		}
		exception := ohanakoutilgo.CastTo[*etype.Exception](r)
		do(exception)
	}
}

func HandleExceptionIsNil(exception *etype.Exception, postprocess func()) {
	if exception == nil {
		if postprocess != nil {
			postprocess()
		}
	}
}

func HandleException(exception *etype.Exception, postprocess func()) {
	if exception != nil {
		_, file, line, _ := runtime.Caller(1)
		slog.Errorf("exception encountered[%s:%d]", file, line)
		slog.Errorf("message: %s", exception.Error())
		if postprocess != nil {
			postprocess()
		}
	}
}
