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

func ThrowMsgF[T any](format string, args any) {
	exception := iNewException[T](fmt.Errorf(format, args))
	Throw(exception)
}

func iThrow(exception *Exception) {
	panic(exception)
}
