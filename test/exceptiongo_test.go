package test

import (
	"fmt"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/exceptiongo/pkg/ehandler"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"github.com/ohanakogo/ohanakoutilgo"
	"testing"
)

type Common any

type Standard any

type MustThrow any

var GlobalHandler *ehandler.ExceptionHandler

func init() {
	GlobalHandler = exceptiongo.NewDefaultExceptionHandler()
	GlobalHandler.OnHandle = func(exception *etype.Exception) {
		switch exception.Type() {
		case ohanakoutilgo.ActualTypeOf[Common]():
			fmt.Println("catching a common exception:", exception)

		case ohanakoutilgo.ActualTypeOf[Standard]():
			fmt.Println("catching a standard exception:", exception)

		case ohanakoutilgo.ActualTypeOf[MustThrow]():
			fmt.Println("catching a must throw exception:", exception)
			exceptiongo.Throw(exception)
		}
	}
}

func TestException(t *testing.T) {
	ex := exceptiongo.NewExceptionF[Common]("test error")

	switch ex.Type() {
	case ohanakoutilgo.ActualTypeOf[Common]():
		fmt.Println("exception has been detected")
	}

	ex.PrintStackTrace()
}

func TestExceptionHandler(t *testing.T) {
	commonException := exceptiongo.NewExceptionF[Common]("test common error")
	standardException := exceptiongo.NewException[Standard]("test standard error")
	mustThrowException := exceptiongo.NewExceptionF[MustThrow]("test must throw error")

	handler := exceptiongo.NewExceptionHandler(func(e *etype.Exception) {
		switch e.Type() {
		case ohanakoutilgo.ActualTypeOf[MustThrow]():
			exceptiongo.Throw(e)
		default:
			fmt.Printf("normally handle exception: %v\n", e)
		}
	})
	handler.Handle(commonException)
	handler.Handle(standardException)

	defer GlobalHandler.Deploy()
	exceptiongo.Throw(mustThrowException)
}

func TestTryCatch(t *testing.T) {
	commonException := exceptiongo.NewExceptionF[Common]("test common error")

	exceptiongo.TryCatch[Common](func() {
		exceptiongo.Throw(commonException)
	}, func(exception *etype.Exception) {
		t.Log(exception.GetStackTraceMessage())
	})

	exceptiongo.TryCatch[any](func() {
		exceptiongo.Throw(commonException)
	}, func(exception *etype.Exception) {
		t.Logf("the type of [%v] has been catched by the type of [any] catcher!", exception.Type())
		t.Log(exception.GetStackTraceMessage())
	})
}
