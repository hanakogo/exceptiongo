package exceptiongo

type ExceptionHandler struct {
	OnHandle func(*Exception)
}

func NewDefaultExceptionHandler() *ExceptionHandler {
	return NewExceptionHandler(func(e *Exception) {
		panic(e)
	})
}

func NewExceptionHandler(handle func(*Exception)) (e *ExceptionHandler) {
	e = new(ExceptionHandler)
	e.OnHandle = handle
	return
}

func (e *ExceptionHandler) Deploy() {
	defer handleRecoverException(func(exception *Exception) {
		e.OnHandle(exception)
	})
	if r := recover(); r != nil {
		panic(r)
	}
}

func (e *ExceptionHandler) Handle(fun func()) {
	defer e.Deploy()
	fun()
}
