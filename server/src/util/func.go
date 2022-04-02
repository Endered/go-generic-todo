package util

func Combine[T, U, V any](f func(T) U, g func(U) V) func(T) V {
	return func(v T) V {
		return g(f(v))
	}
}
