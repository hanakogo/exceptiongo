package exceptiongo

import "fmt"

func Throw(exception *Exception) {
	iThrow(exception)
}

func ThrowErr[T any](err error) {
	if err == nil {
		return
	}
	exception := iNewException[T](err)
	Throw(exception)
}

func ThrowMsg[T any](msg string) {
	exception := iNewException[T](fmt.Errorf(msg))
	Throw(exception)
}

func ThrowMsgF[T any](format string, args ...any) {
	var err error
	if len(args) > 0 {
		err = fmt.Errorf(format, args...)
	} else {
		err = fmt.Errorf(format)
	}
	exception := iNewException[T](err)
	Throw(exception)
}

func iThrow(exception *Exception) {
	panic(exception)
}
