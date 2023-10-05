package exceptiongo

import "fmt"

func Throw(exception *Exception) {
	iThrow(exception)
}

func QuickThrow[T any](err error) {
	if err == nil {
		return
	}
	exception := iNewException[T](err)
	Throw(exception)
}

func QuickThrowMsg[T any](msg string) {
	exception := iNewException[T](fmt.Errorf(msg))
	Throw(exception)
}

func iThrow(exception *Exception) {
	panic(exception)
}
