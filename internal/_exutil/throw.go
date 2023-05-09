package _exutil

import "github.com/ohanakogo/exceptiongo/pkg/etype"

func InternalThrow(exception *etype.Exception) {
	panic(exception)
}
