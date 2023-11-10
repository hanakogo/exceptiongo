package exceptiongo

import (
	"fmt"
	"github.com/hanakogo/hanakoutilgo"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type Exception struct {
	error
	exType     reflect.Type
	stackTrace []string
}

func NewExceptionF[T any](format string, a ...any) *Exception {
	var err error
	switch len(a) {
	case 0:
		err = fmt.Errorf(format)
	default:
		err = fmt.Errorf(format, a)
	}
	return iNewException[T](err)
}

func NewException[T any](message string) *Exception {
	return iNewException[T](fmt.Errorf(message))
}

func NewErrException[T any](err error) *Exception {
	if err == nil {
		return nil
	}
	return iNewException[T](err)
}

func iNewException[T any](err error) *Exception {
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
		// hardcode pop last 2 elements to remove common basic level stacktrace
		stackTrace = stackTrace[:len(stackTrace)-2]
		return
	}
	exType := hanakoutilgo.TypeOf[T]()
	if exType.Kind() == reflect.Ptr {
		exType = exType.Elem()
	}
	return &Exception{
		error:      err,
		exType:     exType,
		stackTrace: getStackTrace(),
	}
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
	return e.exType
}

func (e *Exception) Error() error {
	return e.error
}

func (e *Exception) TypeName() string {
	return e.Type().String()
}
