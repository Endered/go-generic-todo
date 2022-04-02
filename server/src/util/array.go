package util

import "sort"

func Map[T any, U any](arr []T, f func(T) U) []U {
	var res []U
	for _, v := range arr {
		res = append(res, f(v))
	}
	return res
}

func Foreach[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func Filter[T any](arr []T, f func(T) bool) []T {
	var res []T
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func Split[T any](arr []T, f func(T) bool) Tuple2[[]T, []T] {
	var trues []T
	var falses []T
	for _, v := range arr {
		if f(v) {
			trues = append(trues, v)
		} else {
			falses = append(falses, v)
		}
	}
	return NewTuple2(trues, falses)
}

func Take[T any](arr []T, n int) []T {
	if len(arr) <= n {
		return arr
	} else {
		return arr[:n]
	}
}

func Skip[T any](arr []T, n int) []T {
	if len(arr) <= n {
		return nil
	} else {
		return arr[n:]
	}
}

func GroupBy[T any, U comparable](arr []T, f func(T) U) map[U][]T {
	var res map[U][]T = make(map[U][]T)
	for _, v := range arr {
		res[f(v)] = append(res[f(v)], v)
	}
	return res
}

func FlatMap[T any, U any](arr []T, f func(T) []U) []U {
	var res []U
	for _, v := range arr {
		res = append(res, f(v)...)
	}
	return res
}

func SortedBy[T any](arr []T, less func(T, T) bool) []T {
	var res []T
	Foreach(arr, func(v T) { res = append(res, v) })
	sort.Slice(res, func(i, j int) bool {
		return less(res[i], res[j])
	})
	return res
}
