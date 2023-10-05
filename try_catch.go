package exceptiongo

import (
	"github.com/hanakogo/hanakoutilgo"
	"reflect"
)

func TryCatch[T any](try func(), catch func(exception *Exception)) {
	defer handleRecoverException(func(exception *Exception) {
		exType := exception.Type()
		switch {
		case exType == hanakoutilgo.TypeOf[T]():
			catch(exception)
		case hanakoutilgo.TypeOf[T]().Kind() == reflect.Interface:
			catch(exception)
		default:
			Throw(exception)
		}
	})
	try()
}
