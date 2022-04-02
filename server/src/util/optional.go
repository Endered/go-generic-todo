package util

type Optional[T any] struct {
	Value T
	Has   bool
}

func Some[T any](v T) Optional[T] {
	return Optional[T]{Value: v, Has: true}
}

func None[T any]() Optional[T] {
	return Optional[T]{Has: false}
}

func MapOptional[T any, U any](m Optional[T], f func(T) U) Optional[U] {
	if m.Has {
		return Optional[U]{Value: f(m.Value), Has: true}
	} else {
		return Optional[U]{Has: false}
	}
}

func AppendOptional[T any](m Optional[Optional[T]]) Optional[T] {
	if m.Has {
		return m.Value
	} else {
		return None[T]()
	}
}

func FlatMapOptional[T any, U any](m Optional[T], f func(T) Optional[U]) Optional[U] {
	if m.Has {
		return f(m.Value)
	} else {
		return None[U]()
	}
}

func OptionalValues[T any](m Optional[T]) (T, bool) {
	return m.Value, m.Has
}

func OrElse[T any](m Optional[T], v T) T {
	if m.Has {
		return m.Value
	} else {
		return v
	}
}

func OrElseFunc[T any](m Optional[T], f func() T) T {
	if m.Has {
		return m.Value
	} else {
		return f()
	}
}
