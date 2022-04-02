package util

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elements ...T) Set[T] {
	res := map[T]struct{}{}
	Foreach(elements, func(v T) {
		res[v] = struct{}{}
	})
	return res
}

func SetToList[T comparable](s Set[T]) []T {
	return MapKeys(s)
}

func SetHas[T comparable](s Set[T], e T) bool {
	_, ok := s[e]
	return ok
}

func Add[T comparable](s Set[T], e T) {
	s[e] = struct{}{}
}

func Del[T comparable](s Set[T], e T) {
	if SetHas(s, e) {
		delete(s, e)
	}
}
