package util

type Tuple2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

func NewTuple2[T1, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{
		V1: v1,
		V2: v2,
	}
}

func (t Tuple2[T1, T2]) Values() (T1, T2) {
	return t.V1, t.V2
}

func ZipWith2[T, U any](arr1 []T, arr2 []U) []Tuple2[T, U] {
	var res []Tuple2[T, U]
	for i := 0; i < len(arr1) && i < len(arr2); i += 1 {
		res = append(res, NewTuple2(arr1[i], arr2[i]))
	}
	return res
}
