package util

type Result[T any] struct {
	Value   T
	Error   error
	Success bool
}

func ResultValues[T any](r Result[T]) (T, error) {
	return r.Value, r.Error
}

func NewResultSuccess[T any](v T) Result[T] {
	return Result[T]{
		Value:   v,
		Success: true,
	}
}
func NewResultFailed[T any](e error) Result[T] {
	return Result[T]{
		Error:   e,
		Success: false,
	}
}

func NewResult[T any](v T, e error) Result[T] {
	if e != nil {
		return NewResultFailed[T](e)
	} else {
		return NewResultSuccess(v)
	}
}

func MapResult[T, U any](r Result[T], f func(T) U) Result[U] {
	if r.Success {
		return NewResultSuccess(f(r.Value))
	}
	return NewResultFailed[U](r.Error)
}

func MapResultError[T any](r Result[T], f func(error) error) Result[T] {
	if r.Success {
		return r
	} else {
		return NewResultFailed[T](f(r.Error))
	}
}

func FlatMapResult[T, U any](r Result[T], f func(T) Result[U]) Result[U] {
	if r.Success {
		return f(r.Value)
	} else {
		return NewResultFailed[U](r.Error)
	}
}
