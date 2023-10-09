package test

import (
	"fmt"
	"github.com/hanakogo/exceptiongo"
	"github.com/hanakogo/hanakoutilgo"
	"testing"
)

type Common any

type Standard any

type MustThrow any

var GlobalHandler *exceptiongo.ExceptionHandler

func init() {
	GlobalHandler = exceptiongo.NewDefaultExceptionHandler()
	GlobalHandler.OnHandle = func(e *exceptiongo.Exception) {
		switch e.Type() {
		case hanakoutilgo.ActualTypeOf[Common]():
			fmt.Println("catching a common e:", e.GetStackTraceMessage())

		case hanakoutilgo.ActualTypeOf[Standard]():
			fmt.Println("catching a standard e:", e.GetStackTraceMessage())

		case hanakoutilgo.ActualTypeOf[MustThrow]():
			fmt.Println("catching a must throw e and simulated throw:", e.GetStackTraceMessage())
		}
	}
}

func TestException(t *testing.T) {
	ex := exceptiongo.NewExceptionF[Common]("test error")

	switch ex.Type() {
	case hanakoutilgo.ActualTypeOf[Common]():
		fmt.Println("exception has been detected")
	}

	ex.PrintStackTrace()
}

func TestExceptionHandler(t *testing.T) {
	commonException := exceptiongo.NewExceptionF[Common]("test common error")
	standardException := exceptiongo.NewException[Standard]("test standard error")
	mustThrowException := exceptiongo.NewExceptionF[MustThrow]("test must throw error")

	handler := exceptiongo.NewExceptionHandler(func(e *exceptiongo.Exception) {
		switch e.Type() {
		case hanakoutilgo.ActualTypeOf[MustThrow]():
			exceptiongo.Throw(e)
		default:
			fmt.Printf("normally handle exception: %s", e.GetStackTraceMessage())
		}
	})
	handler.Handle(func() {
		exceptiongo.Throw(commonException)
	})
	handler.Handle(func() {
		exceptiongo.Throw(standardException)
	})

	defer GlobalHandler.Deploy()
	exceptiongo.Throw(mustThrowException)
}

func TestTryCatch(t *testing.T) {
	commonException := exceptiongo.NewExceptionF[Common]("test common error")

	exceptiongo.TryCatch[Common](func() {
		exceptiongo.Throw(commonException)
	}, func(exception *exceptiongo.Exception) {
		t.Log(exception.GetStackTraceMessage())
	})

	exceptiongo.TryCatch[any](func() {
		exceptiongo.Throw(commonException)
	}, func(exception *exceptiongo.Exception) {
		t.Logf("the type of [%v] has been catched by the type of [any] catcher!", exception.Type())
		t.Log(exception.GetStackTraceMessage())
	})
}
