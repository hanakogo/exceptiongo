package exceptiongo

import (
	"fmt"
	"github.com/gookit/slog"
	"github.com/hanakogo/hanakoutilgo"
	"runtime"
)

func handleRecoverException(do func(exception *Exception)) {
	if r := recover(); r != nil {
		if !hanakoutilgo.Is[*Exception](r) {
			iThrow(iNewException[any](fmt.Errorf("handle failed: \"%v\" is not a Exception, panic with default function", r)))
		}
		exception := hanakoutilgo.CastTo[*Exception](r)
		do(exception)
	}
}

func handleExceptionIsNil(exception *Exception, postprocess func()) {
	if exception == nil {
		if postprocess != nil {
			postprocess()
		}
	}
}

func handleException(exception *Exception, postprocess func()) {
	if exception != nil {
		_, file, line, _ := runtime.Caller(1)
		slog.Errorf("exception encountered[%s:%d]", file, line)
		slog.Errorf("message: %s", exception.Error())
		if postprocess != nil {
			postprocess()
		}
	}
}
