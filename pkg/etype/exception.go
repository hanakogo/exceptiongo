package etype

import (
	"fmt"
	"github.com/ohanakogo/ohanakoutilgo"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type Exception struct {
	error
	kind       reflect.Type
	stackTrace []string
}

func InternalException[T any](err error) *Exception {
	getStackTrace := func() (stackTrace []string) {
		for i := 3; ; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			funcName := runtime.FuncForPC(pc).Name()
			filePath := strings.Split(file, "/")
			fileName := filePath[len(filePath)-1]
			stackTrace = append(stackTrace, fmt.Sprintf("%s <%s:%d>", funcName, fileName, line))
		}
		return
	}
	return &Exception{
		error:      err,
		kind:       ohanakoutilgo.TypeOf[T]().Elem(),
		stackTrace: getStackTrace(),
	}
}

func (e *Exception) Compare(p reflect.Type) bool {
	return e.Type() == p
}

func (e *Exception) GetStackTraceMessage() string {
	parseOutputStackTrace := func() (output string) {
		for _, s := range e.stackTrace {
			output += fmt.Sprintf("\t -> at %s\n", s)
		}
		return
	}
	return fmt.Sprintf(
		"Exception[%s] encountered: %s\n%s",
		e.TypeName(),
		e.Error(),
		parseOutputStackTrace(),
	)
}

func (e *Exception) PrintStackTrace() {
	_, err := fmt.Fprint(os.Stderr, e.GetStackTraceMessage())
	if err != nil {
		panic(err)
	}
}

func (e *Exception) Type() reflect.Type {
	if e == nil {
		return nil
	}
	return e.kind
}

func (e *Exception) Error() error {
	return e.error
}

func (e *Exception) TypeName() string {
	return e.Type().String()
}
